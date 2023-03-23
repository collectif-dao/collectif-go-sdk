package fvm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type Record struct {
	preCommitDeposit *big.Int
	initialPledge    *big.Int
}

func (c *LotusClient) UpdateRecord(epoch *big.Int, preCommitDeposit *big.Int, initialPledge *big.Int, send bool) (*types.Transaction, error) {
	var opts = c.Signer

	if send {
		opts.NoSend = false
	} else {
		opts.NoSend = true
	}

	tx, err := c.Oracle.UpdateRecord(opts, epoch, preCommitDeposit, initialPledge)
	if err != nil {
		panic(err)
	}

	return tx, nil
}

func (c *LotusClient) GetLastRecord() (Record, error) {
	preCommitDeposit, initialPledge, err := c.Oracle.GetLastRecord(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	return Record{
		preCommitDeposit,
		initialPledge,
	}, nil
}

func (c *LotusClient) GetPledgeFees() (*big.Int, error) {
	return c.Oracle.GetPledgeFees(&bind.CallOpts{})
}
