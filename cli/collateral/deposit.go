package collateral

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"fmt"

	"math/big"

	"github.com/spf13/cobra"
)

var (
	amount int
	run    bool
)

func depositCollateral(amount int, run bool) (string, error) {
	depositAmt := big.NewInt(int64(amount))

	config, err := config.LoadConfig("./config")
	if err != nil {
		return "", err
	}

	client, err := fvm.NewLotusClient(config)

	tx, err := client.Deposit(depositAmt, run)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

var depositCmd = &cobra.Command{
	Use:   "deposit",
	Short: "Deposit collateral into StorageProviderCollateral contract",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := depositCollateral(amount, run); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	depositCmd.Flags().IntVarP(&amount, "amount", "a", 0, "Deposit amount")
	depositCmd.Flags().BoolVarP(&run, "run", "r", true, "Execute transaction")

	if err := depositCmd.MarkFlagRequired("amount"); err != nil {
		fmt.Println(err)
	}

	CollateralCmd.AddCommand(depositCmd)
}
