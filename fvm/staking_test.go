package fvm

import (
	"collective-go-sdk/config"
	"collective-go-sdk/keystore"
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
)

func TestSymbol(t *testing.T) {
	ctx := context.Background()

	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	symbol, err := client.Staking.Contract.Symbol(&bind.CallOpts{})
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, "clFIL", symbol)
}

func TestName(t *testing.T) {
	ctx := context.Background()

	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	symbol, err := client.Staking.Contract.Name(&bind.CallOpts{})
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, "Collectif Staked FIL", symbol)
}

func TestTotalAssets(t *testing.T) {
	ctx := context.Background()

	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	assets, err := client.TotalAssets(ctx)
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotEqual(t, ZERO_BN.String(), assets.String())
}

func TestTotalFILAvailable(t *testing.T) {
	ctx := context.Background()

	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	fil, err := client.TotalFilAvailable(ctx)
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotEqual(t, ZERO_BN.String(), fil.String())
}

func TestTotalFILPledged(t *testing.T) {
	ctx := context.Background()

	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	fil, err := client.TotalFilPledged(ctx)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, ZERO_BN.String(), fil.String())
}

func TestTotalFees(t *testing.T) {
	ctx := context.Background()

	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	ownerId := getOwnerId(t, ctx, client)
	fil, err := client.TotalFees(ctx, ownerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, ZERO_BN.String(), fil.String())
}

func TestPledge(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, keystore.FSKeyStore)

	callData, err := client.calculateCalldata("pledge", client.Staking.ABI, amount, minerId)
	if err != nil {
		assert.Error(t, err)
	}

	res, err := client.Pledge(ctx, amount, minerId, false)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, callData, res.Data)
}

func TestStake(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, keystore.FSKeyStore)

	callData, err := client.calculateCalldata("stake", client.Staking.ABI)
	if err != nil {
		assert.Error(t, err)
	}

	res, err := client.Stake(ctx, amount, false)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, callData, res.Data)
}
