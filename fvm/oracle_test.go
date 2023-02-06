package fvm

import (
	"collective-go-sdk/config"
	"collective-go-sdk/contracts/PledgeOracle"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

var epoch = big.NewInt(58660)
var preCommitDeposit = big.NewInt(1000)
var initialPledge = big.NewInt(230000)

func TestUpdateRecord(t *testing.T) {
	config, err := config.LoadConfig("../config")
	if err != nil {
		assert.Error(t, err)
	}

	abi, err := PledgeOracle.PledgeOracleMetaData.GetAbi()
	if err != nil {
		assert.Error(t, err)
	}
	callData, err := abi.Pack("updateRecord", epoch, preCommitDeposit, initialPledge)

	client, err := NewLotusClient(config)
	tx, err := client.UpdateRecord(epoch, preCommitDeposit, initialPledge, false)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, tx.Data(), callData)
}
