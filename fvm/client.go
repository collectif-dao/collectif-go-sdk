package fvm

import (
	"collective-go-sdk/config"
	"collective-go-sdk/contracts/LiquidStaking"
	"collective-go-sdk/contracts/PledgeOracle"
	"collective-go-sdk/contracts/StorageProviderCollateral"
	"collective-go-sdk/contracts/StorageProviderRegistry"
	"collective-go-sdk/keystore"
	"context"

	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filecoin-project/lotus/chain/wallet"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/umbracle/ethgo/contract"
	"github.com/valyala/fasthttp"
)

type LotusClient struct {
	Host          string
	Client        *ethclient.Client
	rpcClient     *fasthttp.HostClient
	MessageSigner *MessageSigner
	Signer        *bind.TransactOpts
	Address       *common.Address
	Registry      *StorageProviderRegistry.StorageProviderRegistry
	RegistryABI   *bind.MetaData
	Staking       *LiquidStaking.LiquidStaking
	StakingABI    *bind.MetaData
	Oracle        *PledgeOracle.PledgeOracle
	Collateral    *StorageProviderCollateral.StorageProviderCollateral
	CollateralABI *bind.MetaData
}

type Contract struct {
	contract.Contract
}

type CacheType string

const (
	MemoryKeyStore CacheType = "memory"
	FSKeyStore     CacheType = "filesystem"
)

func NewLotusClient(ctx context.Context, cfg *config.Config, cache CacheType) (*LotusClient, error) {
	c := &LotusClient{}

	client, err := ethclient.Dial(cfg.RPCAddress)

	if err != nil {
		return nil, err
	}

	c.Client = client
	c.rpcClient = &fasthttp.HostClient{
		Addr: cfg.RPCAddress,
	}
	c.Host = cfg.RPCAddress

	key, err := c.getPrivateKey(cfg)
	if err != nil {
		return nil, err
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(int64(cfg.ChainID)))
	if err != nil {
		return nil, err
	}

	c.Signer = transactor
	c.Address = &c.Signer.From

	ks, err := cache.prepareKeystore(cfg.FSKeyStoreDir)
	if err != nil {
		return nil, err
	}

	wallet, err := wallet.NewWallet(ks)
	if err != nil {
		return nil, err
	}

	c.MessageSigner = NewMessageSigner(*wallet)

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

	c.Registry = registry
	c.RegistryABI = StorageProviderRegistry.StorageProviderRegistryMetaData

	c.Oracle = oracle

	c.Collateral = collateral
	c.CollateralABI = StorageProviderCollateral.StorageProviderCollateralMetaData

	c.Staking = staking
	c.StakingABI = LiquidStaking.LiquidStakingMetaData

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

func (cache CacheType) prepareKeystore(dir string) (keystore.KeyStore, error) {
	var ks keystore.KeyStore
	var err error

	switch cache {
	case MemoryKeyStore:
		ks = keystore.NewMemKeystore()
	case FSKeyStore:
		ks, err = keystore.NewFSKeystore(dir)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("keystore type isn't specified")
	}

	return ks, nil
}
