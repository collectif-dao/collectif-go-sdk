package beneficiary

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
	minerId     string
	beneficiary string
	quota       int64
	expiration  int64
	run         bool
)

func changeBeneficiaryAddress(minerId string, beneficiary string, quota int64, expiration int64, run bool) (*fvm.MessageResponse, error) {
	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, keystore.FSKeyStore, "./")
	if err != nil {
		return nil, err
	}

	qA := utils.GetAttoFilFromFIL(quota)

	mID, err := utils.LookupIdAddress(ctx, minerId, sdk.Client)
	if err != nil {
		return nil, err
	}

	bID, err := utils.LookupIdAddress(ctx, beneficiary, sdk.Client)
	if err != nil {
		return nil, err
	}

	res, err := sdk.Client.ChangeBeneficiaryAddress(ctx, mID, bID, qA, expiration, run)
	if err != nil {
		return res, err
	}

	return res, nil
}

var ChangeBeneficiaryCmd = &cobra.Command{
	Use:   "change-beneficiary",
	Short: "Change beneficiary address in the Collectif DAO protocol",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if msg, err := changeBeneficiaryAddress(minerId, beneficiary, quota, expiration, run); err != nil {
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
	ChangeBeneficiaryCmd.Flags().StringVarP(&minerId, "minerId", "m", "", "Miner actor address (either ID address or actor address)")
	ChangeBeneficiaryCmd.Flags().StringVarP(&beneficiary, "beneficiary", "b", "", "New beneficiary address")
	ChangeBeneficiaryCmd.Flags().Int64VarP(&quota, "quota", "q", 0, "New beneficiary quota")
	ChangeBeneficiaryCmd.Flags().Int64VarP(&expiration, "expiration", "e", 0, "New expiration of beneficiary term")
	ChangeBeneficiaryCmd.Flags().BoolVarP(&run, "run", "r", true, "Run transaction")

	if err := ChangeBeneficiaryCmd.MarkFlagRequired("minerId"); err != nil {
		fmt.Println(err)
	}

	if err := ChangeBeneficiaryCmd.MarkFlagRequired("beneficiary"); err != nil {
		fmt.Println(err)
	}

	if err := ChangeBeneficiaryCmd.MarkFlagRequired("quota"); err != nil {
		fmt.Println(err)
	}

	if err := ChangeBeneficiaryCmd.MarkFlagRequired("expiration"); err != nil {
		fmt.Println(err)
	}
}
