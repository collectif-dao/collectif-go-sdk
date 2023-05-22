package restaking

import (
	"github.com/spf13/cobra"
)

var RestakingCmd = &cobra.Command{
	Use:   "restaking",
	Short: "Restaking section of the Storage Provider",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
