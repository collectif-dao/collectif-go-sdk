package wallet

import (
	"collective-go-sdk/fvm"
	"collective-go-sdk/keystore"
	"collective-go-sdk/sdk"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func getDefaultWallet() (string, error) {
	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, fvm.DefaultNetwork, keystore.FSKeyStore, "./")
	if err != nil {
		return "", err
	}

	address, err := sdk.Client.MessageSigner.Wallet.GetDefault()
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
