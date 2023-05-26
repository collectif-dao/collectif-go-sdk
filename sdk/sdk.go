package sdk

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"collective-go-sdk/keystore"
	"context"

	"github.com/filecoin-project/go-address"
)

type CollectifSDK struct {
	Host    string
	Address *address.Address
	Client  *fvm.LotusClient
}

func NewCollectifSDK(ctx context.Context, network string, cache keystore.CacheType, path string) (*CollectifSDK, error) {
	s := CollectifSDK{}
	var cfg *config.Config
	var err error

	if path == "" {
		cfg, err = config.LoadConfig("../")
	} else {
		cfg, err = config.LoadConfig(path)
	}

	if err != nil {
		return nil, err
	}

	client, err := fvm.NewLotusClient(ctx, cfg, network, cache)
	if err != nil {
		return nil, err
	}

	s.Client = client

	return &s, nil
}
