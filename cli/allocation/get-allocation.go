package allocation

import (
	"collective-go-sdk/fvm"
	"collective-go-sdk/keystore"
	"collective-go-sdk/sdk"
	fUtils "collective-go-sdk/utils"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	minerId string
)

func getAllocationInfo(minerId string) (*fvm.SPAllocation, error) {
	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, keystore.FSKeyStore, "./")
	if err != nil {
		return nil, err
	}

	if minerId == "" {
		minerId = sdk.Client.Address.String()
	}

	idAddr := fUtils.GetIdAddress(ctx, minerId, sdk.Client)

	allocation, err := sdk.Client.GetAllocations(ctx, idAddr)
	if err != nil {
		return nil, err
	}

	return allocation, nil
}

var getAllocation = &cobra.Command{
	Use:   "get-allocation",
	Short: "Returns the allocation information for the Storage Provider",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if allocation, err := getAllocationInfo(minerId); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Allocation limit: ", allocation.AllocationLimit.String())
			fmt.Println("Repayment: ", allocation.Repayment.String())
			fmt.Println("Used allocation: ", allocation.UsedAllocation.String())
			fmt.Println("Daily allocation: ", allocation.DailyAllocation.String())
			fmt.Println("Accrued rewards: ", allocation.AccruedRewards.String())
			fmt.Println("Repaid pledge: ", allocation.RepaidPledge.String())
		}
	},
}

func init() {
	getAllocation.Flags().StringVarP(&minerId, "minerId", "o", "", "Storage Provider's miner address (or miner actor ID)")
	AllocationCmd.AddCommand(getAllocation)
}
