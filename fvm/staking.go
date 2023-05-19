package fvm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *LotusClient) Pledge(amount *big.Int, send bool) (*types.Transaction, error) {
	var opts = c.Signer

	if !send {
		opts.NoSend = true
	} else {
		opts.NoSend = false
	}

	tx, err := c.Staking.Pledge(opts, amount)

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
