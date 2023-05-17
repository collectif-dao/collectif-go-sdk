package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("./")

	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, 3141, config.ChainID)
	assert.Equal(t, "https://filecoin-hyperspace.chainstacklabs.com/rpc/v1", config.RPCAddress)
}
