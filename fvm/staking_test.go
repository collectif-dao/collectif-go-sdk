package fvm

import (
	"collective-go-sdk/config"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
)

func TestSymbol(t *testing.T) {
	config, err := config.LoadConfig("../config")

	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(config)

	symbol, err := client.staking.Symbol(&bind.CallOpts{})
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, "clFIL", symbol)
}

func TestName(t *testing.T) {
	config, err := config.LoadConfig("../config")

	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(config)

	symbol, err := client.staking.Name(&bind.CallOpts{})
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, "Collective Staked FIL", symbol)
}
