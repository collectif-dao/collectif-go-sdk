package fvm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type Collateral struct {
	availableCollateral *big.Int
	lockedCollateral    *big.Int
}

func (c *LotusClient) Deposit(amount *big.Int, send bool) (*types.Transaction, error) {
	var opts = c.Signer

	if !send {
		opts.NoSend = true
		opts.Value = amount
	} else {
		opts.NoSend = false
		opts.Value = amount
	}

	tx, err := c.Collateral.Deposit(opts)

	if err != nil {
		panic(err)
	}

	return tx, nil
}

func (c *LotusClient) Withdraw(amount *big.Int, send bool) (*types.Transaction, error) {
	var opts = c.Signer

	if !send {
		opts.NoSend = true
	} else {
		opts.NoSend = false
	}

	tx, err := c.Collateral.Withdraw(opts, amount)

	if err != nil {
		panic(err)
	}

	return tx, nil
}

func (c *LotusClient) GetCollateral(provider []byte) (*big.Int, *big.Int, error) {
	availableCollateral, lockedCollateral, err := c.Collateral.GetCollateral(&bind.CallOpts{}, provider)
	if err != nil {
		panic(err)
	}

	return availableCollateral, lockedCollateral, nil
}

func (c *LotusClient) GetLockedCollateral(provider []byte) (*big.Int, error) {
	return c.Collateral.GetLockedCollateral(&bind.CallOpts{}, provider)
}

func (c *LotusClient) GetAvailableCollateral(provider []byte) (*big.Int, error) {
	return c.Collateral.GetAvailableCollateral(&bind.CallOpts{}, provider)
}
