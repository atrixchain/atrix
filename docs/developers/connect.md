<!--
order: 2
-->

# Quick Connect

Quickly connect your app or client to Atrix services {synopsis}

## Public Available Endpoints

Below is a list of publicly available endpoints that you can use to connect to the Atrix mainnet and
public testnets:

::: tip
You can also use [chainlist.org](https://chainlist.org/) to add the node directly to [Metamask](./../users/wallets/metamask.md#automatic-import).
:::

<!-- markdown-link-check-disable -->
### Mainnet

| Address                                       | Category               | Maintainer                              |
| --------------------------------------------- | ---------------------- | --------------------------------------- |
| `https://grpc.bd.Atrix.org:9090`              | `Cosmos` `gRPC`        | [Blockdaemon](https://blockdaemon.com/) |
| `https://rest.bd.Atrix.org:1317`              | `Cosmos` `REST`        | [Blockdaemon](https://blockdaemon.com/) |
| `https://tendermint.bd.Atrix.org:26657`       | `Tendermint` `RPC`     | [Blockdaemon](https://blockdaemon.com/) |
| `https://eth.bd.Atrix.org:8545`               | `Ethereum` `JSON-RPC`  | [Blockdaemon](https://blockdaemon.com/) |
| `wss://eth.bd.Atrix.org:8546`                 | `Ethereum` `Websocket` | [Blockdaemon](https://blockdaemon.com/) |
| `https://Atrix-json-rpc.stakely.io`           | `Ethereum` `JSON-RPC`  | [Stakely](https://stakely.io/)          |
| `https://Atrix-rpc.stakely.io`                | `Cosmos` `RPC`         | [Stakely](https://stakely.io/)          |
| `https://Atrix-lcd.stakely.io`                | `Cosmos` `REST`        | [Stakely](https://stakely.io/)          |
| `https://jsonrpc-Atrix-ia.cosmosia.notional.ventures/` | `Ethereum` `JSON-RPC`  | [Notional](https://notional.ventures/)  |
| `https://rpc-Atrix-ia.cosmosia.notional.ventures:443`  | `Tendermint` `RPC`     | [Notional](https://notional.ventures/)  |
| `https://grpc-Atrix-ia.cosmosia.notional.ventures:443` | `Tendermint` `gRPC`    | [Notional](https://notional.ventures/)  |
| `https://api-Atrix-ia.cosmosia.notional.ventures:443`  | `Tendermint` `RPC`     | [Notional](https://notional.ventures/)  |
| `https://rpc.Atrix.nodestake.top`             | `Tendermint` `RPC`     | [NodeStake](https://nodestake.top/)     |
| `https://grpc.Atrix.nodestake.top`            | `Cosmos` `gRPC`        | [NodeStake](https://nodestake.top/)     |
| `https://api.Atrix.nodestake.top`             | `Cosmos` `REST`        | [NodeStake](https://nodestake.top/)     |
| `https://jsonrpc.Atrix.nodestake.top`         | `Ethereum` `JSON-RPC`  | [NodeStake](https://nodestake.top/)     |
| `https://rpc.Atrix.chaintools.tech/`          | `Tendermint` `RPC`     | [ChainTools](https://chaintools.tech/)  |
| `https://Atrix.grpcui.chaintools.host`        | `Cosmos` `gRPC`        | [ChainTools](https://chaintools.tech/)  |
| `https://api.Atrix.chaintools.tech/`          | `Tendermint` `API`     | [ChainTools](https://chaintools.tech/)  |
| `https://rpc.Atrix.silknodes.io`              | `Tendermint` `RPC`     | [Silk Nodes](https://silknodes.io/)     |
| `https://grpc.Atrix.silknodes.io`             | `Cosmos` `gRPC`        | [Silk Nodes](https://silknodes.io/)     |
| `https://api.Atrix.silknodes.io`              | `Cosmos` `REST`        | [Silk Nodes](https://silknodes.io/)     |
| `https://Atrix-mainnet.public.blastapi.io`    | `Ethereum` `JSON-RPC`  | [BLAST](https://blastapi.io/)           |
| `wss://Atrix-mainnet.public.blastapi.io`      | `Ethereum` `Websocket` | [BLAST](https://blastapi.io/)           |
| `https://Atrix-evm.publicnode.com`            | `Ethereum` `JSON-RPC`  | [PublicNode (by Allnodes)](https://Atrix.publicnode.com/) |
| `https://Atrix-rpc.publicnode.com`            | `Tendermint` `RPC`     | [PublicNode (by Allnodes)](https://Atrix.publicnode.com/) |
| `https://Atrix-rest.publicnode.com`           | `Cosmos` `REST`        | [PublicNode (by Allnodes)](https://Atrix.publicnode.com/) |
| `https://Atrix-api.validatrium.club`           | `Tendermint` `API`        | [Validatrium](https://validatrium.com/) |
| `https://Atrix-rpc.validatrium.club`           | `Tendermint` `RPC`        | [Validatrium](https://validatrium.com/) |
| `https://Atrix-rpc.gateway.pokt.network`      | `Ethereum` `JSON-RPC`  | [PocketNetwork](https://www.pokt.network/)  |

### Testnet
<!-- markdown-link-check-disable -->

| Address                                      | Category               | Maintainer                              |
| -------------------------------------------- | ---------------------- | --------------------------------------- |
| `https://grpc.bd.Atrix.dev:9090`             | `Cosmos` `gRPC`        | [Blockdaemon](https://blockdaemon.com/) |
| `https://rest.bd.Atrix.dev:1317`             | `Cosmos` `REST`        | [Blockdaemon](https://blockdaemon.com/) |
| `https://tendermint.bd.Atrix.dev:26657`      | `Tendermint` `RPC`     | [Blockdaemon](https://blockdaemon.com/) |
| `https://eth.bd.Atrix.dev:8545`              | `Ethereum` `JSON-RPC`  | [Blockdaemon](https://blockdaemon.com/) |
| `wss://eth.bd.Atrix.dev:8546`                | `Ethereum` `Websocket` | [Blockdaemon](https://blockdaemon.com/) |
| `https://Atrix-testnet-rpc.polkachu.com:443` | `Tendermint` `RPC`     | [Polkachu](https://polkachu.com)        |
| `https://rpc-t.Atrix.nodestake.top`          | `Tendermint` `RPC`     | [NodeStake](https://nodestake.top/)     |
| `https://grpc-t.Atrix.nodestake.top`         | `Cosmos` `gRPC`        | [NodeStake](https://nodestake.top/)     |
| `https://api-t.Atrix.nodestake.top`          | `Cosmos` `REST`        | [NodeStake](https://nodestake.top/)     |
| `https://jsonrpc-t.Atrix.nodestake.top`      | `Ethereum` `JSON-RPC`  | [NodeStake](https://nodestake.top/)     |
| `https://Atrix-testnet-rpc.qubelabs.io`      | `Tendermint` `RPC`     | [Qubelabs](https://qubelabs.io/)        |
| `https://Atrix-testnet-lcd.qubelabs.io`      | `Cosmos` `REST`        | [Qubelabs](https://qubelabs.io/)        |
| `https://Atrix-testnet-grpc.qubelabs.io`     | `Cosmos` `gRPC`        | [Qubelabs](https://qubelabs.io/)        |
