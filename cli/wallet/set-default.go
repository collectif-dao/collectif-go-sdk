package wallet

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"

	"github.com/spf13/cobra"
)

func setDefaultWallet(addr string) error {
	config, err := config.LoadConfig("./config")
	if err != nil {
		return err
	}

	ctx := context.Background()

	client, err := fvm.NewLotusClient(ctx, config, fvm.FSKeyStore)
	if err != nil {
		return err
	}

	fAddr, err := address.NewFromString(addr)
	if err != nil {
		return err
	}

	err = client.MessageSigner.Wallet.SetDefault(fAddr)
	if err != nil {
		return err
	}

	return nil
}

var setDefaultCmd = &cobra.Command{
	Use:   "set-default",
	Short: "Set default wallet for SDK",
	Long:  ``,
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
