package wallet

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"context"
	"fmt"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/spf13/cobra"
)

var key string

func newWallet(key string) (string, error) {
	config, err := config.LoadConfig("./config")
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	client, err := fvm.NewLotusClient(ctx, config, fvm.FSKeyStore)
	if err != nil {
		return "", err
	}

	address, err := client.MessageSigner.Wallet.WalletNew(ctx, types.KeyType(key))
	if err != nil {
		return "", err
	}

	return address.String(), nil
}

var newWalletCmd = &cobra.Command{
	Use:   "new-wallet",
	Short: "Create new wallet for Collectif SDK",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := newWallet(key); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	newWalletCmd.Flags().StringVarP(&key, "key", "k", "secp256k1", "Filecoin wallet key type: bls or secp256k1")
	WalletCmd.AddCommand(newWalletCmd)
}
