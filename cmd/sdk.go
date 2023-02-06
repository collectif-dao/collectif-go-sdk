package main

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"collective-go-sdk/utils"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// fmt.Println("Hello, Modules!")

	// cli.Execute()

	var minerAddr string = "t2vb6iahjntzweoxb7ozhond4jalwf5azy2xzk2oa"

	firstAddr := utils.ConvertAddress(minerAddr)
	fmt.Printf("minerAddress []uint8 :\t %v \n", firstAddr)

	str := hex.EncodeToString(firstAddr)
	fmt.Printf("minerAddress hex string:\t %v \n", str)

	config, err := config.LoadConfig("./config")
	if err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	client, err := fvm.NewLotusClient(config)
	if err != nil {
		panic(fmt.Errorf("invalid FVM client initialization: %s", err))
	}

	totalAssets, err := client.TotalAssets()
	if err != nil {
		panic(fmt.Errorf("invalid FVM TotalAssets call: %s", err))
	}
	fmt.Println(totalAssets.String())

	fees, err := client.GetPledgeFees()
	if err != nil {
		panic(fmt.Errorf("invalid FVM TotalAssets call: %s", err))
	}
	fmt.Println(fees.String())

	var minerBytes = utils.ConvertAddress(minerAddr)
	var fil = big.NewInt(1000000000000000000)
	var allocationLimit = fil.Mul(fil, big.NewInt(1000))
	var period = big.NewInt(131548)

	tx, err := client.Register(minerBytes, common.HexToAddress(config.LiquidStaking), allocationLimit, period, false)
	if err != nil {
		panic(fmt.Errorf("invalid FVM Register call: %s", err))
	}

	// tx, err := client.ChangeBeneficiaryAddress(common.HexToAddress(config.LiquidStaking), false)
	// if err != nil {
	// 	panic(fmt.Errorf("invalid ChangeBeneficiaryAddress call: %s", err))
	// }

	fmt.Println(hex.EncodeToString(tx.Data()))

	// tx, err = client.UpdateRecord(epoch, preCommitDeposit, initialPledge, true)

	// if err != nil {
	// 	panic(fmt.Errorf("invalid FVM TotalAssets call: %s", err))
	// }

	// fmt.Println(hex.EncodeToString(tx.Data()))

	// fmt.Printf(config.LiquidStaking)
}
