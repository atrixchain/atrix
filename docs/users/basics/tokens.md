<!--
order: 2
-->

# Tokens

Learn about the the different types of tokens available in Atrix. {synopsis}

## Introduction

Atrix is a Cosmos-based chain with full Ethereum Virtual Machine (EVM) support. Because of this [architecture](./../technical_concepts/architecture.md), tokens and assets in the network may come from different independent sources.

## The Atrix Token

The denomination used for staking, governance and gas consumption on the EVM is the Atrix. The Atrix provides the utility of: securing the Proof-of-Stake chain, token used for governance proposals, distribution of fees to validator and users, and as a mean of gas for running smart contracts on the EVM.

Atrix uses [Atto](https://en.wikipedia.org/wiki/Atto-) Atrix as the base denomination to maintain parity with Ethereum.

1 Atrix = 10<sup>18</sup> aAtrix

This matches Ethereum denomination of:

1 ETH = 10<sup>18</sup> wei

## Cosmos Coins

Accounts can own Cosmos coins in their balance, which are used for operations with other Cosmos and transactions. Examples of these are using the coins for staking, IBC transfers, governance deposits and EVM.

## EVM Tokens

Atrix is compatible with ERC20 tokens and other non-fungible token standards (EIP721, EIP1155)
that are natively supported by the EVM.
