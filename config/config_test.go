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

	assert.Equal(t, 314159, config.RPCConfig["calibration"].ChainID)
	assert.Equal(t, 31415926, config.RPCConfig["localnet"].ChainID)
	assert.Equal(t, "https://api.calibration.node.glif.io/rpc/v1", config.RPCConfig["calibration"].Address)
	assert.Equal(t, "http://127.0.0.1:1234/rpc/v1", config.RPCConfig["localnet"].Address)
}
