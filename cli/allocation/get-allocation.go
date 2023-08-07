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

func getAllocationInfo(minerAddr string) (*fvm.SPAllocation, error) {
	ctx := context.Background()
	sdk, err := sdk.NewCollectifSDK(ctx, keystore.FSKeyStore, "./")
	if err != nil {
		return nil, err
	}

	idAddr := fUtils.GetIdAddress(ctx, minerAddr, sdk.Client)

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

		if allocation, err := getAllocationInfo(minerAddr); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Allocation limit: ", fvm.AttoFIL2FIL_str(allocation.AllocationLimit), " FIL")
			fmt.Println("Repayment: ", fvm.AttoFIL2FIL_str(allocation.Repayment), " FIL")
			fmt.Println("Used allocation: ", fvm.AttoFIL2FIL_str(allocation.UsedAllocation), " FIL")
			fmt.Println("Daily allocation: ", fvm.AttoFIL2FIL_str(allocation.DailyAllocation), " FIL")
			fmt.Println("Accrued rewards: ", fvm.AttoFIL2FIL_str(allocation.AccruedRewards), " FIL")
			fmt.Println("Repaid pledge: ", fvm.AttoFIL2FIL_str(allocation.RepaidPledge), " FIL")
		}
	},
}

func init() {
	getAllocation.Flags().StringVarP(&minerAddr, "minerAddr", "m", "", "Miner actor address (either ID address or actor address)")
	if err := getAllocation.MarkFlagRequired("minerAddr"); err != nil {
		fmt.Println(err)
	}

	AllocationCmd.AddCommand(getAllocation)
}
