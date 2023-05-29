package fvm

import (
	"collective-go-sdk/config"
	"collective-go-sdk/contracts/LiquidStaking"
	"collective-go-sdk/contracts/StorageProviderCollateral"
	"collective-go-sdk/contracts/StorageProviderRegistry"
	"collective-go-sdk/keystore"
	ms "collective-go-sdk/messagesigner"
	"collective-go-sdk/rpc"
	"context"
	"path"
	"runtime"

	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/filecoin-project/lotus/chain/wallet"
	"github.com/filecoin-project/lotus/chain/wallet/key"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

var DefaultNetwork string

type LotusClient struct {
	Host       string
	Client     *ethclient.Client
	Signer     *bind.TransactOpts
	Address    *address.Address
	EthAddress *common.Address

	RPCClient     *rpc.RPCClient
	MessageSigner *ms.MessageSigner

	Registry   *Registry
	Collateral *Collateral
	Staking    *Staking
}

type Registry struct {
	Contract      *StorageProviderRegistry.StorageProviderRegistry
	Address       ethtypes.EthAddress
	NativeAddress address.Address
	ABI           *abi.ABI
}

type Collateral struct {
	Contract      *StorageProviderCollateral.StorageProviderCollateral
	Address       ethtypes.EthAddress
	NativeAddress address.Address
	ABI           *abi.ABI
}

type Staking struct {
	Contract      *LiquidStaking.LiquidStaking
	Address       ethtypes.EthAddress
	NativeAddress address.Address
	ABI           *abi.ABI
}

func NewLotusClient(ctx context.Context, cfg *config.Config, cache keystore.CacheType) (*LotusClient, error) {
	c := &LotusClient{}
	DefaultNetwork = cfg.DefaultNetwork
	rpcAddr := cfg.RPCConfig[DefaultNetwork].Address

	client, err := ethclient.Dial(rpcAddr)
	if err != nil {
		return nil, err
	}

	c.Client = client
	c.Host = rpcAddr

	c.RPCClient = rpc.NewRPCClient(rpcAddr, cfg.RPCConfig[DefaultNetwork].APIToken)

	ks, err := cache.PrepareKeystore(cfg.FSKeyStoreDir)
	if err != nil {
		return nil, err
	}

	w, err := wallet.NewWallet(ks)
	if err != nil {
		return nil, err
	}

	c.MessageSigner = ms.NewMessageSigner(*w, c.RPCClient)

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
	} else {
		c.Address = &addr
	}

	registry, err := initRegistry(cfg.Addresses[DefaultNetwork].StorageProviderRegistry, client)
	if err != nil {
		return nil, err
	}
	c.Registry = registry

	collateral, err := initCollateral(cfg.Addresses[DefaultNetwork].StorageProviderCollateral, client)
	if err != nil {
		return nil, err
	}
	c.Collateral = collateral

	staking, err := initStaking(cfg.Addresses[DefaultNetwork].LiquidStaking, client)
	if err != nil {
		return nil, err
	}
	c.Staking = staking

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

func initRegistry(addr string, client *ethclient.Client) (*Registry, error) {
	c, err := StorageProviderRegistry.NewStorageProviderRegistry(common.HexToAddress(addr), client)
	if err != nil {
		return nil, err
	}

	ethAddr, err := ethtypes.ParseEthAddress(addr)
	if err != nil {
		return nil, err
	}

	fAddr, err := ethAddr.ToFilecoinAddress()
	if err != nil {
		return nil, err
	}

	abi, err := StorageProviderRegistry.StorageProviderRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return &Registry{
		Contract:      c,
		Address:       ethAddr,
		NativeAddress: fAddr,
		ABI:           abi,
	}, nil
}

func initCollateral(addr string, client *ethclient.Client) (*Collateral, error) {
	c, err := StorageProviderCollateral.NewStorageProviderCollateral(common.HexToAddress(addr), client)
	if err != nil {
		return nil, err
	}

	ethAddr, err := ethtypes.ParseEthAddress(addr)
	if err != nil {
		return nil, err
	}

	fAddr, err := ethAddr.ToFilecoinAddress()
	if err != nil {
		return nil, err
	}

	abi, err := StorageProviderCollateral.StorageProviderCollateralMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return &Collateral{
		Contract:      c,
		Address:       ethAddr,
		NativeAddress: fAddr,
		ABI:           abi,
	}, nil
}

func initStaking(addr string, client *ethclient.Client) (*Staking, error) {
	sc, err := LiquidStaking.NewLiquidStaking(common.HexToAddress(addr), client)
	if err != nil {
		return nil, err
	}

	ethAddr, err := ethtypes.ParseEthAddress(addr)
	if err != nil {
		return nil, err
	}

	fAddr, err := ethAddr.ToFilecoinAddress()
	if err != nil {
		return nil, err
	}

	abi, err := LiquidStaking.LiquidStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return &Staking{
		Contract:      sc,
		Address:       ethAddr,
		NativeAddress: fAddr,
		ABI:           abi,
	}, nil
}
