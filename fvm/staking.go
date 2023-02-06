package fvm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *LotusClient) Pledge(sectorNumber uint64, proof []byte, send bool) (*types.Transaction, error) {
	var opts = c.signer

	if !send {
		opts.NoSend = true
	} else {
		opts.NoSend = false
	}

	tx, err := c.staking.Pledge(opts, sectorNumber, proof)

	if err != nil {
		panic(err)
	}

	return tx, nil
}

func (c *LotusClient) PledgeAggregate(sectorNumber []uint64, proof [][]byte, send bool) (*types.Transaction, error) {
	var opts = c.signer

	if !send {
		opts.NoSend = true
	} else {
		opts.NoSend = false
	}

	tx, err := c.staking.PledgeAggregate(opts, sectorNumber, proof)

	if err != nil {
		panic(err)
	}

	return tx, nil
}

func (c *LotusClient) WithdrawBalance(miner []byte, amount *big.Int, send bool) (*types.Transaction, error) {
	var opts = c.signer

	if !send {
		opts.NoSend = true
	} else {
		opts.NoSend = false
	}

	tx, err := c.staking.WithdrawRewards(opts, miner, amount)

	if err != nil {
		panic(err)
	}

	return tx, nil
}

func (c *LotusClient) TotalAssets() (*big.Int, error) {
	return c.staking.TotalAssets(&bind.CallOpts{})
}

func (c *LotusClient) TotalFilAvailable() (*big.Int, error) {
	return c.staking.TotalFilAvailable(&bind.CallOpts{})
}
