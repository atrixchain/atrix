#!/bin/bash

KEY="dev0"
CHAINID="Atrix_9000-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t Atrix-datadir.XXXXX)

echo "create and add new keys"
./Atrixd keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Atrix with moniker=$MONIKER and chain-id=$CHAINID"
./Atrixd init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./Atrixd add-genesis-account \
"$(./Atrixd keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000aAtrix,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./Atrixd gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./Atrixd collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./Atrixd validate-genesis --home $DATA_DIR

echo "starting Atrix node $i in background ..."
./Atrixd start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started Atrix node"
tail -f /dev/null