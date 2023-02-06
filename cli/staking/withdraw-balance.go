package staking

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"collective-go-sdk/utils"
	"fmt"
	"math/big"

	"github.com/spf13/cobra"
)

var (
	amount int64
	miner  string
)

func withdrawBalance(miner string, amount int64, run bool) (string, error) {
	bytesAddr := utils.ConvertAddress(miner)
	withdrawAmt := big.NewInt(amount)

	config, err := config.LoadConfig("./config")
	if err != nil {
		return "", err
	}

	client, err := fvm.NewLotusClient(config)
	tx, err := client.WithdrawBalance(bytesAddr, withdrawAmt, run)

	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

var withdrawBalanceCmd = &cobra.Command{
	Use:   "withdraw-balance",
	Short: "Withdraw balance from the Miner actor",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := withdrawBalance(miner, amount, run); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	withdrawBalanceCmd.Flags().StringVarP(&miner, "miner", "m", "", "Miner address (filecoin)")
	withdrawBalanceCmd.Flags().Int64VarP(&amount, "amount", "a", 0, "Amount to withdraw")
	withdrawBalanceCmd.Flags().BoolVarP(&run, "run", "r", true, "Execute transaction")

	if err := withdrawBalanceCmd.MarkFlagRequired("miner"); err != nil {
		fmt.Println(err)
	}

	if err := withdrawBalanceCmd.MarkFlagRequired("amount"); err != nil {
		fmt.Println(err)
	}

	StakingCmd.AddCommand(withdrawBalanceCmd)
}
