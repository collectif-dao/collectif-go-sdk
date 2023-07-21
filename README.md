# collectif-go-sdk

Collectif-Go-SDK is the main client for Filecoin Storage Providers to interact with Collective Liquid Staking protocol it wraps around of all the complexities of joining Collectif protocol, pledging capital from the liquid staking pool into miner's Initial Pledge, and manage SPs collateral.

## collectif-client

Collectif client is an automated worker for pledging on a daily basis. SP's can run worker to automate the process which simplifies an integration with Collectif DAO liquid staking protocol.

The main scenario is auto pledging, which is executed on a daily basis (with 24 hour delay). If there is not enough funds in the collateral to cover the slashing risks for stakers worker would automatically deposit FIL into the `StorageProviderCollateral` smart contract. The deposit amount is determined during the execution and wouldn't cover the next day pledge allocation from the pool. It's recommended to deposit some FIL funds upfront to make sure that SP is able to get allocation without a need to deposit FIL in the collateral on a daily basis.

## Setup

Please make sure to install the following before working with codebase:

[Go-lang](https://go.dev/doc/install)

## Build

To build an sdk please run: `make`
This command would build both collectif-client and collectif-go-sdk.

To build individual binaries you can run: `make build-sdk` or `make build-client`. Both packages rely on the [filecoin-ffi](https://github.com/filecoin-project/filecoin-ffi/) to work with owner id private keys. Make sure to run `make filecoin-ffi` beforehand otherwise non of the packages would pass build.

If you're using a Mac with Apple Silicon (M1/M2 processors) it's likely that you may face an error with `ld: library not found for -lhwloc` in this case run `export LIBRARY_PATH=/opt/homebrew/lib` in your terminal and try to build packages again

## Modifying config

Both packages are using config to find where your private key is stored. You can provide an existing location of your lotus keystore in the `config.yaml` file. Also config provides a default Filecoin network and a list of miners for automatic pledging. Ideally you'll need to add your miner actor ID into after the registration.

In case if you don't need to provide a location of the existing lotus keystore you can simply import your private key into the SDK using CLI commands.

## Importing lotus wallet

To import your existing lotus wallet you need to build an SDK, and run `./collectif-go-sdk wallet import -p <PRIVATE_KEY>`. Make sure to path a private key string from your existing wallet that is an owner of your miner actors that you want to use with liquid staking protocol.

After your wallet import this key would be used as your default wallet.

## Using packages

### Using collectif-go-sdk CLI

To start using SDK via CLI make sure to build a package and run `./collectif-go-sdk help` in your terminal. You will see a list of commands available for you.

```
Collective-Go-SDK is the main client for Filecoin Storage Providers
        to interact with Collective Liquid Staking protocol it wraps
        around of all the complexities of joining Collectif protocol, pledging capital from the liquid staking pool
        into miner's Initial Pledge, and manage SPs collateral.

Usage:
  collective-go-sdk [command]

Available Commands:
  allocation         Allocation section for Storage Provider
  change-beneficiary Change beneficiary address in the Collectif DAO protocol
  collateral         This is the section for interacting with StorageProviderCollateral contract
  completion         Generate the autocompletion script for the specified shell
  help               Help about any command
  register           Register Storage Provider in the Collectif DAO protocol
  restaking          Restaking section of the Storage Provider
  staking            Interact with Liquid Staking contract to pledge FIL
  wallet             Interact with Filecoin/Ethereum wallet
```

### Registering in Collectif DAO

To register your miner actor in Collectif DAO you'll need to run `./collectif-go-sdk register -m <MINER_ID> -l <TOTAL_ALLOCATION> -d <DAILY_ALLOCATION>` where:

- `<MINER_ID>` is your miner actor id in the filecoin network
- `<TOTAL_ALLOCATION>` is an amount of FIL (not attoFIL) that your miner actor would need to pledge overtime
- `<DAILY_ALLOCATION>` is a FIL (not attoFIL) amount that you expect to utilize in 1 day. This amount is dependent on your sealing speed

Registration process happens on a per miner actor basis. But all miners owned by you are within the same collateral bucket. This means that you can register several miners and manage collateral for them in aggregate basis.

### Depositing FIL collateral

In order to deposit FIL as collateral into the liquid staking protocol you can run `./collectif-go-sdk collateral deposit -a <DEPOSIT_AMOUNT>` where your `<DEPOSIT_AMOUNT>` is provided in FIL. Make sure to provide FIL value not attoFIL.

### Integration with lotus node

The Collectif DAO Go-lang SDK provides an easy access point to integrate into existing lotus node. The main point of integration is at the end of sealing process to acquire initial pledge collateral from the network. To integrate SDK into your lotus/venus node you can use the following example:

```
package main

import (
	"collective-go-sdk/keystore"
	"collective-go-sdk/sdk"
	"collective-go-sdk/utils"
	"context"
)

func main() {
	ctx := context.Background()

	ksType := keystore.FSKeyStore
	configPath := "./"

	sdk, err := sdk.NewCollectifSDK(ctx, ksType, configPath)
	if err != nil {
		panic(err)
	}

	pledgeAmount := utils.GetAttoFilFromFIL(1000)

	sdk.Client.Pledge(ctx, pledgeAmount, true)
}
```

## Running collectif-client

After registering your miner actors in the Collectif DAO liquid staking protocol you can add your miner ID into the config file `config.yaml` and start worker by running `./collectif-client`

## Documentation

You can find more details in our docs at Collectif DAO [documentation](https://docs.collectif.finance).

## License

MIT
