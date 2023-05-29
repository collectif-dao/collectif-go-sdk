package rpc

import (
	"collective-go-sdk/config"
	"context"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	rpcClient := NewRPCClient(config.RPCConfig[config.DefaultNetwork].Address, config.RPCConfig[config.DefaultNetwork].APIToken)

	version, err := rpcClient.Version(ctx)
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotEmpty(t, version.APIVersion)
	assert.NotEmpty(t, version.BlockDelay)
	assert.NotEmpty(t, version.Version)
}

func TestVerifyCid(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	rpcClient := NewRPCClient(config.RPCConfig[config.DefaultNetwork].Address, config.RPCConfig[config.DefaultNetwork].APIToken)

	status, err := rpcClient.VerifyCid(ctx, "bafy2bzacedeih423wkxn225h3r356jlu3qqrpmg43s2bcx6bq7kxhftncaiue")
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, status, true)
}

func TestGetNonce(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	rpcClient := NewRPCClient(config.RPCConfig[config.DefaultNetwork].Address, config.RPCConfig[config.DefaultNetwork].APIToken)

	addr, err := address.NewFromString("t410fjtte454usorrzedcbgslfmiwg5yy6i2a2pxoyky")
	if err != nil {
		assert.Error(t, err)
	}

	nonce, err := rpcClient.GetNonce(ctx, addr)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, nonce, uint64(5))
}

func TestLookupId(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	rpcClient := NewRPCClient(config.RPCConfig[config.DefaultNetwork].Address, config.RPCConfig[config.DefaultNetwork].APIToken)

	fAddr := "t410fmiuoet6qvv3jowc33r7bvlvcpgiu25kjec5t2qa"

	addr, err := address.NewFromString(fAddr)
	if err != nil {
		assert.Error(t, err)
	}

	key, err := rpcClient.GetTipSetKey(ctx)
	if err != nil {
		assert.Error(t, err)
	}

	idAddr, err := rpcClient.LookupId(ctx, addr, key)
	if err != nil {
		assert.Error(t, err)
	}

	id, err := address.IDFromAddress(*idAddr)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Greater(t, id, uint64(0))
}

func TestGetChainHead(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	rpcClient := NewRPCClient(config.RPCConfig[config.DefaultNetwork].Address, config.RPCConfig[config.DefaultNetwork].APIToken)

	tipSet, err := rpcClient.GetChainHead(ctx)
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotEmpty(t, tipSet)
}

func TestGetTipSetKey(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	rpcClient := NewRPCClient(config.RPCConfig[config.DefaultNetwork].Address, config.RPCConfig[config.DefaultNetwork].APIToken)

	key, err := rpcClient.GetTipSetKey(ctx)
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotEmpty(t, key)
}

func TestStateMinerInfo(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	rpcClient := NewRPCClient(config.RPCConfig[config.DefaultNetwork].Address, config.RPCConfig[config.DefaultNetwork].APIToken)

	fAddr := "t2fp6lovd2v565vli6e2edltx6tgwopqepfd3rarq"
	// fAddr := "t035955"

	mAddr, err := address.NewFromString(fAddr)
	if err != nil {
		assert.Error(t, err)
	}

	key, err := rpcClient.GetTipSetKey(ctx)
	if err != nil {
		assert.Error(t, err)
	}

	info, err := rpcClient.StateMinerInfo(ctx, mAddr, key)
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotEmpty(t, info)
}
