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

type SPAllocation struct {
	AllocationLimit *big.Int
	Repayment       *big.Int
	UsedAllocation  *big.Int
	DailyAllocation *big.Int
	AccruedRewards  *big.Int
	RepaidPledge    *big.Int
}

type SPRestaking struct {
	RestakingRatio   *big.Int
	RestakingAddress common.Address
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

func (c *LotusClient) GetAllocations(ownerId uint64) (*SPAllocation, error) {
	var opts = &bind.CallOpts{}
	var err error

	allocation := SPAllocation{}
	allocation, err = c.Registry.StorageProviderRegistryCaller.Allocations(opts, ownerId)
	if err != nil {
		panic(err)
	}

	return &allocation, nil
}

func (c *LotusClient) GetRestaking(ownerId uint64) (*SPRestaking, error) {
	var opts = &bind.CallOpts{}
	var err error

	restaking := SPRestaking{}
	restaking, err = c.Registry.StorageProviderRegistryCaller.Restakings(opts, ownerId)
	if err != nil {
		panic(err)
	}

	return &restaking, nil
}

func (c *LotusClient) GetBeneficiaryStatus(ownerId uint64) (bool, error) {
	var opts = &bind.CallOpts{}

	status, err := c.Registry.StorageProviderRegistryCaller.BeneficiaryStatus(opts, ownerId)
	if err != nil {
		panic(err)
	}

	return status, nil
}

func (c *LotusClient) GetSectorSize(ownerId uint64) (uint64, error) {
	var opts = &bind.CallOpts{}

	size, err := c.Registry.StorageProviderRegistryCaller.SectorSizes(opts, ownerId)
	if err != nil {
		panic(err)
	}

	return size, nil
}

func (c *LotusClient) SetRestaking(restakingRatio *big.Int, restakingAddress common.Address, send bool) (*types.Transaction, error) {
	var opts = c.Signer

	if !send {
		opts.NoSend = true
	}

	tx, err := c.Registry.SetRestaking(opts, restakingRatio, restakingAddress)

	if err != nil {
		panic(err)
	}

	return tx, nil
}

func (c *LotusClient) RequestAllocationLimitUpdate(allocationLimit *big.Int, dailyAllocation *big.Int, send bool) (*types.Transaction, error) {
	var opts = c.Signer

	if !send {
		opts.NoSend = true
	}

	tx, err := c.Registry.RequestAllocationLimitUpdate(opts, allocationLimit, dailyAllocation)

	if err != nil {
		panic(err)
	}

	return tx, nil
}
