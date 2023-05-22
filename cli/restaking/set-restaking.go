package restaking

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var (
	ratio int
	addr  string
	run   bool
)

func setRestaking(ratio int, addr string, run bool) (string, error) {
	config, err := config.LoadConfig("./config")
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	client, err := fvm.NewLotusClient(ctx, config, fvm.FSKeyStore)
	if err != nil {
		return "", err
	}

	restakingRatio := big.NewInt(int64(ratio))
	valid := common.IsHexAddress(addr)
	var ethAddress common.Address

	if valid {
		ethAddress = common.HexToAddress(addr)
	}

	abi, err := client.RegistryABI.GetAbi()
	if err != nil {
		return "", err
	}

	callData, err := abi.Pack("setRestaking", restakingRatio, ethAddress)

	return hex.EncodeToString(callData), nil
}

var SetRestakingCmd = &cobra.Command{
	Use:   "set-restaking",
	Short: "Update restaking parameters for SP in the Collectif DAO protocol",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := setRestaking(ratio, addr, run); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	SetRestakingCmd.Flags().IntVarP(&ratio, "ratio", "v", 0, "Restaking ratio")
	SetRestakingCmd.Flags().StringVarP(&addr, "address", "a", "0x", "Filecoin address to receive clFIL tokens")
	SetRestakingCmd.Flags().BoolVarP(&run, "run", "r", true, "Execute transaction")

	if err := SetRestakingCmd.MarkFlagRequired("ratio"); err != nil {
		fmt.Println(err)
	}

	RestakingCmd.AddCommand(SetRestakingCmd)
}
