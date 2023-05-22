package restaking

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	fUtils "collective-go-sdk/fvm/utils"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ownerId string
)

func getRestakingInfo(ownerId string) (string, string, error) {
	config, err := config.LoadConfig("./config")
	if err != nil {
		return "", "", err
	}

	ctx := context.Background()
	client, err := fvm.NewLotusClient(ctx, config, fvm.FSKeyStore)
	if err != nil {
		return "", "", err
	}

	if ownerId == "" {
		ownerId = client.Address.String()
	}

	fmt.Print(ownerId)
	idAddr := fUtils.GetIdAddress(ctx, ownerId, client)

	restaking, err := client.GetRestaking(idAddr)
	if err != nil {
		return "", "", err
	}

	return restaking.RestakingRatio.String(), restaking.RestakingAddress.String(), nil
}

var getRestaking = &cobra.Command{
	Use:   "get-restaking",
	Short: "Returns the restaking information for the Storage Provider",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if ratio, addr, err := getRestakingInfo(ownerId); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Restaking ratio: ", ratio)
			fmt.Println("Locked collateral: ", addr)
		}
	},
}

func init() {
	getRestaking.Flags().StringVarP(&ownerId, "ownerId", "o", "", "Storage Provider ownerID (filecoin)")
	RestakingCmd.AddCommand(getRestaking)
}
