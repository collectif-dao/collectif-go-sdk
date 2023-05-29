package fvm

import (
	"context"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type CollateralInfo struct {
	AvailableCollateral *big.Int
	LockedCollateral    *big.Int
}

func (c *LotusClient) Deposit(ctx context.Context, amount *big.Int, send bool) (*MessageResponse, error) {
	method := "deposit"
	calldata, err := c.calculateCalldata(method, c.Collateral.ABI)
	if err != nil {
		return nil, err
	}

	res, err := c.performLotusMessage(ctx, &c.Collateral.NativeAddress, method, amount, calldata, send)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c *LotusClient) Withdraw(ctx context.Context, amount *big.Int, send bool) (*MessageResponse, error) {
	method := "withdraw"
	calldata, err := c.calculateCalldata(method, c.Collateral.ABI, amount)
	if err != nil {
		return nil, err
	}

	res, err := c.performLotusMessage(ctx, &c.Collateral.NativeAddress, method, big.NewInt(0), calldata, send)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c *LotusClient) GetCollateral(ctx context.Context, ownerId uint64) (*CollateralInfo, error) {
	method := "getCollateral"
	var res [2]interface{}

	callData, err := c.calculateEthCalldata(method, c.Collateral.ABI, ownerId)
	if err != nil {
		return nil, err
	}

	if err = c.performEthCall(ctx, nil, &c.Collateral.Address, method, callData, c.Collateral.ABI, &res); err != nil {
		return nil, err
	}

	collateral := &CollateralInfo{
		AvailableCollateral: *abi.ConvertType(res[0], new(*big.Int)).(**big.Int),
		LockedCollateral:    *abi.ConvertType(res[1], new(*big.Int)).(**big.Int),
	}

	return collateral, nil
}

func (c *LotusClient) GetLockedCollateral(ctx context.Context, ownerId uint64) (*big.Int, error) {
	method := "getLockedCollateral"
	res := &big.Int{}

	callData, err := c.calculateEthCalldata(method, c.Collateral.ABI, ownerId)
	if err != nil {
		return nil, err
	}

	if err = c.performEthCall(ctx, nil, &c.Collateral.Address, method, callData, c.Collateral.ABI, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *LotusClient) GetAvailableCollateral(ctx context.Context, ownerId uint64) (*big.Int, error) {
	method := "getAvailableCollateral"
	res := &big.Int{}

	callData, err := c.calculateEthCalldata(method, c.Collateral.ABI, ownerId)
	if err != nil {
		return nil, err
	}

	if err = c.performEthCall(ctx, nil, &c.Collateral.Address, method, callData, c.Collateral.ABI, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *LotusClient) GetTotalSlashing(ctx context.Context, ownerId uint64) (*big.Int, error) {
	method := "slashings"
	res := &big.Int{}

	callData, err := c.calculateEthCalldata(method, c.Collateral.ABI, ownerId)
	if err != nil {
		return nil, err
	}

	if err = c.performEthCall(ctx, nil, &c.Collateral.Address, method, callData, c.Collateral.ABI, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *LotusClient) IsActiveSlashing(ctx context.Context, ownerId uint64) (bool, error) {
	method := "activeSlashings"
	var res bool

	callData, err := c.calculateEthCalldata(method, c.Collateral.ABI, ownerId)
	if err != nil {
		return false, err
	}

	if err = c.performEthCall(ctx, nil, &c.Collateral.Address, method, callData, c.Collateral.ABI, &res); err != nil {
		return false, err
	}

	return res, nil
}
