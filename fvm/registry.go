package fvm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type StorageProvider struct {
	isActive   bool
	targetPool common.Address
	minerId    uint64
	lastEpoch  int64
}

func (c *LotusClient) Register(minerId uint64, allocationLimit *big.Int, dailyAllocation *big.Int, send bool) (*types.Transaction, error) {
	var opts = c.Signer

	if !send {
		opts.NoSend = true
	}

	tx, err := c.Registry.Register(opts, minerId, allocationLimit, dailyAllocation)

	if err != nil {
		panic(err)
	}

	return tx, nil
}

func (c *LotusClient) GetStorageProvider(ownerId uint64) (StorageProvider, error) {
	var sp StorageProvider
	var opts = &bind.CallOpts{}

	isActive, targetPool, miner, maxRedeemablePeriod, err := c.Registry.StorageProviderRegistryCaller.GetStorageProvider(opts, ownerId)
	if err != nil {
		panic(err)
	}

	sp.isActive = isActive
	sp.targetPool = targetPool
	sp.minerId = miner
	sp.lastEpoch = maxRedeemablePeriod

	return sp, nil
}

func (c *LotusClient) IsActiveProvider(ownerId uint64) (bool, error) {
	var opts = &bind.CallOpts{}

	isActive, err := c.Registry.IsActiveProvider(opts, ownerId)
	if err != nil {
		panic(err)
	}

	return isActive, nil
}
