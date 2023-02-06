package fvm

import (
	"collective-go-sdk/config"
	"collective-go-sdk/contracts/LiquidStaking"
	"collective-go-sdk/contracts/PledgeOracle"
	"collective-go-sdk/contracts/StorageProviderCollateral"
	"collective-go-sdk/contracts/StorageProviderRegistry"

	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/umbracle/ethgo/contract"
)

type LotusClient struct {
	host       string
	client     *ethclient.Client
	signer     *bind.TransactOpts
	address    *common.Address
	registry   *StorageProviderRegistry.StorageProviderRegistry
	staking    *LiquidStaking.LiquidStaking
	oracle     *PledgeOracle.PledgeOracle
	collateral *StorageProviderCollateral.StorageProviderCollateral
}

type Contract struct {
	contract.Contract
}

func NewLotusClient(cfg *config.Config) (*LotusClient, error) {
	c := &LotusClient{}

	client, err := ethclient.Dial(cfg.RPCAddress)

	if err != nil {
		return nil, err
	}

	c.client = client
	c.host = cfg.RPCAddress

	key, err := c.getPrivateKey(cfg)
	if err != nil {
		return nil, err
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(int64(cfg.ChainID)))
	if err != nil {
		return nil, err
	}

	c.signer = transactor
	c.address = &c.signer.From

	registry, err := StorageProviderRegistry.NewStorageProviderRegistry(common.HexToAddress(cfg.StorageProviderRegistry), client)
	if err != nil {
		return nil, err
	}

	oracle, err := PledgeOracle.NewPledgeOracle(common.HexToAddress(cfg.PledgeOracle), client)
	if err != nil {
		return nil, err
	}

	collateral, err := StorageProviderCollateral.NewStorageProviderCollateral(common.HexToAddress(cfg.StorageProviderCollateral), client)
	if err != nil {
		return nil, err
	}

	staking, err := LiquidStaking.NewLiquidStaking(common.HexToAddress(cfg.LiquidStaking), client)
	if err != nil {
		return nil, err
	}

	c.registry = registry
	c.oracle = oracle
	c.collateral = collateral
	c.staking = staking

	return c, nil
}

func (c *LotusClient) getPrivateKey(cfg *config.Config) (*ecdsa.PrivateKey, error) {
	if cfg.PrivateKey != "" {
		key, err := crypto.HexToECDSA(cfg.PrivateKey)
		return key, err
	}

	if cfg.MnemonicPhrase != "" {
		wallet, err := hdwallet.NewFromMnemonic(cfg.MnemonicPhrase)
		if err != nil {
			return nil, err
		}
		path := hdwallet.DefaultBaseDerivationPath

		if cfg.HDDerivationPath != "" {
			parsedPath, err := hdwallet.ParseDerivationPath(cfg.HDDerivationPath)
			if err != nil {
				return nil, err
			}
			path = parsedPath
		}

		account, err := wallet.Derive(path, true)
		if err != nil {
			return nil, err
		}
		key, err := wallet.PrivateKey(account)
		if err != nil {
			return nil, err
		}
		return key, nil
	}

	return nil, fmt.Errorf("private key or mnemonic phrase isn't specified")
}
