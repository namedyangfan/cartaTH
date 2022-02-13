package vest_calcualtor

import (
	"carta/internal/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type VestCalculatorTestSuite struct {
	suite.Suite
	targetDate time.Time
}

func (s *VestCalculatorTestSuite) SetupTest() {
	s.targetDate, _ = time.Parse("2006-01-02", "2020-04-01")
}

func TestVestCalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(VestCalculatorTestSuite))
}

func (s *VestCalculatorTestSuite) TestFilterTargetDate() {
	// Setup
	dataChan := make(chan types.VestingEvent)
	go func(dataChan chan types.VestingEvent) {
		defer close(dataChan)
		dataChan <- types.VestingEvent{
			VestingAction: types.VestAction,
			EmployeeId:    "E001",
			EmployeeName:  "Alice Smith",
			AwardId:       "ISO-001",
			Date:          s.targetDate.AddDate(0, -1, 0),
			Quantity:      decimal.RequireFromString("1000"),
		}
		dataChan <- types.VestingEvent{
			VestingAction: types.VestAction,
			EmployeeId:    "E001",
			EmployeeName:  "Alice Smith",
			AwardId:       "ISO-001",
			Date:          s.targetDate.AddDate(1, 1, 0),
			Quantity:      decimal.RequireFromString("1000"),
		}
		dataChan <- types.VestingEvent{
			VestingAction: types.VestAction,
			EmployeeId:    "E001",
			EmployeeName:  "Alice Smith",
			AwardId:       "ISO-002",
			Date:          s.targetDate.AddDate(-1, 0, 0),
			Quantity:      decimal.RequireFromString("300"),
		}
		dataChan <- types.VestingEvent{
			VestingAction: types.VestAction,
			EmployeeId:    "E001",
			EmployeeName:  "Alice Smith",
			AwardId:       "ISO-002",
			Date:          s.targetDate.AddDate(-1, 0, 0),
			Quantity:      decimal.RequireFromString("800"),
		}
		dataChan <- types.VestingEvent{
			VestingAction: types.VestAction,
			EmployeeId:    "E003",
			EmployeeName:  "Cat Helms",
			AwardId:       "NSO-002",
			Date:          s.targetDate.AddDate(3, 0, 0),
			Quantity:      decimal.RequireFromString("100"),
		}
	}(dataChan)
	// Execution
	client := NewClient(s.targetDate)
	results := client.CalculateTotalVestedShares(dataChan)
	s.Assert().EqualValues(results[types.VestedResultKey{EmployeeId: "E001", AwardId: "ISO-001", EmployeeName: "Alice Smith"}].String(), "1000")
	s.Assert().EqualValues(results[types.VestedResultKey{EmployeeId: "E001", AwardId: "ISO-002", EmployeeName: "Alice Smith"}].String(), "1100")
	s.Assert().EqualValues(results[types.VestedResultKey{EmployeeId: "E003", AwardId: "ISO-002", EmployeeName: "Cat Helms"}].String(), "0")
}

func (s *VestCalculatorTestSuite) TestCancelAction() {
	// Setup
	dataChan := make(chan types.VestingEvent)
	go func(dataChan chan types.VestingEvent) {
		defer close(dataChan)
		dataChan <- types.VestingEvent{
			VestingAction: types.VestAction,
			EmployeeId:    "E001",
			EmployeeName:  "Alice Smith",
			AwardId:       "ISO-001",
			Date:          s.targetDate.AddDate(0, -1, 0),
			Quantity:      decimal.RequireFromString("1000"),
		}
		dataChan <- types.VestingEvent{
			VestingAction: types.CancelAction,
			EmployeeId:    "E001",
			EmployeeName:  "Alice Smith",
			AwardId:       "ISO-001",
			Date:          s.targetDate.AddDate(0, -1, 0),
			Quantity:      decimal.RequireFromString("500"),
		}
		dataChan <- types.VestingEvent{
			VestingAction: types.CancelAction,
			EmployeeId:    "E001",
			EmployeeName:  "Alice Smith",
			AwardId:       "ISO-001",
			Date:          s.targetDate.AddDate(0, -1, 0),
			Quantity:      decimal.RequireFromString("300"),
		}
		dataChan <- types.VestingEvent{
			VestingAction: types.VestAction,
			EmployeeId:    "E001",
			EmployeeName:  "Alice Smith",
			AwardId:       "ISO-001",
			Date:          s.targetDate.AddDate(0, -1, 0),
			Quantity:      decimal.RequireFromString("100"),
		}
	}(dataChan)
	// Execution
	client := NewClient(s.targetDate)
	results := client.CalculateTotalVestedShares(dataChan)
	s.Assert().EqualValues(results[types.VestedResultKey{EmployeeId: "E001", AwardId: "ISO-001", EmployeeName: "Alice Smith"}].String(), "300")
}
