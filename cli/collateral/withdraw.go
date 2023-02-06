package collateral

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"fmt"

	"math/big"

	"github.com/spf13/cobra"
)

func withdrawCollateral(amount int, run bool) (string, error) {
	withdrawal := big.NewInt(int64(amount))

	config, err := config.LoadConfig("./config")
	if err != nil {
		return "", err
	}

	client, err := fvm.NewLotusClient(config)

	tx, err := client.Withdraw(withdrawal, run)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

var withdrawCmd = &cobra.Command{
	Use:   "withdraw",
	Short: "Withdraw collateral from the StorageProviderCollateral contract",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := withdrawCollateral(amount, run); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	withdrawCmd.Flags().IntVarP(&amount, "amount", "a", 0, "Deposit amount")
	withdrawCmd.Flags().BoolVarP(&run, "run", "r", true, "Execute transaction")

	if err := withdrawCmd.MarkFlagRequired("amount"); err != nil {
		fmt.Println(err)
	}

	CollateralCmd.AddCommand(withdrawCmd)
}
