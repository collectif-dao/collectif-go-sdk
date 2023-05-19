package fvm

import (
	"collective-go-sdk/config"
	"context"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

var minerId = uint64(0100)
var result = big.NewInt(0)
var amount = big.NewInt(100)

// func TestDepositCallData(t *testing.T) {
// 	config, err := config.LoadConfig("../config")
// 	if err != nil {
// 		assert.Error(t, err)
// 	}

// 	client, err := NewLotusClient(config)

// 	abi, err := StorageProviderCollateral.StorageProviderCollateralMetaData.GetAbi()
// 	if err != nil {
// 		assert.Error(t, err)
// 	}

// 	callData, err := abi.Pack("deposit")

// 	tx, err := client.Deposit(amount, false)
// 	if err != nil {
// 		assert.Error(t, err)
// 	}

// 	assert.Equal(t, tx.Data(), callData)
// }

// func TestWithdrawCallData(t *testing.T) {
// 	config, err := config.LoadConfig("../config")
// 	if err != nil {
// 		assert.Error(t, err)
// 	}

// 	client, err := NewLotusClient(config)

// 	tx, err := client.Withdraw(amount, false)
// 	if err != nil {
// 		assert.Error(t, err)
// 	}

// 	abi, err := StorageProviderCollateral.StorageProviderCollateralMetaData.GetAbi()
// 	if err != nil {
// 		assert.Error(t, err)
// 	}

// 	callData, err := abi.Pack("withdraw", amount)
// 	assert.Equal(t, tx.Data(), callData)
// }

func TestGetCollateral(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadConfig("../config")

	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, "memory")
	if err != nil {
		assert.Error(t, err)
	}

	available, locked, err := client.GetCollateral(minerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, result.String(), available.String())
	assert.Equal(t, result.String(), locked.String())
}

func TestGetLockedCollateral(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadConfig("../config")

	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	lockedCollateral, err := client.GetLockedCollateral(minerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, result.String(), lockedCollateral.String())
}

func TestGetAvailableCollateral(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadConfig("../config")

	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	availableCollateral, err := client.GetAvailableCollateral(minerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, result.String(), availableCollateral.String())
}