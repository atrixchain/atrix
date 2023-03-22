#!/bin/bash

KEY="mykey"
CHAINID="${CHAIN_ID:-Atrix_9000-1}"
MONIKER="localtestnet"
KEYRING="test" # remember to change to other types of keyring like 'file' in-case exposing to outside world, otherwise your balance will be wiped quickly. The keyring test does not require private key to steal tokens from you
KEYALGO="eth_secp256k1"
LOGLEVEL="info"
# to trace evm
#TRACE="--trace"
TRACE=""

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# used to exit on first error (any non-zero exit code)
set -e

# Set client config
Atrixd config keyring-backend "$KEYRING"
Atrixd config chain-id "$CHAINID"

# if $KEY exists it should be deleted
Atrixd keys add "$KEY" --keyring-backend $KEYRING --algo "$KEYALGO"

# Set moniker and chain-id for Atrix (Moniker can be anything, chain-id must be an integer)
Atrixd init "$MONIKER" --chain-id "$CHAINID"

# Change parameter token denominations to aAtrix
cat "$HOME"/.Atrixd/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="aAtrix"' > "$HOME"/.Atrixd/config/tmp_genesis.json && mv "$HOME"/.Atrixd/config/tmp_genesis.json "$HOME"/.Atrixd/config/genesis.json
cat "$HOME"/.Atrixd/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="aAtrix"' > "$HOME"/.Atrixd/config/tmp_genesis.json && mv "$HOME"/.Atrixd/config/tmp_genesis.json "$HOME"/.Atrixd/config/genesis.json
cat "$HOME"/.Atrixd/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="aAtrix"' > "$HOME"/.Atrixd/config/tmp_genesis.json && mv "$HOME"/.Atrixd/config/tmp_genesis.json "$HOME"/.Atrixd/config/genesis.json
cat "$HOME"/.Atrixd/config/genesis.json | jq '.app_state["evm"]["params"]["evm_denom"]="aAtrix"' > "$HOME"/.Atrixd/config/tmp_genesis.json && mv "$HOME"/.Atrixd/config/tmp_genesis.json "$HOME"/.Atrixd/config/genesis.json
cat "$HOME"/.Atrixd/config/genesis.json | jq '.app_state["inflation"]["params"]["mint_denom"]="aAtrix"' > "$HOME"/.Atrixd/config/tmp_genesis.json && mv "$HOME"/.Atrixd/config/tmp_genesis.json "$HOME"/.Atrixd/config/genesis.json

# set gov proposing && voting period
cat "$HOME"/.Atrixd/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["max_deposit_period"]="30s"' > "$HOME"/.Atrixd/config/tmp_genesis.json && mv "$HOME"/.Atrixd/config/tmp_genesis.json "$HOME"/.Atrixd/config/genesis.json
cat "$HOME"/.Atrixd/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="30s"' > "$HOME"/.Atrixd/config/tmp_genesis.json && mv "$HOME"/.Atrixd/config/tmp_genesis.json "$HOME"/.Atrixd/config/genesis.json

# Set gas limit in genesis
cat "$HOME"/.Atrixd/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="10000000"' > "$HOME"/.Atrixd/config/tmp_genesis.json && mv "$HOME"/.Atrixd/config/tmp_genesis.json "$HOME"/.Atrixd/config/genesis.json

# Set claims start time
node_address=$(Atrixd keys list | grep  "address: " | cut -c12-)
current_date=$(date -u +"%Y-%m-%dT%TZ")
cat "$HOME"/.Atrixd/config/genesis.json | jq -r --arg current_date "$current_date" '.app_state["claims"]["params"]["airdrop_start_time"]=$current_date' > "$HOME"/.Atrixd/config/tmp_genesis.json && mv "$HOME"/.Atrixd/config/tmp_genesis.json "$HOME"/.Atrixd/config/genesis.json

# Set claims records for validator account
amount_to_claim=10000
cat "$HOME"/.Atrixd/config/genesis.json | jq -r --arg node_address "$node_address" --arg amount_to_claim "$amount_to_claim" '.app_state["claims"]["claims_records"]=[{"initial_claimable_amount":$amount_to_claim, "actions_completed":[false, false, false, false],"address":$node_address}]' > "$HOME"/.Atrixd/config/tmp_genesis.json && mv "$HOME"/.Atrixd/config/tmp_genesis.json "$HOME"/.Atrixd/config/genesis.json

# Set claims decay
cat "$HOME"/.Atrixd/config/genesis.json | jq '.app_state["claims"]["params"]["duration_of_decay"]="1000000s"' > "$HOME"/.Atrixd/config/tmp_genesis.json && mv "$HOME"/.Atrixd/config/tmp_genesis.json "$HOME"/.Atrixd/config/genesis.json
cat "$HOME"/.Atrixd/config/genesis.json | jq '.app_state["claims"]["params"]["duration_until_decay"]="100000s"' > "$HOME"/.Atrixd/config/tmp_genesis.json && mv "$HOME"/.Atrixd/config/tmp_genesis.json "$HOME"/.Atrixd/config/genesis.json

