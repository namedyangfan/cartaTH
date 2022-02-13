package vest_calcualtor

import (
	"carta/internal/types"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"time"
)

type Client interface {
	// CalculateTotalVestedShares calculate the vesting schedule based on the vesting event
	CalculateTotalVestedShares(dataChan chan types.VestingEvent) types.VestedResults
}

type client struct {
	targetDate time.Time
}

func NewClient(targetDate time.Time) Client {
	return &client{targetDate: targetDate}
}

func (c *client) CalculateTotalVestedShares(dataChan chan types.VestingEvent) types.VestedResults {
	vestedResults := make(types.VestedResults)

	for event := range dataChan {
		key := types.VestedResultKey{
			EmployeeId:   event.EmployeeId,
			AwardId:      event.AwardId,
			EmployeeName: event.EmployeeName,
		}

		if event.Date.Equal(c.targetDate) || event.Date.Before(c.targetDate) {
			switch event.VestingAction {
			case types.VestAction:
				vestedResults[key] = vestedResults[key].Add(event.Quantity)
			case types.CancelAction:
				vestedResults[key] = vestedResults[key].Sub(event.Quantity)
			default:
				zap.S().Panic("unrecognized action:", event.VestingAction)
			}
		} else {
			if _, ok := vestedResults[key]; !ok {
				vestedResults[key] = decimal.Decimal{}
			}
		}
	}
	return vestedResults
}
