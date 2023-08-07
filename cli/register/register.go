package register

import (
	"collective-go-sdk/fvm"
	"collective-go-sdk/keystore"
	"collective-go-sdk/sdk"
	"collective-go-sdk/utils"
	"context"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

var (
	minerAddr  string
	totalLimit int64
	dailyLimit int64
	run        bool
)

func registerStorageProvider(minerAddr string, totalLimit int64, dailyLimit int64, run bool) (*fvm.MessageResponse, error) {
	if totalLimit == 0 || dailyLimit == 0 {
		return nil, xerrors.Errorf("Incorrect params for Register transaction")
	}

	dailyAllocation := utils.GetAttoFilFromFIL(dailyLimit)
	allocationLimit := utils.GetAttoFilFromFIL(totalLimit)

	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, keystore.FSKeyStore, "./")
	if err != nil {
		return nil, err
	}

	minerId := utils.GetIdAddress(ctx, minerAddr, sdk.Client)
	msg, err := sdk.Client.Register(ctx, minerId, allocationLimit, dailyAllocation, run)
	if err != nil {
		return msg, err
	}

	return msg, nil
}

var RegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Register Storage Provider in the Collectif DAO protocol",
	Long:  `Register Storage Provider in the Collectif DAO protocol. Please provide your miner actor address with -m (or --minerAddr) flag, your total desired amount of FIL from the liquid staking pool with -l (or --allocation-limit) flag, and daily FIL pledge with -d (or --daily-allocation) flag. After the registration step you would need to wait for Collectif DAO to onboard you and after that you can change your beneficiary address to Reward Collector smart contract`,
	Run: func(cmd *cobra.Command, args []string) {

		if msg, err := registerStorageProvider(minerAddr, totalLimit, dailyLimit, run); err != nil {
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
	RegisterCmd.Flags().StringVarP(&minerAddr, "minerAddr", "m", "", "Miner actor address (either ID address or actor address)")
	RegisterCmd.Flags().Int64VarP(&totalLimit, "allocation-limit", "l", 0, "Total FIL allocation for pledge")
	RegisterCmd.Flags().Int64VarP(&dailyLimit, "daily-allocation", "d", 0, "Daily FIL allocation for pledge")
	RegisterCmd.Flags().BoolVarP(&run, "execute", "e", true, "Execute transaction")

	if err := RegisterCmd.MarkFlagRequired("minerAddr"); err != nil {
		fmt.Println(err)
	}
}
