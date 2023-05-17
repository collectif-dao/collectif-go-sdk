package fvm

import (
	"collective-go-sdk/config"
	"collective-go-sdk/contracts/LiquidStaking"
	"collective-go-sdk/contracts/StorageProviderCollateral"
	"collective-go-sdk/contracts/StorageProviderRegistry"
	"collective-go-sdk/keystore"
	"context"
	"path"
	"runtime"

	"github.com/ybbus/jsonrpc/v3"

	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/wallet"
	"github.com/filecoin-project/lotus/chain/wallet/key"
	"github.com/sirupsen/logrus"
	"github.com/umbracle/ethgo/contract"
	"github.com/valyala/fasthttp"
	"golang.org/x/xerrors"
)

type LotusClient struct {
	Host          string
	Client        *ethclient.Client
	rpcClient     *fasthttp.HostClient
	rpcClient2    jsonrpc.RPCClient
	MessageSigner *MessageSigner
	Signer        *bind.TransactOpts
	Address       *address.Address
	Registry      *StorageProviderRegistry.StorageProviderRegistry
	RegistryABI   *bind.MetaData
	Staking       *LiquidStaking.LiquidStaking
	StakingABI    *bind.MetaData
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
	c.Host = cfg.RPCAddress

	c.rpcClient = &fasthttp.HostClient{
		Addr: cfg.RPCAddress,
	}

	rpcClient := jsonrpc.NewClient(cfg.RPCAddress)

	c.rpcClient2 = rpcClient

	ks, err := cache.prepareKeystore(cfg.FSKeyStoreDir)
	if err != nil {
		return nil, err
	}

	w, err := wallet.NewWallet(ks)
	if err != nil {
		return nil, err
	}

	c.MessageSigner = NewMessageSigner(*w)

	addr, err := w.GetDefault()
	if err != nil {
		key, err := key.GenerateKey("secp256k1")
		if err != nil {
			return nil, err
		}

		if err := ks.Put(wallet.KNamePrefix+key.Address.String(), key.KeyInfo); err != nil {
			return nil, xerrors.Errorf("saving to keystore: %w", err)
		}

		err = ks.Put("default", key.KeyInfo)

		c.Address = &key.Address
	}

	c.Address = &addr

	registry, err := StorageProviderRegistry.NewStorageProviderRegistry(common.HexToAddress(cfg.StorageProviderRegistry), client)
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

	c.Collateral = collateral
	c.CollateralABI = StorageProviderCollateral.StorageProviderCollateralMetaData

	c.Staking = staking
	c.StakingABI = LiquidStaking.LiquidStakingMetaData

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return "", fmt.Sprintf("%s:%d:", filename, f.Line)
		},
	})

	logrus.SetLevel(logrus.InfoLevel)

	return c, nil
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
