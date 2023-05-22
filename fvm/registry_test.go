package fvm

import (
	"collective-go-sdk/config"
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/stretchr/testify/assert"
)

func getOwnerId(t *testing.T, ctx context.Context, client *LotusClient) uint64 {
	ownerAddress := "t3r3nqkq7ybn2ozj4lvvpdwayyy3b5v6l7gnrgg7yd5d42ti65qfaamvdphnqscek3ruoo7rnulixbsyh52tla"
	fAddr, err := address.NewFromString(ownerAddress)
	if err != nil {
		assert.Error(t, err)
	}

	key, err := client.GetTipSetKey(ctx)
	if err != nil {
		panic(fmt.Sprintf("unable to get tipset key for chain head: %v", err))
	}

	idAddr, err := client.LookupId(ctx, fAddr, key)
	if err != nil {
		panic(fmt.Sprintf("unable to get find actor ID for specified address: %v", err))
	}

	ownerId, err := address.IDFromAddress(*idAddr)
	if err != nil {
		panic(fmt.Sprintf("unable to get ID from the address.Address: %v", err))
	}

	return ownerId
}

func TestGetStorageProvider(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadConfig("../config")

	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, "memory")
	if err != nil {
		assert.Error(t, err)
	}

	ownerId := getOwnerId(t, ctx, client)
	sp, err := client.GetStorageProvider(ownerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, sp.isActive, false)
	assert.Equal(t, sp.lastEpoch, int64(0))
	assert.Equal(t, sp.minerId, uint64(0))
	assert.Equal(t, sp.targetPool, common.HexToAddress("0x"))
}

func TestGetAllocations(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadConfig("../config")

	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, "memory")
	if err != nil {
		assert.Error(t, err)
	}

	ownerId := getOwnerId(t, ctx, client)
	allocation, err := client.GetAllocations(ownerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, allocation.AccruedRewards.String(), "0")
	assert.Equal(t, allocation.AllocationLimit.String(), "0")
	assert.Equal(t, allocation.DailyAllocation.String(), "0")
	assert.Equal(t, allocation.RepaidPledge.String(), "0")
	assert.Equal(t, allocation.Repayment.String(), "0")
	assert.Equal(t, allocation.UsedAllocation.String(), "0")
}

func TestGetRestakings(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadConfig("../config")

	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, "memory")
	if err != nil {
		assert.Error(t, err)
	}

	ownerId := getOwnerId(t, ctx, client)
	restaking, err := client.GetRestaking(ownerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, restaking.RestakingRatio.String(), "0")
	assert.Equal(t, restaking.RestakingAddress, common.HexToAddress("0x"))
}

func TestGetBeneficiaryStatus(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadConfig("../config")

	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, "memory")
	if err != nil {
		assert.Error(t, err)
	}

	ownerId := getOwnerId(t, ctx, client)
	status, err := client.GetBeneficiaryStatus(ownerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, status, false)
}

func TestGetSectorSize(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadConfig("../config")

	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, "memory")
	if err != nil {
		assert.Error(t, err)
	}

	ownerId := getOwnerId(t, ctx, client)
	size, err := client.GetSectorSize(ownerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, size, uint64(0))
}
