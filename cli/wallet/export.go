package wallet

import (
	"collective-go-sdk/keystore"
	"collective-go-sdk/sdk"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/spf13/cobra"
)

var addr string

func exportWallet(addr string) ([]string, error) {
	emptyKey := make([]string, 0, 0)

	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, keystore.FSKeyStore, "./")
	if err != nil {
		return nil, err
	}

	lAddr, err := address.NewFromString(addr)
	if err != nil {
		return emptyKey, err
	}

	ki, err := sdk.Client.MessageSigner.Wallet.WalletExport(ctx, lAddr)
	if err != nil {
		return emptyKey, err
	}

	b, err := json.Marshal(ki)
	if err != nil {
		return emptyKey, err
	}

	keyList := []string{hex.EncodeToString(b), hex.EncodeToString(ki.PrivateKey)}

	return keyList, nil
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export Filecoin wallet from SDK",
	Long:  `Export a private key of your Filecoin wallet by providing your wallet address with -a flag`,
	Run: func(cmd *cobra.Command, args []string) {

		if res, err := exportWallet(addr); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Private Key: ", res[0])
			fmt.Println("Key Info hex: ", res[1])
		}
	},
}

func init() {
	exportCmd.Flags().StringVarP(&addr, "address", "a", "", "Filecoin wallet address")

	if err := exportCmd.MarkFlagRequired("address"); err != nil {
		fmt.Println(err)
	}

	WalletCmd.AddCommand(exportCmd)
}
