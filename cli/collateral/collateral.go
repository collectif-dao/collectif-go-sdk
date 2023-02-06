package collateral

import (
	"github.com/spf13/cobra"
)

var CollateralCmd = &cobra.Command{
	Use:   "collateral",
	Short: "This is the section for interacting with StorageProviderCollateral contract",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
