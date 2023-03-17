# collectif-go-sdk

Collective-Go-SDK is the main client for Filecoin Storage Providers to interact with Collective Liquid Staking protocol it wraps around of all the complexities of joining Collectif protocol, pledging capital from the liquid staking pool into miner's Initial Pledge, and manage SPs collateral.

## Setup

Please make sure to install the following before working with codebase:

[Go-lang](https://go.dev/doc/install)

## Build

To build an sdk please run: `make build`

## Documentation

You can find Go-lang SDK documentation over at Collectif DAO [documentation website](https://docs.collectif.finance).

## Using the SDK

### Using via CLI

Go-lang SDK provides an CLI tool to interact with Collectif DAO protocol on FVM. To learn more about available commands please run `collective-go-sdk help` on your terminal

### Integration with lotus node

The Collectif DAO Go-lang SDK provides an easy access point to integrate into existing lotus node. The main point of integration is at the end of sealing process to acquire initial pledge collateral from the network.