# Claim module account:
# 0xA61808Fe40fEb8B3433778BBC2ecECCAA47c8c47 || Atrix15cvq3ljql6utxseh0zau9m8ve2j8erz89m5wkz
cat "$HOME"/.Atrixd/config/genesis.json | jq -r --arg amount_to_claim "$amount_to_claim" '.app_state["bank"]["balances"] += [{"address":"Atrix15cvq3ljql6utxseh0zau9m8ve2j8erz89m5wkz","coins":[{"denom":"aAtrix", "amount":$amount_to_claim}]}]' > "$HOME"/.Atrixd/config/tmp_genesis.json && mv "$HOME"/.Atrixd/config/tmp_genesis.json "$HOME"/.Atrixd/config/genesis.json

# disable produce empty block
if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' 's/create_empty_blocks = true/create_empty_blocks = false/g' "$HOME"/.Atrixd/config/config.toml
  else
    sed -i 's/create_empty_blocks = true/create_empty_blocks = false/g' "$HOME"/.Atrixd/config/config.toml
fi

if [[ $1 == "pending" ]]; then
  if [[ "$OSTYPE" == "darwin"* ]]; then
      sed -i '' 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' "$HOME"/.Atrixd/config/config.toml
      sed -i '' 's/timeout_propose = "3s"/timeout_propose = "30s"/g' "$HOME"/.Atrixd/config/config.toml
      sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' "$HOME"/.Atrixd/config/config.toml
      sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' "$HOME"/.Atrixd/config/config.toml
      sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' "$HOME"/.Atrixd/config/config.toml
      sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' "$HOME"/.Atrixd/config/config.toml
      sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' "$HOME"/.Atrixd/config/config.toml
      sed -i '' 's/timeout_commit = "5s"/timeout_commit = "150s"/g' "$HOME"/.Atrixd/config/config.toml
      sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' "$HOME"/.Atrixd/config/config.toml
  else
      sed -i 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' "$HOME"/.Atrixd/config/config.toml
      sed -i 's/timeout_propose = "3s"/timeout_propose = "30s"/g' "$HOME"/.Atrixd/config/config.toml
      sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' "$HOME"/.Atrixd/config/config.toml
      sed -i 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' "$HOME"/.Atrixd/config/config.toml
      sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' "$HOME"/.Atrixd/config/config.toml
      sed -i 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' "$HOME"/.Atrixd/config/config.toml
      sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' "$HOME"/.Atrixd/config/config.toml
      sed -i 's/timeout_commit = "5s"/timeout_commit = "150s"/g' "$HOME"/.Atrixd/config/config.toml
      sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' "$HOME"/.Atrixd/config/config.toml
  fi
fi

# Allocate genesis accounts (cosmos formatted addresses)
Atrixd add-genesis-account $KEY 100000000000000000000000000aAtrix --keyring-backend $KEYRING

# Update total supply with claim values
# Bc is required to add this big numbers
# total_supply=$(bc <<< "$amount_to_claim+$validators_supply")
total_supply=100000000000000000000010000
cat "$HOME"/.Atrixd/config/genesis.json | jq -r --arg total_supply "$total_supply" '.app_state["bank"]["supply"][0]["amount"]=$total_supply' > "$HOME"/.Atrixd/config/tmp_genesis.json && mv "$HOME"/.Atrixd/config/tmp_genesis.json "$HOME"/.Atrixd/config/genesis.json

# Sign genesis transaction
Atrixd gentx $KEY 1000000000000000000000aAtrix --keyring-backend $KEYRING --chain-id "$CHAINID"
## In case you want to create multiple validators at genesis
## 1. Back to `Atrixd keys add` step, init more keys
## 2. Back to `Atrixd add-genesis-account` step, add balance for those
## 3. Clone this ~/.Atrixd home directory into some others, let's say `~/.clonedAtrixd`
## 4. Run `gentx` in each of those folders
## 5. Copy the `gentx-*` folders under `~/.clonedAtrixd/config/gentx/` folders into the original `~/.Atrixd/config/gentx`

# Collect genesis tx
Atrixd collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
Atrixd validate-genesis

if [[ $1 == "pending" ]]; then
  echo "pending mode is on, please wait for the first block committed."
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
Atrixd start --pruning=nothing "$TRACE" --log_level $LOGLEVEL --minimum-gas-prices=0.0001aAtrix --json-rpc.api eth,txpool,personal,net,debug,web3
