package fvm

import (
	"context"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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

func (c *LotusClient) Register(ctx context.Context, minerId uint64, allocationLimit *big.Int, dailyAllocation *big.Int, send bool) (*MessageResponse, error) {
	method := "register"
	calldata, err := c.calculateCalldata(method, c.Registry.ABI, minerId, allocationLimit, dailyAllocation)
	if err != nil {
		return nil, err
	}

	res, err := c.performLotusMessage(ctx, &c.Registry.NativeAddress, method, big.NewInt(0), calldata, send)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c *LotusClient) GetStorageProvider(ctx context.Context, ownerId uint64) (*StorageProvider, error) {
	method := "getStorageProvider"
	var res [4]interface{}

	callData, err := c.calculateEthCalldata(method, c.Registry.ABI, ownerId)
	if err != nil {
		return nil, err
	}

	if err = c.performEthCall(ctx, nil, &c.Registry.Address, method, callData, c.Registry.ABI, &res); err != nil {
		return nil, err
	}

	sp := &StorageProvider{
		isActive:   *abi.ConvertType(res[0], new(bool)).(*bool),
		targetPool: *abi.ConvertType(res[1], new(common.Address)).(*common.Address),
		minerId:    *abi.ConvertType(res[2], new(uint64)).(*uint64),
		lastEpoch:  *abi.ConvertType(res[3], new(int64)).(*int64),
	}

	return sp, nil
}

func (c *LotusClient) IsActiveProvider(ctx context.Context, ownerId uint64) (bool, error) {
	method := "isActiveProvider"
	var res bool

	callData, err := c.calculateEthCalldata(method, c.Registry.ABI, ownerId)
	if err != nil {
		return false, err
	}

	if err = c.performEthCall(ctx, nil, &c.Registry.Address, method, callData, c.Registry.ABI, &res); err != nil {
		return false, err
	}

	return res, nil
}

func (c *LotusClient) GetAllocations(ctx context.Context, ownerId uint64) (*SPAllocation, error) {
	method := "allocations"
	var res [6]interface{}

	callData, err := c.calculateEthCalldata(method, c.Registry.ABI, ownerId)
	if err != nil {
		return nil, err
	}

	if err = c.performEthCall(ctx, nil, &c.Registry.Address, method, callData, c.Registry.ABI, &res); err != nil {
		return nil, err
	}

	allocation := &SPAllocation{
		AllocationLimit: *abi.ConvertType(res[0], new(*big.Int)).(**big.Int),
		Repayment:       *abi.ConvertType(res[1], new(*big.Int)).(**big.Int),
		UsedAllocation:  *abi.ConvertType(res[2], new(*big.Int)).(**big.Int),
		DailyAllocation: *abi.ConvertType(res[3], new(*big.Int)).(**big.Int),
		AccruedRewards:  *abi.ConvertType(res[4], new(*big.Int)).(**big.Int),
		RepaidPledge:    *abi.ConvertType(res[5], new(*big.Int)).(**big.Int),
	}

	return allocation, nil
}

func (c *LotusClient) GetRestaking(ctx context.Context, ownerId uint64) (*SPRestaking, error) {
	method := "restakings"
	var res [2]interface{}

	callData, err := c.calculateEthCalldata(method, c.Registry.ABI, ownerId)
	if err != nil {
		return nil, err
	}

	if err = c.performEthCall(ctx, nil, &c.Registry.Address, method, callData, c.Registry.ABI, &res); err != nil {
		return nil, err
	}

	restaking := &SPRestaking{
		RestakingRatio:   *abi.ConvertType(res[0], new(*big.Int)).(**big.Int),
		RestakingAddress: *abi.ConvertType(res[1], new(common.Address)).(*common.Address),
	}

	return restaking, nil
}

func (c *LotusClient) GetBeneficiaryStatus(ctx context.Context, ownerId uint64) (bool, error) {
	method := "getBeneficiaryStatus"
	var res bool

	callData, err := c.calculateEthCalldata(method, c.Registry.ABI, ownerId)
	if err != nil {
		return false, err
	}

	if err = c.performEthCall(ctx, nil, &c.Registry.Address, method, callData, c.Registry.ABI, &res); err != nil {
		return false, err
	}

	return res, nil
}

func (c *LotusClient) GetSectorSize(ctx context.Context, ownerId uint64) (uint64, error) {
	method := "getSectorSize"
	var res uint64

	callData, err := c.calculateEthCalldata(method, c.Registry.ABI, ownerId)
	if err != nil {
		return 0, err
	}

	if err = c.performEthCall(ctx, nil, &c.Registry.Address, method, callData, c.Registry.ABI, &res); err != nil {
		return 0, err
	}

	return res, nil
}

func (c *LotusClient) SetRestaking(ctx context.Context, restakingRatio *big.Int, restakingAddress common.Address, send bool) (*MessageResponse, error) {
	method := "setRestaking"
	calldata, err := c.calculateCalldata(method, c.Registry.ABI, restakingRatio, restakingAddress)
	if err != nil {
		return nil, err
	}

	res, err := c.performLotusMessage(ctx, &c.Registry.NativeAddress, method, big.NewInt(0), calldata, send)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c *LotusClient) RequestAllocationLimitUpdate(ctx context.Context, allocationLimit *big.Int, dailyAllocation *big.Int, send bool) (*MessageResponse, error) {
	method := "requestAllocationLimitUpdate"
	calldata, err := c.calculateCalldata(method, c.Registry.ABI, allocationLimit, dailyAllocation)
	if err != nil {
		return nil, err
	}

	res, err := c.performLotusMessage(ctx, &c.Registry.NativeAddress, method, big.NewInt(0), calldata, send)
	if err != nil {
		return res, err
	}

	return res, nil
}
