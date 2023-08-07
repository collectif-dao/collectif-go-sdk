package fvm

import (
	"context"
	"fmt"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	fAbi "github.com/filecoin-project/go-state-types/abi"
	fBig "github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/builtin/v9/miner"
	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/types"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type StorageProvider struct {
	isActive   bool
	targetPool common.Address
	ownerId    uint64
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
	log.Info("Registering ", minerId, " minerID with ", AttoFIL2FIL_str(allocationLimit), " FIL allocation limit and ", AttoFIL2FIL_str(dailyAllocation), " FIL daily allocation")

	mAddr, err := address.NewIDAddress(minerId)
	if err != nil {
		return nil, err
	}

	method := "register"
	calldata, err := c.calculateCalldata(method, c.Registry.ABI, minerId, allocationLimit, dailyAllocation)
	if err != nil {
		return nil, err
	}

	res, err := c.performLotusMessage(ctx, &c.Registry.NativeAddress, method, big.NewInt(0), calldata, send)
	if err != nil {
		return res, err
	}

	log.Info("Succesfully registered ", mAddr.String(), " minerId")
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
		ownerId:    *abi.ConvertType(res[2], new(uint64)).(*uint64),
		lastEpoch:  *abi.ConvertType(res[3], new(int64)).(*int64),
	}

	return sp, nil
}

func (c *LotusClient) GetMinerOwner(ctx context.Context, minerId uint64) (uint64, error) {
	sp, err := c.GetStorageProvider(ctx, minerId)
	if err != nil {
		return 0, err
	}

	return sp.ownerId, nil
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

func (c *LotusClient) RequestAllocationLimitUpdate(ctx context.Context, minerId uint64, allocationLimit *big.Int, dailyAllocation *big.Int, send bool) (*MessageResponse, error) {
	method := "requestAllocationLimitUpdate"
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

func (c *LotusClient) ChangeBeneficiaryAddress(ctx context.Context, minerAddr *address.Address, beneficiaryAddr *address.Address, quota *big.Int, expiration int64, send bool) (*MessageResponse, error) {
	tk, err := c.RPCClient.GetTipSetKey(ctx)
	if err != nil {
		return nil, xerrors.Errorf("unable to get tipset head key: %w", err)
	}

	idAddr, err := c.RPCClient.LookupId(ctx, *beneficiaryAddr, tk)
	if err != nil {
		return nil, xerrors.Errorf("invalid beneficiary address param: %w", err)
	}

	quotaVal, err := fBig.FromString(quota.String())
	if err != nil {
		return nil, xerrors.Errorf("invalid quota param: %w", err)
	}

	params := &miner.ChangeBeneficiaryParams{
		NewBeneficiary: *idAddr,
		NewQuota:       quotaVal,
		NewExpiration:  fAbi.ChainEpoch(expiration),
	}

	sp, err := actors.SerializeParams(params)
	if err != nil {
		return nil, xerrors.Errorf("serializing params: %w", err)
	}

	mi, err := c.RPCClient.StateMinerInfo(ctx, *minerAddr, tk)
	if err != nil {
		return nil, xerrors.Errorf("invalid beneficiary address param: %w", err)
	}

	if mi.Beneficiary == mi.Owner && *idAddr == mi.Owner {
		return nil, fmt.Errorf("beneficiary %s already set to owner address", mi.Beneficiary)
	}

	msg := &types.Message{
		From:   *c.Address,
		To:     *minerAddr,
		Method: builtin.MethodsMiner.ChangeBeneficiary,
		Value:  fBig.Zero(),
		Params: sp,
	}

	var res MessageResponse
	res.Data = sp

	res, err = c.executeNativeMessage(ctx, msg, res, send)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
