package wallet

import (
	"collective-go-sdk/fvm"
	"collective-go-sdk/keystore"
	"collective-go-sdk/sdk"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func list() ([]string, error) {
	addrList := make([]string, 0, 2)

	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, fvm.DefaultNetwork, keystore.FSKeyStore, "./")
	if err != nil {
		return addrList, err
	}

	addrs, err := sdk.Client.MessageSigner.Wallet.WalletList(ctx)
	if err != nil {
		return addrList, err
	}

	if len(addrs) == 0 {
		return addrList, nil
	}

	for _, addr := range addrs {
		addrList = append(addrList, addr.String())
	}

	return addrList, nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List wallets used by SDK",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := list(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	WalletCmd.AddCommand(listCmd)
}
