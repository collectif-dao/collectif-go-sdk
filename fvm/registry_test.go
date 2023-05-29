package fvm

import (
	"collective-go-sdk/config"
	"collective-go-sdk/keystore"
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/stretchr/testify/assert"
)

func getAttoFilFromFIL(amount int) *big.Int {
	fil := big.NewInt(1000000000000000000)
	v := big.NewInt(int64(amount))

	return big.NewInt(1).Mul(v, fil)
}

func getOwnerId(t *testing.T, ctx context.Context, client *LotusClient) uint64 {
	ownerAddress := "t3r3nqkq7ybn2ozj4lvvpdwayyy3b5v6l7gnrgg7yd5d42ti65qfaamvdphnqscek3ruoo7rnulixbsyh52tla"
	fAddr, err := address.NewFromString(ownerAddress)
	if err != nil {
		assert.Error(t, err)
	}

	key, err := client.RPCClient.GetTipSetKey(ctx)
	if err != nil {
		panic(fmt.Sprintf("unable to get tipset key for chain head: %v", err))
	}

	idAddr, err := client.RPCClient.LookupId(ctx, fAddr, key)
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

	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	ownerId := getOwnerId(t, ctx, client)
	sp, err := client.GetStorageProvider(ctx, ownerId)
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

	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	ownerId := getOwnerId(t, ctx, client)
	allocation, err := client.GetAllocations(ctx, ownerId)
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

	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	ownerId := getOwnerId(t, ctx, client)
	restaking, err := client.GetRestaking(ctx, ownerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, restaking.RestakingRatio.String(), "0")
	assert.Equal(t, restaking.RestakingAddress, common.HexToAddress("0x"))
}

func TestGetBeneficiaryStatus(t *testing.T) {
	ctx := context.Background()

	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	ownerId := getOwnerId(t, ctx, client)
	status, err := client.GetBeneficiaryStatus(ctx, ownerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, status, false)
}

func TestGetSectorSize(t *testing.T) {
	ctx := context.Background()

	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	ownerId := getOwnerId(t, ctx, client)
	size, err := client.GetSectorSize(ctx, ownerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, size, uint64(0))
}

func TestIsActiveSP(t *testing.T) {
	ctx := context.Background()

	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(ctx, config, keystore.MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	ownerId := getOwnerId(t, ctx, client)
	status, err := client.IsActiveProvider(ctx, ownerId)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, status, false)
}

func TestRegisterCallData(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, keystore.FSKeyStore)

	minerId := uint64(53149)
	allocationLimit := getAttoFilFromFIL(100000)
	dailyAllocation := getAttoFilFromFIL(1000)

	callData, err := client.calculateCalldata("register", client.Registry.ABI, minerId, allocationLimit, dailyAllocation)
	if err != nil {
		assert.Error(t, err)
	}

	res, err := client.Register(ctx, minerId, allocationLimit, dailyAllocation, false)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, callData, res.Data)
}

func TestSetRestaking(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, keystore.FSKeyStore)

	restakingRatio := big.NewInt(8000)
	restakingAddress := common.HexToAddress("0x000000000000000000000000000000000000dEaD")

	callData, err := client.calculateCalldata("setRestaking", client.Registry.ABI, restakingRatio, restakingAddress)
	if err != nil {
		assert.Error(t, err)
	}

	res, err := client.SetRestaking(ctx, restakingRatio, restakingAddress, false)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, callData, res.Data)
}

func TestRequestAllocationUpdate(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, keystore.FSKeyStore)

	allocationLimit := getAttoFilFromFIL(100000)
	dailyAllocation := getAttoFilFromFIL(1000)

	callData, err := client.calculateCalldata("requestAllocationLimitUpdate", client.Registry.ABI, allocationLimit, dailyAllocation)
	if err != nil {
		assert.Error(t, err)
	}

	res, err := client.RequestAllocationLimitUpdate(ctx, allocationLimit, dailyAllocation, false)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, callData, res.Data)
}

func TestChangeBeneficiary(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, keystore.FSKeyStore)

	mA := "t2fp6lovd2v565vli6e2edltx6tgwopqepfd3rarq"
	bA := "t3q5szei4jstv3aqnqkimkehpbhiwj2urjt4pgiih2s5ky6cuynjedhfufvdvjybw3rnurfhwfwpcbnqa6q2tq"
	quota := getAttoFilFromFIL(1000)
	expiration := int64(384804)

	mAddr, err := address.NewFromString(mA)
	if err != nil {
		assert.Error(t, err)
	}

	bAddr, err := address.NewFromString(bA)
	if err != nil {
		assert.Error(t, err)
	}

	res, err := client.ChangeBeneficiaryAddress(ctx, &mAddr, &bAddr, quota, expiration, false)
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotEmpty(t, res.Data)
}
