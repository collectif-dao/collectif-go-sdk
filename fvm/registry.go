package fvm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type StorageProvider struct {
	isActive            bool
	targetPool          common.Address
	miner               []byte
	allocationLimit     *big.Int
	usedAllocation      *big.Int
	accruedRewards      *big.Int
	lockedRewards       *big.Int
	maxRedeemablePeriod *big.Int
}

func (c *LotusClient) Register(miner []byte, targetPool common.Address, allocationLimit *big.Int, period *big.Int, send bool) (*types.Transaction, error) {
	var opts = c.signer

	if !send {
		opts.NoSend = true
	}

	tx, err := c.registry.Register(opts, miner, targetPool, allocationLimit, period)

	if err != nil {
		panic(err)
	}

	return tx, nil
}

func (c *LotusClient) ChangeBeneficiaryAddress(beneficiaryAddress common.Address, send bool) (*types.Transaction, error) {
	var opts = c.signer

	if !send {
		opts.NoSend = true
	}

	tx, err := c.registry.ChangeBeneficiaryAddress(opts, beneficiaryAddress)

	if err != nil {
		panic(err)
	}

	return tx, nil
}

func (c *LotusClient) GetStorageProvider(provider []byte) (StorageProvider, error) {
	var sp StorageProvider
	var opts = &bind.CallOpts{}

	isActive, targetPool, miner, allocation, usedAllocation, accruedRewards, lockedRewards, maxRedeemablePeriod, err := c.registry.StorageProviderRegistryCaller.GetStorageProvider(opts, provider)
	if err != nil {
		panic(err)
	}

	sp.isActive = isActive
	sp.targetPool = targetPool
	sp.miner = miner
	sp.allocationLimit = allocation
	sp.usedAllocation = usedAllocation
	sp.accruedRewards = accruedRewards
	sp.lockedRewards = lockedRewards
	sp.maxRedeemablePeriod = maxRedeemablePeriod

	return sp, nil
}

func (c *LotusClient) IsActiveProvider(provider []byte) (bool, error) {
	var opts = &bind.CallOpts{}

	isActive, err := c.registry.IsActiveProvider(opts, provider)
	if err != nil {
		panic(err)
	}

	return isActive, nil
}
