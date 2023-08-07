package wallet

import (
	"collective-go-sdk/keystore"
	"collective-go-sdk/sdk"
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"

	"github.com/spf13/cobra"
)

func setDefaultWallet(addr string) error {
	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, keystore.FSKeyStore, "./")
	if err != nil {
		return err
	}

	fAddr, err := address.NewFromString(addr)
	if err != nil {
		return err
	}

	err = sdk.Client.MessageSigner.Wallet.SetDefault(fAddr)
	if err != nil {
		return err
	}

	return nil
}

var setDefaultCmd = &cobra.Command{
	Use:   "set-default",
	Short: "Set default wallet for SDK",
	Long:  `To set a wallet as default please provide your Filecoin address with -a flag. Make sure to import or create the wallet in the first place.`,
	Run: func(cmd *cobra.Command, args []string) {

		if err := setDefaultWallet(addr); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Succesfully set", addr, "as default wallet")
		}
	},
}

func init() {
	setDefaultCmd.Flags().StringVarP(&addr, "address", "a", "", "Filecoin wallet address")

	if err := setDefaultCmd.MarkFlagRequired("address"); err != nil {
		fmt.Println(err)
	}

	WalletCmd.AddCommand(setDefaultCmd)
}
