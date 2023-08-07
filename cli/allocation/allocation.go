package allocation

import (
	"github.com/spf13/cobra"
)

var AllocationCmd = &cobra.Command{
	Use:   "allocation",
	Short: "Allocation section for Storage Provider",
	Long:  `Allocation section helps you understand how much FIL is available for pledges for your miner actor. Also you can request an update to change your FIL allocation from the liquid staking pool`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
