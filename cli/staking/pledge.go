package staking

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
	amount    int64
	run       bool
	minerAddr string
)

func pledge(amount int64, minerAddr string, run bool) (*fvm.MessageResponse, error) {
	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, keystore.FSKeyStore, "./")
	if err != nil {
		return nil, err
	}

	value := utils.GetAttoFilFromFIL(amount)

	minerId := utils.GetIdAddress(ctx, minerAddr, sdk.Client)

	msg, err := sdk.Client.Pledge(ctx, value, minerId, run)
	if err != nil {
		return msg, err
	}

	return msg, nil
}

var pledgeCmd = &cobra.Command{
	Use:   "pledge",
	Short: "Pledge part of daily allocation to cover initial pledge for Storage Provider",
	Long:  `This operation is pledging FIL from the liquid staking pool to your miner actor. Make sure to provide an amount of FIL with -a flag, and miner actor address (or actor ID) with -m flag`,
	Run: func(cmd *cobra.Command, args []string) {

		if msg, err := pledge(amount, minerAddr, run); err != nil {
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
	pledgeCmd.Flags().StringVarP(&minerAddr, "minerAddr", "m", "", "Miner actor address (either ID address or actor address)")
	pledgeCmd.Flags().Int64VarP(&amount, "amount", "a", 0, "Amount of FIL to pledge (not attoFIL)")
	pledgeCmd.Flags().BoolVarP(&run, "execute", "e", true, "Execute transaction")

	if err := pledgeCmd.MarkFlagRequired("amount"); err != nil {
		fmt.Println(err)
	}

	if err := pledgeCmd.MarkFlagRequired("minerAddr"); err != nil {
		fmt.Println(err)
	}

	StakingCmd.AddCommand(pledgeCmd)
}
