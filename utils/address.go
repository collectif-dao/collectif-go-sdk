package utils

import (
	"collective-go-sdk/fvm"
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
)

// convertAddress for converting string address to bytes for Solidity interactions
func ConvertAddress(addr string) []byte {
	fAddr, err := address.NewFromString(addr)
	if err != nil {
		panic(fmt.Sprintf("go-address is unable to initiate Address struct: %v", err))
	}

	return fAddr.Bytes()
}

func EncodePacked(address string) []byte {
	return encodePacked(encodeBytesString(address))
}

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
		key, err := c.RPCClient.GetTipSetKey(ctx)
		if err != nil {
			panic(fmt.Sprintf("unable to get tipset key for chain head: %v", err))
		}

		idAddr, err := c.RPCClient.LookupId(ctx, fAddr, key)
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

func LookupIdAddress(ctx context.Context, addr string, c *fvm.LotusClient) (*address.Address, error) {
	var err error
	var fAddr address.Address

	fAddr, err = address.NewFromString(addr)
	if err != nil {
		ethAddr, err := ethtypes.ParseEthAddress(addr)

		fAddr, err = ethAddr.ToFilecoinAddress()
		if err != nil {
			return nil, fmt.Errorf("unable to get Filecoin address: %v", err)
		}
	}

	key, err := c.RPCClient.GetTipSetKey(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get tipset key for chain head: %v", err)
	}

	idAddr, err := c.RPCClient.LookupId(ctx, fAddr, key)
	if err != nil {
		return nil, fmt.Errorf("unable to get find actor ID for specified address: %v", err)
	}

	return idAddr, nil
}
