package worker

import (
	"collective-go-sdk/fvm"
	"collective-go-sdk/utils"
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/filecoin-project/go-address"
	log "github.com/sirupsen/logrus"
)

type Miner struct {
	Address                *address.Address
	MinerId                uint64
	OwnerId                uint64
	DailyAllocation        *big.Int
	TotalAllocation        *big.Int
	Allocated              *big.Int
	AvailableCollateral    *big.Int
	LockedCollateral       *big.Int
	CollateralRequirements *big.Int
	RepaidPledge           *big.Int
	Debt                   *big.Int
}

type Worker struct {
	Ticker              *time.Ticker
	Client              *fvm.LotusClient
	Miners              []Miner
	ExecuteTransactions bool
}

func NewWorker(t time.Duration, c *fvm.LotusClient, m []Miner, executeTxs bool) *Worker {
	return &Worker{
		Ticker:              time.NewTicker(t),
		Client:              c,
		Miners:              m,
		ExecuteTransactions: executeTxs,
	}
}

func (w *Worker) Start(ctx context.Context) {
	log.Info("Collectif Client started")

	for {
		select {
		case <-w.Ticker.C:
			for _, m := range w.Miners {
				if m.Address == nil {
					log.Warn("Miner address is nil ")
					return
				}

				log.Info("Working with miner ", m.Address.String())

				msg, err := w.tryPledge(ctx, &m)
				if err != nil {
					log.Error("Unable to pledge", err)
				}

				log.Info("Successfull tranasction with ", msg.Message, " CID")
			}
		}
	}
}

func (w *Worker) tryPledge(ctx context.Context, m *Miner) (*fvm.MessageResponse, error) {
	availableFIL, err := w.Client.TotalFilAvailable(ctx)
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	if m.DailyAllocation.Cmp(availableFIL) == 1 {
		err := errors.New("Not enough FIL in the Liquid Staking contract")
		log.Warn(err)

		return nil, err
	}

	var newAllocation big.Int
	newAllocation.Add(m.Allocated, m.DailyAllocation)

	if newAllocation.Cmp(m.TotalAllocation) == 1 {
		err := errors.New("Total allocation overflow")
		log.Warn(err)

		return nil, err
	}

	msg, err := w.increaseCollateral(ctx, m, &newAllocation)
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	msg, err = w.Client.Pledge(ctx, m.DailyAllocation, m.MinerId, w.ExecuteTransactions)
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	return msg, nil
}

func (w *Worker) increaseCollateral(ctx context.Context, m *Miner, allocated *big.Int) (*fvm.MessageResponse, error) {
	var requirements = big.NewInt(0)
	var totalReq big.Int
	var amount big.Int

	if allocated.Cmp(m.RepaidPledge) == 1 {
		requirements.Sub(allocated, m.RepaidPledge)
	}

	requirements.Mul(requirements, m.CollateralRequirements)
	requirements.Div(requirements, big.NewInt(10000))

	totalReq.Add(requirements, m.Debt)

	var totalCollateral big.Int
	totalCollateral.Add(m.AvailableCollateral, m.LockedCollateral)

	if totalCollateral.Cmp(&totalReq) >= 0 {
		return nil, nil
	}

	amount.Sub(&totalReq, &totalCollateral)

	log.Info("Collateral deposit is required before pledging for  ", amount.String(), " amount")

	msg, err := w.Client.Deposit(ctx, &amount, w.ExecuteTransactions)
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	return msg, nil
}

func PrepareMiners(ctx context.Context, minerList []string, client *fvm.LotusClient) ([]Miner, error) {
	miners := make([]Miner, 0)

	for i := 0; i < len(minerList); i++ {
		idAddr := utils.GetIdAddress(ctx, minerList[i], client)

		allocations, err := client.GetAllocations(ctx, idAddr)
		if err != nil {
			return nil, err
		}

		ownerId, err := client.GetMinerOwner(ctx, idAddr)
		if err != nil {
			return nil, err
		}

		collateral, err := client.GetCollateral(ctx, ownerId)
		if err != nil {
			return nil, err
		}

		collateralRequirements, err := client.GetCollateralRequirements(ctx, ownerId)
		if err != nil {
			return nil, err
		}

		debt, err := client.GetDebt(ctx, ownerId)
		if err != nil {
			return nil, err
		}

		mAddr, err := utils.LookupIdAddress(ctx, minerList[i], client)
		if err != nil {
			return nil, err
		}

		miner := Miner{
			Address:                mAddr,
			MinerId:                idAddr,
			OwnerId:                ownerId,
			DailyAllocation:        allocations.DailyAllocation,
			Allocated:              allocations.UsedAllocation,
			TotalAllocation:        allocations.AllocationLimit,
			RepaidPledge:           allocations.RepaidPledge,
			AvailableCollateral:    collateral.AvailableCollateral,
			LockedCollateral:       collateral.LockedCollateral,
			CollateralRequirements: collateralRequirements,
			Debt:                   debt,
		}

		miners = append(miners, miner)
	}

	return miners, nil
}
