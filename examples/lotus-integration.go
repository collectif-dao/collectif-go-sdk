package main

import (
	"collective-go-sdk/keystore"
	"collective-go-sdk/sdk"
	"collective-go-sdk/utils"
	"context"
)

var minerId = uint64(1000)

func main() {
	ctx := context.Background()

	ksType := keystore.FSKeyStore
	configPath := "./"

	sdk, err := sdk.NewCollectifSDK(ctx, ksType, configPath)
	if err != nil {
		panic(err)
	}

	pledgeAmount := utils.GetAttoFilFromFIL(1000)

	sdk.Client.Pledge(ctx, pledgeAmount, minerId, true)
}
