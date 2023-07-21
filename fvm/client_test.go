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

	assert.Equal(t, big.NewInt(314159), chainId)
	assert.Equal(t, config.RPCConfig[DefaultNetwork].Address, client.Host)

	config.DefaultNetwork = "calibration"

	client, err = NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, client.Collateral.Address.String(), "0x98a79c415af7b4c0c2c8fb440796d02652abdf87")
	assert.Equal(t, client.Registry.Address.String(), "0xcc6c40da237f2311e6d8e8e7832b1f76aa6115e4")
	assert.Equal(t, client.Staking.Address.String(), "0x19aab7dd96e9eedf9e232fe56d1736f53205834a")
}
