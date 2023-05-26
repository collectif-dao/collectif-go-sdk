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
	sdk, err := sdk.NewCollectifSDK(ctx, fvm.DefaultNetwork, keystore.FSKeyStore, "./")
	if err != nil {
		return nil, err
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
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if collateral, err := getCollateral(address); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Available collateral: ", collateral.AvailableCollateral)
			fmt.Println("Locked collateral: ", collateral.LockedCollateral)
		}
	},
}

func init() {
	getCollateralCmd.Flags().StringVarP(&address, "address", "a", "", "Storage Provider address (filecoin)")

	if err := getCollateralCmd.MarkFlagRequired("address"); err != nil {
		fmt.Println(err)
	}

	CollateralCmd.AddCommand(getCollateralCmd)
}
