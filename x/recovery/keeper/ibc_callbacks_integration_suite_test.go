package keeper_test

import (
	"strconv"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	ibcgotesting "github.com/cosmos/ibc-go/v6/testing"

	ibctesting "github.com/Atrix/Atrix/v11/ibc/testing"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/Atrix/Atrix/v11/app"
	claimstypes "github.com/Atrix/Atrix/v11/x/claims/types"
	inflationtypes "github.com/Atrix/Atrix/v11/x/inflation/types"
	"github.com/Atrix/Atrix/v11/x/recovery/types"
)

type IBCTestingSuite struct {
	suite.Suite
	coordinator *ibcgotesting.Coordinator

	// testing chains used for convenience and readability
	AtrixChain      *ibcgotesting.TestChain
	IBCOsmosisChain *ibcgotesting.TestChain
	IBCCosmosChain  *ibcgotesting.TestChain

	pathOsmosisAtrix  *ibcgotesting.Path
	pathCosmosAtrix   *ibcgotesting.Path
	pathOsmosisCosmos *ibcgotesting.Path
}

var s *IBCTestingSuite

func TestIBCTestingSuite(t *testing.T) {
	s = new(IBCTestingSuite)
	suite.Run(t, s)

	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keeper Suite")
}

func (suite *IBCTestingSuite) SetupTest() {
	// initializes 3 test chains
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 1, 2)
	suite.AtrixChain = suite.coordinator.GetChain(ibcgotesting.GetChainID(1))
	suite.IBCOsmosisChain = suite.coordinator.GetChain(ibcgotesting.GetChainID(2))
	suite.IBCCosmosChain = suite.coordinator.GetChain(ibcgotesting.GetChainID(3))
	suite.coordinator.CommitNBlocks(suite.AtrixChain, 2)
	suite.coordinator.CommitNBlocks(suite.IBCOsmosisChain, 2)
	suite.coordinator.CommitNBlocks(suite.IBCCosmosChain, 2)

	// Mint coins locked on the Atrix account generated with secp.
	coinAtrix := sdk.NewCoin("aAtrix", sdk.NewInt(10000))
	coins := sdk.NewCoins(coinAtrix)
	err := suite.AtrixChain.App.(*app.Atrix).BankKeeper.MintCoins(suite.AtrixChain.GetContext(), inflationtypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.AtrixChain.App.(*app.Atrix).BankKeeper.SendCoinsFromModuleToAccount(suite.AtrixChain.GetContext(), inflationtypes.ModuleName, suite.IBCOsmosisChain.SenderAccount.GetAddress(), coins)
	suite.Require().NoError(err)

	// Mint coins on the osmosis side which we'll use to unlock our aAtrix
	coinOsmo := sdk.NewCoin("uosmo", sdk.NewInt(10))
	coins = sdk.NewCoins(coinOsmo)
	err = suite.IBCOsmosisChain.GetSimApp().BankKeeper.MintCoins(suite.IBCOsmosisChain.GetContext(), minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.IBCOsmosisChain.GetSimApp().BankKeeper.SendCoinsFromModuleToAccount(suite.IBCOsmosisChain.GetContext(), minttypes.ModuleName, suite.IBCOsmosisChain.SenderAccount.GetAddress(), coins)
	suite.Require().NoError(err)

	// Mint coins on the cosmos side which we'll use to unlock our aAtrix
	coinAtom := sdk.NewCoin("uatom", sdk.NewInt(10))
	coins = sdk.NewCoins(coinAtom)
	err = suite.IBCCosmosChain.GetSimApp().BankKeeper.MintCoins(suite.IBCCosmosChain.GetContext(), minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.IBCCosmosChain.GetSimApp().BankKeeper.SendCoinsFromModuleToAccount(suite.IBCCosmosChain.GetContext(), minttypes.ModuleName, suite.IBCCosmosChain.SenderAccount.GetAddress(), coins)
	suite.Require().NoError(err)

	claimparams := claimstypes.DefaultParams()
	claimparams.AirdropStartTime = suite.AtrixChain.GetContext().BlockTime()
	claimparams.EnableClaims = true
	suite.AtrixChain.App.(*app.Atrix).ClaimsKeeper.SetParams(suite.AtrixChain.GetContext(), claimparams)

	params := types.DefaultParams()
	params.EnableRecovery = true
	suite.AtrixChain.App.(*app.Atrix).RecoveryKeeper.SetParams(suite.AtrixChain.GetContext(), params)

	suite.pathOsmosisAtrix = ibctesting.NewTransferPath(suite.IBCOsmosisChain, suite.AtrixChain) // clientID, connectionID, channelID empty
	suite.pathCosmosAtrix = ibctesting.NewTransferPath(suite.IBCCosmosChain, suite.AtrixChain)
	suite.pathOsmosisCosmos = ibctesting.NewTransferPath(suite.IBCCosmosChain, suite.IBCOsmosisChain)
	suite.coordinator.Setup(suite.pathOsmosisAtrix) // clientID, connectionID, channelID filled
	suite.coordinator.Setup(suite.pathCosmosAtrix)
	suite.coordinator.Setup(suite.pathOsmosisCosmos)
	suite.Require().Equal("07-tendermint-0", suite.pathOsmosisAtrix.EndpointA.ClientID)
	suite.Require().Equal("connection-0", suite.pathOsmosisAtrix.EndpointA.ConnectionID)
	suite.Require().Equal("channel-0", suite.pathOsmosisAtrix.EndpointA.ChannelID)
}

var timeoutHeight = clienttypes.NewHeight(1000, 1000)

func (suite *IBCTestingSuite) SendAndReceiveMessage(path *ibcgotesting.Path, origin *ibcgotesting.TestChain, coin string, amount int64, sender string, receiver string, seq uint64) {
	// Send coin from A to B
	transferMsg := transfertypes.NewMsgTransfer(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, sdk.NewCoin(coin, sdk.NewInt(amount)), sender, receiver, timeoutHeight, 0, "")
	_, err := origin.SendMsgs(transferMsg)
	suite.Require().NoError(err) // message committed
	// Recreate the packet that was sent
	transfer := transfertypes.NewFungibleTokenPacketData(coin, strconv.Itoa(int(amount)), sender, receiver, "")
	packet := channeltypes.NewPacket(transfer.GetBytes(), seq, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, 0)
	// Receive message on the counterparty side, and send ack
	err = path.RelayPacket(packet)
	suite.Require().NoError(err)
}

func CreatePacket(amount, denom, sender, receiver, srcPort, srcChannel, dstPort, dstChannel string, seq, timeout uint64) channeltypes.Packet {
	transfer := transfertypes.FungibleTokenPacketData{
		Amount:   amount,
		Denom:    denom,
		Receiver: sender,
		Sender:   receiver,
	}
	return channeltypes.NewPacket(
		transfer.GetBytes(),
		seq,
		srcPort,
		srcChannel,
		dstPort,
		dstChannel,
		clienttypes.ZeroHeight(), // timeout height disabled
		timeout,
	)
}
