package worker

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"collective-go-sdk/keystore"
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewWorker(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := fvm.NewLotusClient(ctx, config, keystore.FSKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	minerList := config.Miners[config.DefaultNetwork]
	miners, err := PrepareMiners(ctx, minerList, client)
	if err != nil {
		assert.Error(t, err)
	}

	worker := NewWorker(time.Duration(config.PledgeTimeout)*time.Hour, client, miners, false)
	assert.Equal(t, worker.Client, client)
	assert.Equal(t, worker.Miners, miners)
}

func TestStart(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := fvm.NewLotusClient(ctx, config, keystore.FSKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	minerList := config.Miners[config.DefaultNetwork]
	miners, err := PrepareMiners(ctx, minerList, client)
	if err != nil {
		assert.Error(t, err)
	}

	_ = NewWorker(time.Duration(1*time.Second), client, miners, false)

	// worker.Start(ctx)
}

func TestTryPledge(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := fvm.NewLotusClient(ctx, config, keystore.FSKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	minerList := []string{"f063869", "f063868"}
	miners, err := PrepareMiners(ctx, minerList, client)
	if err != nil {
		assert.Error(t, err)
	}

	worker := NewWorker(time.Duration(1*time.Second), client, miners, false)

	miners[0].DailyAllocation = big.NewInt(1).Mul(big.NewInt(2441262), big.NewInt(1000000000000000))
	miners[0].TotalAllocation = big.NewInt(2)

	msg, err := worker.tryPledge(ctx, &miners[0])
	if err != nil {
		assert.Equal(t, err, errors.New("Not enough FIL in the Liquid Staking contract"))
	}

	miners[0].DailyAllocation = big.NewInt(1)
	miners[0].TotalAllocation = big.NewInt(0)
	miners[0].CollateralRequirements = big.NewInt(3000)

	msg, err = worker.tryPledge(ctx, &miners[0])
	if err != nil {
		assert.Equal(t, err, errors.New("Total allocation overflow"))
	}

	miners[0].TotalAllocation = big.NewInt(2)

	msg, err = worker.tryPledge(ctx, &miners[0])
	if err != nil {
		assert.Error(t, err)
	}

	if msg != nil {
		assert.NotEmpty(t, msg.Data)
	}
}

func TestIncreaseCollateral(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := fvm.NewLotusClient(ctx, config, keystore.FSKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	minerList := []string{"f063869", "f063868"}
	miners, err := PrepareMiners(ctx, minerList, client)
	if err != nil {
		assert.Error(t, err)
	}

	worker := NewWorker(time.Duration(1*time.Second), client, miners, false)

	amt := big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1000000000000000))

	miners[0].DailyAllocation = amt
	miners[0].CollateralRequirements = big.NewInt(3000)

	msg, err := worker.increaseCollateral(ctx, &miners[0], amt)
	if err != nil {
		assert.Error(t, err)
	}
	assert.NotEmpty(t, msg.Data)

	var lc = big.NewInt(0)
	lc.Mul(amt, big.NewInt(3500)).Div(lc, big.NewInt(10000))
	miners[0].LockedCollateral = lc

	msg, err = worker.increaseCollateral(ctx, &miners[0], amt)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Nil(t, msg)
}
