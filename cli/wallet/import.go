package wallet

import (
	"collective-go-sdk/keystore"
	"collective-go-sdk/sdk"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/spf13/cobra"
)

var (
	privateKey string
	isDefault  bool
)

func importWallet(pk string, isDefault bool) error {
	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, keystore.FSKeyStore, "./")
	if err != nil {
		return err
	}

	var ki types.KeyInfo

	data, err := hex.DecodeString(strings.TrimSpace(pk))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &ki); err != nil {
		return err
	}

	addr, err := sdk.Client.MessageSigner.Wallet.WalletImport(ctx, &ki)
	if err != nil {
		return err
	}

	if isDefault {
		if err := sdk.Client.MessageSigner.Wallet.SetDefault(addr); err != nil {
			return fmt.Errorf("failed to set default key: %w", err)
		}
	}

	fmt.Printf("imported key %s successfully!\n", addr)

	return nil
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import Filecoin wallet for SDK",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if err := importWallet(privateKey, isDefault); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	importCmd.Flags().StringVarP(&privateKey, "private-key", "p", "", "Private Key")
	importCmd.Flags().BoolVarP(&isDefault, "use-default", "d", true, "Set as default")

	if err := importCmd.MarkFlagRequired("private-key"); err != nil {
		fmt.Println(err)
	}

	WalletCmd.AddCommand(importCmd)
}
