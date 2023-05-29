package restaking

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
	ownerId string
)

func getRestakingInfo(ownerId string) (*fvm.SPRestaking, error) {
	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, keystore.FSKeyStore, "./")
	if err != nil {
		return nil, err
	}

	if ownerId == "" {
		ownerId = sdk.Client.Address.String()
	}

	idAddr := fUtils.GetIdAddress(ctx, ownerId, sdk.Client)

	restaking, err := sdk.Client.GetRestaking(ctx, idAddr)
	if err != nil {
		return nil, err
	}

	return restaking, nil
}

var getRestaking = &cobra.Command{
	Use:   "get-restaking",
	Short: "Returns the restaking information for the Storage Provider",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if restaking, err := getRestakingInfo(ownerId); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Restaking ratio: ", restaking.RestakingRatio.String())
			fmt.Println("Restaking address: ", restaking.RestakingAddress.String())
		}
	},
}

func init() {
	getRestaking.Flags().StringVarP(&ownerId, "ownerId", "o", "", "Storage Provider ownerID (filecoin)")
	RestakingCmd.AddCommand(getRestaking)
}
