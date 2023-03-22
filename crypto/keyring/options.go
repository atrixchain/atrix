// Copyright 2022 Atrix Foundation
// This file is part of the Atrix Network packages.
//
// Atrix is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Atrix packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Atrix packages. If not, see https://github.com/Atrix/Atrix/blob/main/LICENSE

package keyring

import (
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cosmosLedger "github.com/cosmos/cosmos-sdk/crypto/ledger"
	"github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/Atrix/ethermint/crypto/ethsecp256k1"
	"github.com/Atrix/ethermint/crypto/hd"
	"github.com/Atrix/Atrix-ledger-go/ledger"
)

// AppName defines the Ledger app used for signing. Atrix uses the Ethereum app
const AppName = "Ethereum"

var (
	// SupportedAlgorithms defines the list of signing algorithms used on Atrix:
	//  - eth_secp256k1 (Ethereum)
	SupportedAlgorithms = keyring.SigningAlgoList{hd.EthSecp256k1}
	// SupportedAlgorithmsLedger defines the list of signing algorithms used on Atrix for the Ledger device:
	//  - secp256k1 (in order to comply with Cosmos SDK)
	// The Ledger derivation function is responsible for all signing and address generation.
	SupportedAlgorithmsLedger = keyring.SigningAlgoList{hd.EthSecp256k1}
	// LedgerDerivation defines the Atrix Ledger Go derivation (Ethereum app with EIP-712 signing)
	LedgerDerivation = ledger.AtrixLedgerDerivation()
	// CreatePubkey uses the ethsecp256k1 pubkey with Ethereum address generation and keccak hashing
	CreatePubkey = func(key []byte) types.PubKey { return &ethsecp256k1.PubKey{Key: key} }
	// SkipDERConversion represents whether the signed Ledger output should skip conversion from DER to BER.
	// This is set to true for signing performed by the Ledger Ethereum app.
	SkipDERConversion = true
)

// EthSecp256k1Option defines a function keys options for the ethereum Secp256k1 curve.
// It supports eth_secp256k1 keys for accounts.
func Option() keyring.Option {
	return func(options *keyring.Options) {
		options.SupportedAlgos = SupportedAlgorithms
		options.SupportedAlgosLedger = SupportedAlgorithmsLedger
		options.LedgerDerivation = func() (cosmosLedger.SECP256K1, error) { return LedgerDerivation() }
		options.LedgerCreateKey = CreatePubkey
		options.LedgerAppName = AppName
		options.LedgerSigSkipDERConv = SkipDERConversion
	}
}
