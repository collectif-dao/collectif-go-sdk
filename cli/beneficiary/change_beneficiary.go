package beneficiary

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"context"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var (
	pool string
)

func changeBeneficiaryAddress(pool string) (string, error) {
	var poolAddr common.Address
	ctx := context.Background()

	config, err := config.LoadConfig("./config")
	if err != nil {
		return "", err
	}

	if pool == "" {
		poolAddr = common.HexToAddress(config.LiquidStaking)
	} else {
		poolAddr = common.HexToAddress(pool)
	}

	client, err := fvm.NewLotusClient(ctx, config, fvm.FSKeyStore)
	if err != nil {
		return "", err
	}

	abi, err := client.RegistryABI.GetAbi()
	if err != nil {
		return "", err
	}

	callData, err := abi.Pack("changeBeneficiaryAddress", poolAddr)

	return hex.EncodeToString(callData), nil
}

var ChangeBeneficiaryCmd = &cobra.Command{
	Use:   "change-beneficiary",
	Short: "Change beneficiary address in the Collectif DAO protocol",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := changeBeneficiaryAddress(pool); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	ChangeBeneficiaryCmd.Flags().StringVarP(&pool, "staking", "s", "", "Liquid Staking pool address (ethereum address)")

	if err := ChangeBeneficiaryCmd.MarkFlagRequired("staking"); err != nil {
		fmt.Println(err)
	}
}
