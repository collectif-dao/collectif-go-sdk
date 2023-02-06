package cli

import (
	"collective-go-sdk/cli/collateral"
	"collective-go-sdk/cli/register"
	"collective-go-sdk/cli/staking"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "collective-go-sdk",
	Short: "Storage Provider client for Collectif Liquid Staking protocol",
	Long: `Collective-Go-SDK is the main client for Filecoin Storage Providers
	to interact with Collective Liquid Staking protocol it wraps
	around of all the complexities of joining Collectif protocol, pledging capital from the liquid staking pool
	into miner's Initial Pledge, and manage SPs collateral.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(initConfig)

	fmt.Println()
	rootCmd.AddCommand(collateral.CollateralCmd)
	rootCmd.AddCommand(register.RegisterCmd)
	rootCmd.AddCommand(staking.StakingCmd)
}

// func initConfig() {
// 	if cfgFile != "" {
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		home, err := os.UserHomeDir()
// 		cobra.CheckErr(err)
// 		viper.AddConfigPath(home)
// 		viper.SetConfigType("yaml")
// 		viper.SetConfigName(".toolbox")
// 	}

// 	viper.AutomaticEnv()

// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
// 	}
// }
