package allocation

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/spf13/cobra"
)

var (
	allocationLimit int
	dailyAllocation int
	run             bool
)

func requestUpdate(allocationLimit int, dailyAllocation int, run bool) (string, error) {
	config, err := config.LoadConfig("./config")
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	client, err := fvm.NewLotusClient(ctx, config, fvm.FSKeyStore)
	if err != nil {
		return "", err
	}

	limit := big.NewInt(int64(allocationLimit))
	daily := big.NewInt(int64(dailyAllocation))

	abi, err := client.RegistryABI.GetAbi()
	if err != nil {
		return "", err
	}

	callData, err := abi.Pack("requestAllocationUpdate", limit, daily)

	return hex.EncodeToString(callData), nil
}

var RequestAllocationCmd = &cobra.Command{
	Use:   "request-allocation-update",
	Short: "Request allocation update for SP in the Collectif DAO protocol",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := requestUpdate(allocationLimit, dailyAllocation, run); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	RequestAllocationCmd.Flags().IntVarP(&allocationLimit, "limit", "l", 0, "Total allocation limit")
	RequestAllocationCmd.Flags().IntVarP(&dailyAllocation, "daily", "d", 0, "Daily allocation limit")
	RequestAllocationCmd.Flags().BoolVarP(&run, "run", "r", true, "Execute transaction")

	if err := RequestAllocationCmd.MarkFlagRequired("limit"); err != nil {
		fmt.Println(err)
	}
	if err := RequestAllocationCmd.MarkFlagRequired("daily"); err != nil {
		fmt.Println(err)
	}

	AllocationCmd.AddCommand(RequestAllocationCmd)
}
