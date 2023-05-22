package allocation

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	fUtils "collective-go-sdk/fvm/utils"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ownerId string
)

func getAllocationInfo(ownerId string) (*fvm.SPAllocation, error) {
	config, err := config.LoadConfig("./config")
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	client, err := fvm.NewLotusClient(ctx, config, fvm.FSKeyStore)
	if err != nil {
		return nil, err
	}

	if ownerId == "" {
		ownerId = client.Address.String()
	}

	idAddr := fUtils.GetIdAddress(ctx, ownerId, client)

	allocation, err := client.GetAllocations(idAddr)
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

		if allocation, err := getAllocationInfo(ownerId); err != nil {
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
	getAllocation.Flags().StringVarP(&ownerId, "ownerId", "o", "", "Storage Provider ownerID (filecoin)")
	AllocationCmd.AddCommand(getAllocation)
}
