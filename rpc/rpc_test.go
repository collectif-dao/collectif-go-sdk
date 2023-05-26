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
	rpcClient := NewRPCClient(config.RPCAddress)

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
	rpcClient := NewRPCClient(config.RPCAddress)

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
	rpcClient := NewRPCClient(config.RPCAddress)

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
	rpcClient := NewRPCClient(config.RPCAddress)

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
	rpcClient := NewRPCClient(config.RPCAddress)

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
	rpcClient := NewRPCClient(config.RPCAddress)

	key, err := rpcClient.GetTipSetKey(ctx)
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotEmpty(t, key)
}
