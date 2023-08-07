package fvm

import (
	"context"
	"math/big"

	log "github.com/sirupsen/logrus"
)

func (c *LotusClient) Pledge(ctx context.Context, amount *big.Int, minerId uint64, send bool) (*MessageResponse, error) {
	log.Info("Pledging ", AttoFIL2FIL_str(amount), " amount of FIL")
	method := "pledge"
	calldata, err := c.calculateCalldata(method, c.Staking.ABI, amount, minerId)
	if err != nil {
		return nil, err
	}

	res, err := c.performLotusMessage(ctx, &c.Staking.NativeAddress, method, big.NewInt(0), calldata, send)
	if err != nil {
		return res, err
	}

	log.Info("Succesfully pledged ", AttoFIL2FIL_str(amount), " amount of FIL")
	return res, nil
}

func (c *LotusClient) Stake(ctx context.Context, amount *big.Int, send bool) (*MessageResponse, error) {
	log.Info("Staking ", AttoFIL2FIL_str(amount), " amount of FIL")
	method := "stake"
	calldata, err := c.calculateCalldata(method, c.Staking.ABI)
	if err != nil {
		return nil, err
	}

	res, err := c.performLotusMessage(ctx, &c.Staking.NativeAddress, method, amount, calldata, send)
	if err != nil {
		return res, err
	}

	log.Info("Succesfully staked ", AttoFIL2FIL_str(amount), " amount of FIL")
	return res, nil
}

func (c *LotusClient) TotalAssets(ctx context.Context) (*big.Int, error) {
	method := "totalAssets"
	res := &big.Int{}

	callData, err := c.calculateEthCalldata(method, c.Staking.ABI)
	if err != nil {
		return nil, err
	}

	if err = c.performEthCall(ctx, nil, &c.Staking.Address, method, callData, c.Staking.ABI, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *LotusClient) TotalFilAvailable(ctx context.Context) (*big.Int, error) {
	method := "totalFilAvailable"
	res := &big.Int{}

	callData, err := c.calculateEthCalldata(method, c.Staking.ABI)
	if err != nil {
		return nil, err
	}

	if err = c.performEthCall(ctx, nil, &c.Staking.Address, method, callData, c.Staking.ABI, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *LotusClient) TotalFilPledged(ctx context.Context) (*big.Int, error) {
	method := "totalFilPledged"
	res := &big.Int{}

	callData, err := c.calculateEthCalldata(method, c.Staking.ABI)
	if err != nil {
		return nil, err
	}

	if err = c.performEthCall(ctx, nil, &c.Staking.Address, method, callData, c.Staking.ABI, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *LotusClient) TotalFees(ctx context.Context, ownerId uint64) (*big.Int, error) {
	method := "totalFees"
	res := &big.Int{}

	callData, err := c.calculateEthCalldata(method, c.Staking.ABI, ownerId)
	if err != nil {
		return nil, err
	}

	if err = c.performEthCall(ctx, nil, &c.Staking.Address, method, callData, c.Staking.ABI, &res); err != nil {
		return nil, err
	}

	return res, nil
}
