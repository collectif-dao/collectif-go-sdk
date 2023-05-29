package fvm

import (
	"collective-go-sdk/config"
	"collective-go-sdk/keystore"
	"context"
	"testing"

	"math/big"

	"github.com/stretchr/testify/assert"
)

var minerId = uint64(0100)
var ZERO_BN = big.NewInt(0)
var amount = big.NewInt(100)

func TestDepositCallData(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, keystore.FSKeyStore)

	callData, err := client.calculateCalldata("deposit", client.Collateral.ABI)
	if err != nil {
		assert.Error(t, err)
	}

	res, err := client.Deposit(ctx, amount, false)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, callData, res.Data)
}

func TestWithdrawCallData(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, keystore.FSKeyStore)

	res, err := client.Withdraw(ctx, amount, false)
	if err != nil {
		assert.Error(t, err)
	}

	callData, err := client.calculateCalldata("withdraw", client.Collateral.ABI, amount)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, callData, res.Data)
}

func TestGetCollateral(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadConfig("../")

	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	collateral, err := client.GetCollateral(ctx, minerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, ZERO_BN.String(), collateral.AvailableCollateral.String())
	assert.Equal(t, ZERO_BN.String(), collateral.LockedCollateral.String())
}

func TestGetLockedCollateral(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadConfig("../")

	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	lockedCollateral, err := client.GetLockedCollateral(ctx, minerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, ZERO_BN.String(), lockedCollateral.String())
}

func TestGetAvailableCollateral(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadConfig("../")

	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	availableCollateral, err := client.GetAvailableCollateral(ctx, minerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, ZERO_BN.String(), availableCollateral.String())
}

func TestIsActiveSlashing(t *testing.T) {
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
	status, err := client.IsActiveSlashing(ctx, ownerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, status, false)
}

func TestGetTotalSlashing(t *testing.T) {
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
	slashing, err := client.GetTotalSlashing(ctx, ownerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, ZERO_BN.String(), slashing.String())
}
