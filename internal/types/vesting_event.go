package types

import (
	"github.com/shopspring/decimal"
	"time"
)

type VestingEvent struct {
	VestingAction VestingActionType
	EmployeeId    string
	EmployeeName  string
	AwardId       string
	Date          time.Time
	Quantity      decimal.Decimal
}
