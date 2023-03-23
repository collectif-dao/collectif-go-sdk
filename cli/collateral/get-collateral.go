package collateral

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"collective-go-sdk/utils"
	"context"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	address string
)

func getCollateral(address string) (string, string, error) {
	bytesAddr := utils.ConvertAddress(address)
	fmt.Println("Miner Address in bytes: ", hex.EncodeToString(bytesAddr))

	config, err := config.LoadConfig("./config")
	if err != nil {
		return "", "", err
	}

	ctx := context.Background()
	client, err := fvm.NewLotusClient(ctx, config, fvm.FSKeyStore)
	if err != nil {
		return "", "", err
	}

	available, err := client.GetAvailableCollateral(bytesAddr)
	if err != nil {
		return "", "", err
	}

	return available.String(), "0", nil
}

var getCollateralCmd = &cobra.Command{
	Use:   "get-collateral",
	Short: "Returns available and locked collateral amounts for Storage Provider",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if available, locked, err := getCollateral(address); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Available collateral: ", available)
			fmt.Println("Locked collateral: ", locked)
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
