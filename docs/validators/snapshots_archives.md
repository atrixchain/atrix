<!--
order: 6
-->

# Snapshots & Archive Nodes

Quickly sync your node with Atrix using a snapshot or serve queries for prev versions using archive nodes {synopsis}

## List of Snapshots and Archives

Below is a list of publicly available snapshots that you can use to sync with the Atrix mainnet and
archived [9001-1 mainnet](https://github.com/Atrix/mainnet/tree/main/Atrix_9001-1):

<!-- markdown-link-check-disable -->

### Snapshots

| Name        | URL                                                                     |
| -------------|------------------------------------------------------------------------ |
| `Staketab`   | [github.com/staketab/nginx-cosmos-snap](https://github.com/staketab/nginx-cosmos-snap/blob/main/docs/Atrix.md) |
| `Polkachu`   | [polkachu.com](https://www.polkachu.com/tendermint_snapshots/Atrix)                   |
| `Nodes Guru` | [snapshots.nodes.guru/Atrix_9001-2/](snapshots.nodes.guru/Atrix_9001-2/)                   |
| `Notional`   | [mainnet/pruned/Atrix_9001-2(pebbledb)](https://snapshot.notional.ventures/Atrix/) <br> [mainnet/archive/Atrix_9001-2(pebbledb)](https://snapshot.notional.ventures/Atrix-archive/) <br> [testnet/archive/Atrix_9000-4(pebbledb)](https://snapshot.notional.ventures/Atrix-testnet-archive/)                   |

### Archives
<!-- markdown-link-check-disable -->

| Name           | URL                                                                             |
| ---------------|---------------------------------------------------------------------------------|
| `Nodes Guru`   | [snapshots.nodes.guru/Atrix_9001-1](https://snapshots.nodes.guru/Atrix_9001-1/)                                    |
| `Polkachu`     | [polkachu.com/tendermint_snapshots/Atrix](https://www.polkachu.com/tendermint_snapshots/Atrix)                           |
| `Forbole`      | [bigdipper.live/Atrix_9001-1](https://s3.bigdipper.live.eu-central-1.linodeobjects.com/Atrix_9001-1.tar.lz4) |

To access snapshots and archives, follow the process below (this code snippet is to access a snapshot of the current network, `Atrix_9001-2`, from Nodes Guru):

```bash
cd $HOME/.Atrixd/data
wget https://snapshots.nodes.guru/Atrix_9001-2/Atrix_9001-2-410819.tar
tar xf Atrix_9001-2-410819.tar
```

### PebbleDB

To use PebbleDB instead of GoLevelDB when using snapshots from Notional:

Build:

```bash
go mod edit -replace github.com/tendermint/tm-db=github.com/baabeetaa/tm-db@pebble
go mod tidy
go install -tags pebbledb -ldflags "-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb" ./...
```

Download snapshot:

```bash
cd $HOME/.Atrixd/
URL_SNAPSHOT="https://snapshot.notional.ventures/Atrix/data_20221024_193254.tar.gz"
wget -O - "$URL_SNAPSHOT" |tar -xzf -
```

Start:

Set `db_backend = "pebbledb"` in `config.toml` or start with `--db_backend=pebbledb`

```bash
Atrixd start --db_backend=pebbledb
```

**Note**: use this [workaround](https://github.com/notional-labs/cosmosia/blob/main/docs/pebbledb.md) when upgrading a node running PebbleDB.
