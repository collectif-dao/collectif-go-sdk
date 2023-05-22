package utils

import (
	"collective-go-sdk/fvm"
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
)

func GetIdAddress(ctx context.Context, addr string, c *fvm.LotusClient) uint64 {
	var id uint64
	var err error
	var fAddr address.Address

	fAddr, err = address.NewFromString(addr)
	if err != nil {
		ethAddr, err := ethtypes.ParseEthAddress(addr)

		fAddr, err = ethAddr.ToFilecoinAddress()
		if err != nil {
			panic(fmt.Sprintf("unable to get Filecoin address: %v", err))
		}
	}

	protocol := fAddr.Protocol()

	if protocol == address.ID {
		id, err = address.IDFromAddress(fAddr)
		if err != nil {
			panic(fmt.Sprintf("unable to get ID from the address.Address: %v", err))
		}
	} else {
		key, err := c.GetTipSetKey(ctx)
		if err != nil {
			panic(fmt.Sprintf("unable to get tipset key for chain head: %v", err))
		}

		idAddr, err := c.LookupId(ctx, fAddr, key)
		if err != nil {
			panic(fmt.Sprintf("unable to get find actor ID for specified address: %v", err))
		}

		id, err = address.IDFromAddress(*idAddr)
		if err != nil {
			panic(fmt.Sprintf("unable to get ID from the address.Address: %v", err))
		}
	}

	return id
}
