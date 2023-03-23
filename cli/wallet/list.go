package wallet

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func list() ([]string, error) {
	addrList := make([]string, 0, 2)
	config, err := config.LoadConfig("./config")
	if err != nil {
		return addrList, err
	}

	ctx := context.Background()

	client, err := fvm.NewLotusClient(ctx, config, fvm.FSKeyStore)
	if err != nil {
		return addrList, err
	}

	addrs, err := client.MessageSigner.Wallet.WalletList(ctx)
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
