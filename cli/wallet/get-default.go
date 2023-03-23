package wallet

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func getDefaultWallet() (string, error) {
	config, err := config.LoadConfig("./config")
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	client, err := fvm.NewLotusClient(ctx, config, fvm.FSKeyStore)
	if err != nil {
		return "", err
	}

	address, err := client.MessageSigner.Wallet.GetDefault()
	if err != nil {
		return "", err
	}

	return address.String(), nil
}

var getDefaultCmd = &cobra.Command{
	Use:   "get-default",
	Short: "Get default wallet used by SDK",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := getDefaultWallet(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	WalletCmd.AddCommand(getDefaultCmd)
}
