package staking

import (
	"collective-go-sdk/fvm"
	"collective-go-sdk/keystore"
	"collective-go-sdk/sdk"
	"collective-go-sdk/utils"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	amount int64
	run    bool
)

func pledge(amount int64, run bool) (*fvm.MessageResponse, error) {
	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, fvm.DefaultNetwork, keystore.FSKeyStore, "./")
	if err != nil {
		return nil, err
	}

	value := utils.GetAttoFilFromFIL(amount)
	msg, err := sdk.Client.Pledge(ctx, value, run)
	if err != nil {
		return msg, err
	}

	return msg, nil
}

var pledgeCmd = &cobra.Command{
	Use:   "pledge",
	Short: "Pledge part of daily allocation to cover initial pledge for Storage Provider",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if msg, err := pledge(amount, run); err != nil {
			fmt.Println(err)

			fmt.Println("Message calldata: ", msg.Data)
		} else {
			if run {
				fmt.Println("Executed message with: ", msg.Message, " CID")
				fmt.Println("Returned: ", msg.Receipt.Return)
				fmt.Println("Gas spent: ", msg.Receipt.GasUsed)
			}

			fmt.Println("Message calldata: ", msg.Data)
		}
	},
}

func init() {
	pledgeCmd.Flags().Int64VarP(&amount, "amount", "a", 0, "Amount of FIL to pledge (not attoFIL)")
	pledgeCmd.Flags().BoolVarP(&run, "execute", "e", true, "Execute transaction")

	if err := pledgeCmd.MarkFlagRequired("amount"); err != nil {
		fmt.Println(err)
	}

	StakingCmd.AddCommand(pledgeCmd)
}
