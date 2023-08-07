package wallet

import (
	"collective-go-sdk/keystore"
	"collective-go-sdk/sdk"
	"context"
	"fmt"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/spf13/cobra"
)

var key string

func newWallet(key string) (string, error) {
	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, keystore.FSKeyStore, "./")
	if err != nil {
		return "", err
	}

	address, err := sdk.Client.MessageSigner.Wallet.WalletNew(ctx, types.KeyType(key))
	if err != nil {
		return "", err
	}

	if isDefault {
		if err := sdk.Client.MessageSigner.Wallet.SetDefault(address); err != nil {
			return "", fmt.Errorf("failed to set default key: %w", err)
		}
	}

	return address.String(), nil
}

var newWalletCmd = &cobra.Command{
	Use:   "new-wallet",
	Short: "Create new wallet for Collectif SDK.",
	Long:  `Create new wallet for Collectif SDK. Make sure to transfer ownership of your miner actor if you want to use this wallet further. If you already have a wallet that is an owner of your miner actors, it's better to import instead of creating new one.`,
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
	newWalletCmd.Flags().BoolVarP(&isDefault, "use-default", "d", true, "Set as default")
	WalletCmd.AddCommand(newWalletCmd)
}
