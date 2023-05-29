package fvm

import (
	"collective-go-sdk/config"
	"collective-go-sdk/keystore"
	"context"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLotusClient(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	chainId, err := client.Client.ChainID(ctx)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, big.NewInt(3141), chainId)
	assert.Equal(t, config.RPCConfig[DefaultNetwork].Address, client.Host)

	config.DefaultNetwork = "calibration"

	client, err = NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, client.Collateral.Address.String(), "0x000000000000000000000000000000000000dead")
	assert.Equal(t, client.Registry.Address.String(), "0x000000000000000000000000000000000000dead")
	assert.Equal(t, client.Staking.Address.String(), "0x000000000000000000000000000000000000dead")
}
