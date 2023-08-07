package collateral

import (
	"collective-go-sdk/fvm"
	"collective-go-sdk/keystore"
	"collective-go-sdk/sdk"
	"collective-go-sdk/utils"
	"context"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

func withdrawCollateral(amount int64, run bool) (*fvm.MessageResponse, error) {
	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, keystore.FSKeyStore, "./")
	if err != nil {
		return nil, err
	}

	value := utils.GetAttoFilFromFIL(amount)
	msg, err := sdk.Client.Withdraw(ctx, value, run)
	if err != nil {
		return msg, err
	}

	return msg, nil
}

var withdrawCmd = &cobra.Command{
	Use:   "withdraw",
	Short: "Withdraw collateral from the StorageProviderCollateral contract",
	Long:  `To withdraw FIL from the collateral system provide an amount of FIL with -a (or --amount) flag. The transaction would withdraw collateral only from the available balance, and locked portion of collateral would remain in the smart contract`,
	Run: func(cmd *cobra.Command, args []string) {

		if msg, err := withdrawCollateral(amount, run); err != nil {
			fmt.Println(err)

			fmt.Println("Message calldata: ", msg.Data)
		} else {
			if run {
				fmt.Println("Executed message with: ", msg.Message, " CID")
				fmt.Println("Returned: ", msg.Receipt.Return)
				fmt.Println("Gas spent: ", msg.Receipt.GasUsed)
			}

			fmt.Println("Message calldata: ", hex.EncodeToString(msg.Data))
		}
	},
}

func init() {
	withdrawCmd.Flags().Int64VarP(&amount, "amount", "a", 0, "Deposit amount")
	withdrawCmd.Flags().BoolVarP(&run, "execute", "e", true, "Execute transaction")

	if err := withdrawCmd.MarkFlagRequired("amount"); err != nil {
		fmt.Println(err)
	}

	CollateralCmd.AddCommand(withdrawCmd)
}
