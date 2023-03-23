package fvm

import (
	"collective-go-sdk/config"
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
)

func TestSymbol(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadConfig("../config")

	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	symbol, err := client.Staking.Symbol(&bind.CallOpts{})
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, "clFIL", symbol)
}

func TestName(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadConfig("../config")

	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	symbol, err := client.Staking.Name(&bind.CallOpts{})
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, "Collective Staked FIL", symbol)
}
