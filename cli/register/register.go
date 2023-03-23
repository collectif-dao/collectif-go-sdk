package register

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"collective-go-sdk/utils"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var (
	miner  string
	pool   string
	limit  int64
	period int64
	run    bool
)

func registerStorageProvider(miner string, pool string, limit int64, period int64, run bool) (string, error) {
	bytesAddr := utils.ConvertAddress(miner)
	var poolAddr common.Address
	var allocationLimit *big.Int
	var maxPeriod *big.Int
	var attoFIL = big.NewInt(10)
	var fil = attoFIL.Exp(attoFIL, big.NewInt(18), big.NewInt(0))

	config, err := config.LoadConfig("./config")
	if err != nil {
		return "", err
	}

	if pool == "" {
		poolAddr = common.HexToAddress(config.LiquidStaking)
	} else {
		poolAddr = common.HexToAddress(pool)
	}

	if limit == 0 {
		allocationLimit = big.NewInt(config.AllocationLimit)
	} else {
		// allocationLimit = attoFIL.Exp(big.NewInt(limit), big.NewInt(18), big.NewInt(0))
		allocationLimit = fil.Mul(big.NewInt(limit), fil)
	}

	if period == 0 {
		maxPeriod = big.NewInt(config.MaxPeriod)
	} else {
		maxPeriod = big.NewInt(period)
	}

	ctx := context.Background()
	client, err := fvm.NewLotusClient(ctx, config, fvm.FSKeyStore)
	if err != nil {
		return "", err
	}

	abi, err := client.RegistryABI.GetAbi()
	if err != nil {
		return "", err
	}

	callData, err := abi.Pack("register", bytesAddr, poolAddr, allocationLimit, maxPeriod)

	return hex.EncodeToString(callData), nil

	// tx, err := client.Register(bytesAddr, poolAddr, allocationLimit, maxPeriod, run)
	// if err != nil {
	// 	return "", err
	// }

	// return tx.Hash().Hex(), nil
}

var RegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Register Storage Provider in the Collectif DAO protocol",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := registerStorageProvider(miner, pool, limit, period, run); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	RegisterCmd.Flags().StringVarP(&miner, "miner", "m", "", "Storage Provider miner address (filecoin address)")
	RegisterCmd.Flags().StringVarP(&pool, "staking", "s", "", "Liquid Staking pool address (ethereum address)")
	RegisterCmd.Flags().Int64VarP(&limit, "limit", "l", 0, "FIL allocation for pledge")
	RegisterCmd.Flags().Int64VarP(&limit, "period", "p", 0, "Max epoch for pledge operations")
	RegisterCmd.Flags().BoolVarP(&run, "run", "r", true, "Execute transaction")

	if err := RegisterCmd.MarkFlagRequired("miner"); err != nil {
		fmt.Println(err)
	}
}
