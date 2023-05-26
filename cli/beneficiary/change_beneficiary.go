package beneficiary

import (
	"collective-go-sdk/fvm"
	"collective-go-sdk/keystore"
	"collective-go-sdk/sdk"
	"context"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	pool string
)

func changeBeneficiaryAddress(pool string) (string, error) {
	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, fvm.DefaultNetwork, keystore.FSKeyStore, "./")
	if err != nil {
		return "", err
	}

	callData, err := sdk.Client.Registry.ABI.Pack("changeBeneficiaryAddress")

	return hex.EncodeToString(callData), nil
}

var ChangeBeneficiaryCmd = &cobra.Command{
	Use:   "change-beneficiary",
	Short: "Change beneficiary address in the Collectif DAO protocol",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := changeBeneficiaryAddress(pool); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	ChangeBeneficiaryCmd.Flags().StringVarP(&pool, "staking", "s", "", "Liquid Staking pool address (ethereum address)")

	if err := ChangeBeneficiaryCmd.MarkFlagRequired("staking"); err != nil {
		fmt.Println(err)
	}
}
