package collateral

import (
	"collective-go-sdk/fvm"
	"collective-go-sdk/keystore"
	"collective-go-sdk/sdk"
	fUtils "collective-go-sdk/utils"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	address string
)

func getCollateral(address string) (*fvm.CollateralInfo, error) {
	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, keystore.FSKeyStore, "./")
	if err != nil {
		return nil, err
	}

	if address == "" {
		address = sdk.Client.Address.String()
	}

	idAddr := fUtils.GetIdAddress(ctx, address, sdk.Client)

	collateral, err := sdk.Client.GetCollateral(ctx, idAddr)
	if err != nil {
		return nil, err
	}

	return collateral, nil
}

var getCollateralCmd = &cobra.Command{
	Use:   "get-collateral",
	Short: "Returns available and locked collateral amounts for Storage Provider",
	Long:  `Returns available and locked collateral amounts for Storage Provider. Available collateral is a portion that is not used in the collateral system and is available for withdrawals. Locked portion is backing the pledge your miner actor already got from the liquid staking pool.`,
	Run: func(cmd *cobra.Command, args []string) {

		if collateral, err := getCollateral(address); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Available collateral: ", fvm.AttoFIL2FIL_str(collateral.AvailableCollateral))
			fmt.Println("Locked collateral: ", fvm.AttoFIL2FIL_str(collateral.LockedCollateral))
		}
	},
}

func init() {
	getCollateralCmd.Flags().StringVarP(&address, "address", "a", "", "Storage Provider Owner ID address (filecoin)")
	CollateralCmd.AddCommand(getCollateralCmd)
}
