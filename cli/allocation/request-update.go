package allocation

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
	totalLimit int64
	dailyLimit int64
	run        bool
)

func requestUpdate(totalLimit int64, dailyLimit int64, run bool) (*fvm.MessageResponse, error) {
	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, keystore.FSKeyStore, "./")
	if err != nil {
		return nil, err
	}

	allocationLimit := utils.GetAttoFilFromFIL(totalLimit)
	dailyAllocation := utils.GetAttoFilFromFIL(dailyLimit)

	msg, err := sdk.Client.RequestAllocationLimitUpdate(ctx, allocationLimit, dailyAllocation, run)
	if err != nil {
		return msg, err
	}

	return msg, nil
}

var RequestAllocationCmd = &cobra.Command{
	Use:   "request-allocation-update",
	Short: "Request allocation update for SP in the Collectif DAO protocol",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if msg, err := requestUpdate(totalLimit, dailyLimit, run); err != nil {
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
	RequestAllocationCmd.Flags().Int64VarP(&totalLimit, "allocation-limit", "l", 0, "Total allocation limit for update")
	RequestAllocationCmd.Flags().Int64VarP(&dailyLimit, "daily-allocation", "d", 0, "Daily allocation limit for update")
	RequestAllocationCmd.Flags().BoolVarP(&run, "execute", "e", true, "Execute transaction")

	if err := RequestAllocationCmd.MarkFlagRequired("allocation-limit"); err != nil {
		fmt.Println(err)
	}
	if err := RequestAllocationCmd.MarkFlagRequired("daily-allocation"); err != nil {
		fmt.Println(err)
	}

	AllocationCmd.AddCommand(RequestAllocationCmd)
}
