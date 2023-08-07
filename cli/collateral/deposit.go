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

var (
	amount int64
	run    bool
)

func depositCollateral(amount int64, run bool) (*fvm.MessageResponse, error) {
	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, keystore.FSKeyStore, "./")
	if err != nil {
		return nil, err
	}

	value := utils.GetAttoFilFromFIL(amount)
	msg, err := sdk.Client.Deposit(ctx, value, run)
	if err != nil {
		return msg, err
	}

	return msg, nil
}

var depositCmd = &cobra.Command{
	Use:   "deposit",
	Short: "Deposit collateral into StorageProviderCollateral contract",
	Long:  `To deposit FIL into the collateral system provide an amount of FIL with -a (or --amount) flag`,
	Run: func(cmd *cobra.Command, args []string) {

		if msg, err := depositCollateral(amount, run); err != nil {
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
	depositCmd.Flags().Int64VarP(&amount, "amount", "a", 0, "Deposit amount")
	depositCmd.Flags().BoolVarP(&run, "execute", "e", true, "Execute transaction")

	if err := depositCmd.MarkFlagRequired("amount"); err != nil {
		fmt.Println(err)
	}

	CollateralCmd.AddCommand(depositCmd)
}
