package cmd

import (
	"carta/internal/data_repo"
	"carta/internal/types"
	"carta/internal/vest_calcualtor"
	"fmt"
)

// calcualte the vesting schedule based on the input vesting event
type CmdCalcualteVestingSchedule struct {
	Places int32
}

func (c *CmdCalcualteVestingSchedule) Run(dataRepoClient data_repo.Client, vestCalculatorClient vest_calcualtor.Client) {
	dataChan := make(chan types.VestingEvent)

	go dataRepoClient.GetVestingEvents(dataChan)
	vestingResults := vestCalculatorClient.CalculateTotalVestedShares(dataChan)

	keys := vestingResults.GetSortedKeysByEmployeeIdAwardId()

	for _, key := range keys {
		result := vestingResults[key]
		line := fmt.Sprintf("%s,%s,%s,%v", key.EmployeeId, key.EmployeeName, key.AwardId, result.StringFixed(c.Places))
		fmt.Println(line)
	}
}
