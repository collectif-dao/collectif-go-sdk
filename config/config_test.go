package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, 3141, config.RPCConfig["hyperspace"].ChainID)
	assert.Equal(t, 31415926, config.RPCConfig["localnet"].ChainID)
	assert.Equal(t, "https://api.hyperspace.node.glif.io/rpc/v1", config.RPCConfig["hyperspace"].Address)
	assert.Equal(t, "http://127.0.0.1:1234/rpc/v1", config.RPCConfig["localnet"].Address)
}
