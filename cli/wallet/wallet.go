package wallet

import (
	"github.com/spf13/cobra"
)

var WalletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Interact with Filecoin/Ethereum wallet",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
