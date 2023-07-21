package allocation

import (
	"github.com/spf13/cobra"
)

var AllocationCmd = &cobra.Command{
	Use:   "allocation",
	Short: "Allocation section for Storage Provider",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
