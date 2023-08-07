package staking

import (
	"github.com/spf13/cobra"
)

var StakingCmd = &cobra.Command{
	Use:   "staking",
	Short: "Interact with Liquid Staking contract to pledge FIL",
	Long:  `This section provides functionality to interact with the Liquid Staking contract`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
