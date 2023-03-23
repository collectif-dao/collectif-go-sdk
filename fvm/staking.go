package fvm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *LotusClient) Pledge(sectorNumber uint64, proof []byte, send bool) (*types.Transaction, error) {
	var opts = c.Signer

	if !send {
		opts.NoSend = true
	} else {
		opts.NoSend = false
	}

	tx, err := c.Staking.Pledge(opts, sectorNumber, proof)

	if err != nil {
		panic(err)
	}

	return tx, nil
}

func (c *LotusClient) PledgeAggregate(sectorNumber []uint64, proof [][]byte, send bool) (*types.Transaction, error) {
	var opts = c.Signer

	if !send {
		opts.NoSend = true
	} else {
		opts.NoSend = false
	}

	tx, err := c.Staking.PledgeAggregate(opts, sectorNumber, proof)

	if err != nil {
		panic(err)
	}

	return tx, nil
}

func (c *LotusClient) WithdrawBalance(miner []byte, amount *big.Int, send bool) (*types.Transaction, error) {
	var opts = c.Signer

	if !send {
		opts.NoSend = true
	} else {
		opts.NoSend = false
	}

	tx, err := c.Staking.WithdrawRewards(opts, miner, amount)

	if err != nil {
		panic(err)
	}

	return tx, nil
}

func (c *LotusClient) TotalAssets() (*big.Int, error) {
	return c.Staking.TotalAssets(&bind.CallOpts{})
}

func (c *LotusClient) TotalFilAvailable() (*big.Int, error) {
	return c.Staking.TotalFilAvailable(&bind.CallOpts{})
}
