package staking

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"context"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	sector uint64
	proof  string
	run    bool
)

func pledge(sector uint64, proof string, run bool) (string, error) {
	proofBytes := []byte(proof)
	config, err := config.LoadConfig("./config")
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	client, err := fvm.NewLotusClient(ctx, config, fvm.FSKeyStore)
	if err != nil {
		return "", err
	}

	abi, err := client.StakingABI.GetAbi()
	if err != nil {
		return "", err
	}

	callData, err := abi.Pack("pledge", sector, proofBytes)

	return hex.EncodeToString(callData), nil

	// tx, err := client.Pledge(sector, proofBytes, run)

	// if err != nil {
	// 	return "", err
	// }

	// return tx.Hash().Hex(), nil
}

var pledgeCmd = &cobra.Command{
	Use:   "pledge",
	Short: "Pledge sector to finalize sealing and increase initial pledge for Storage Provider",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := pledge(sector, proof, run); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	pledgeCmd.Flags().Uint64VarP(&sector, "sector", "s", 0, "Sector number")
	pledgeCmd.Flags().StringVarP(&proof, "proof", "p", "", "Sector proof")
	pledgeCmd.Flags().BoolVarP(&run, "run", "r", true, "Execute transaction")

	if err := pledgeCmd.MarkFlagRequired("sector"); err != nil {
		fmt.Println(err)
	}

	StakingCmd.AddCommand(pledgeCmd)
}
