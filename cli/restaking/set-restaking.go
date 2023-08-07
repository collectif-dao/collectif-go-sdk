package restaking

import (
	"collective-go-sdk/fvm"
	"collective-go-sdk/keystore"
	"collective-go-sdk/sdk"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var (
	ratio int
	addr  string
	run   bool
)

func setRestaking(ratio int, addr string, run bool) (*fvm.MessageResponse, error) {
	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, keystore.FSKeyStore, "./")
	if err != nil {
		return nil, err
	}

	var ethAddress common.Address
	if addr == "0x" {
		ethAddress = *sdk.Client.EthAddress
	} else {
		valid := common.IsHexAddress(addr)

		if valid {
			ethAddress = common.HexToAddress(addr)
		}
	}

	var restakingRatio *big.Int
	if ratio > 100 && ratio < 10000 {
		restakingRatio = big.NewInt(int64(ratio))
	}

	restakingRatio = big.NewInt(int64(ratio * 100))

	msg, err := sdk.Client.SetRestaking(ctx, restakingRatio, ethAddress, run)
	if err != nil {
		return msg, err
	}

	return msg, nil
}

var SetRestakingCmd = &cobra.Command{
	Use:   "set-restaking",
	Short: "Update restaking parameters for SP in the Collectif DAO protocol",
	Long:  `IF you wish to allow restaking of some of your mining rewards, provide a percentage of rewards that you want to stake in the Liquid Staking pool. Restaking helps you to minimize the cost of capital, and earn some portion of rewards from other SPs. During restaking you'll receive clFIL token. Make sure to provide an f4 or 0x Filecoin wallet with -a flag and your restaking ratio with -r flag.`,
	Run: func(cmd *cobra.Command, args []string) {

		if msg, err := setRestaking(ratio, addr, run); err != nil {
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
	SetRestakingCmd.Flags().IntVarP(&ratio, "ratio", "r", 0, "Restaking ratio in percentage (up to 100%)")
	SetRestakingCmd.Flags().StringVarP(&addr, "address", "a", "0x", "Filecoin address to receive clFIL tokens (f4/0x address type)")
	SetRestakingCmd.Flags().BoolVarP(&run, "execute", "e", true, "Execute transaction")

	if err := SetRestakingCmd.MarkFlagRequired("ratio"); err != nil {
		fmt.Println(err)
	}

	RestakingCmd.AddCommand(SetRestakingCmd)
}
