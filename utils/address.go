package utils

import (
	"fmt"

	fAddress "github.com/filecoin-project/go-address"
)

// convertAddress for converting string address to bytes for Solidity interactions
func ConvertAddress(address string) []byte {
	addr, err := fAddress.NewFromString(address)

	if err != nil {
		panic(fmt.Sprintf("go-address is unable to initiate Address struct: %v", err))
	}

	return addr.Bytes()
}

func EncodePacked(address string) []byte {
	return encodePacked(encodeBytesString(address))
}
