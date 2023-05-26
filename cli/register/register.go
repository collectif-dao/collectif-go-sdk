package register

import (
	"collective-go-sdk/fvm"
	"collective-go-sdk/keystore"
	"collective-go-sdk/sdk"
	"collective-go-sdk/utils"
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

var (
	minerId    uint64
	totalLimit int64
	dailyLimit int64
	run        bool
)

func registerStorageProvider(minerId uint64, totalLimit int64, dailyLimit int64, run bool) (*fvm.MessageResponse, error) {
	if totalLimit == 0 || dailyLimit == 0 {
		return nil, xerrors.Errorf("Incorrect params for Register transaction")
	}

	dailyAllocation := utils.GetAttoFilFromFIL(dailyLimit)
	allocationLimit := utils.GetAttoFilFromFIL(totalLimit)

	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, fvm.DefaultNetwork, keystore.FSKeyStore, "./")
	if err != nil {
		return nil, err
	}

	msg, err := sdk.Client.Register(ctx, minerId, allocationLimit, dailyAllocation, run)
	if err != nil {
		return msg, err
	}

	return msg, nil
}

var RegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Register Storage Provider in the Collectif DAO protocol",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if msg, err := registerStorageProvider(minerId, totalLimit, dailyLimit, run); err != nil {
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
	RegisterCmd.Flags().Uint64VarP(&minerId, "minerId", "m", 0, "Storage Provider miner id (not filecoin address)")
	RegisterCmd.Flags().Int64VarP(&totalLimit, "allocation-limit", "l", 0, "Total FIL allocation for pledge")
	RegisterCmd.Flags().Int64VarP(&dailyLimit, "daily-allocation", "d", 0, "Daily FIL allocation for pledge")
	RegisterCmd.Flags().BoolVarP(&run, "execute", "e", true, "Execute transaction")

	if err := RegisterCmd.MarkFlagRequired("miner"); err != nil {
		fmt.Println(err)
	}
}
