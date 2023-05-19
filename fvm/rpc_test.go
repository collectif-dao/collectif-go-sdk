package fvm

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm/types"
	"context"
	"encoding/json"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	config, err := config.LoadConfig("../config")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	res, err := client.HandleRequest("Filecoin.Version", nil)
	if err != nil {
		assert.Error(t, err)
	}

	rpcRespBody := &types.RPCResponseBody{}
	err = json.Unmarshal(res, rpcRespBody)
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotEmpty(t, rpcRespBody)
}

func TestVerifyCid(t *testing.T) {
	config, err := config.LoadConfig("../config")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	status, err := client.VerifyCid(ctx, "bafy2bzacedeih423wkxn225h3r356jlu3qqrpmg43s2bcx6bq7kxhftncaiue")
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, status, true)
}

func TestGetNonce(t *testing.T) {
	config, err := config.LoadConfig("../config")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	addr, err := address.NewFromString("t410fjtte454usorrzedcbgslfmiwg5yy6i2a2pxoyky")
	if err != nil {
		assert.Error(t, err)
	}

	nonce, err := client.GetNonce(ctx, addr)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, nonce, uint64(5))
}

func TestLookupId(t *testing.T) {
	config, err := config.LoadConfig("../config")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	fAddr := "t410fmiuoet6qvv3jowc33r7bvlvcpgiu25kjec5t2qa"

	addr, err := address.NewFromString(fAddr)
	if err != nil {
		assert.Error(t, err)
	}

	key, err := client.GetTipSetKey(ctx)
	if err != nil {
		assert.Error(t, err)
	}

	idAddr, err := client.LookupId(ctx, addr, key)
	if err != nil {
		assert.Error(t, err)
	}

	id, err := address.IDFromAddress(*idAddr)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Greater(t, id, 0)
}

func TestGetChainHead(t *testing.T) {
	config, err := config.LoadConfig("../config")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	tipSet, err := client.GetChainHead(ctx)
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotEmpty(t, tipSet)
}

func TestGetTipSetKey(t *testing.T) {
	config, err := config.LoadConfig("../config")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	key, err := client.GetTipSetKey(ctx)
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotEmpty(t, key)
}
