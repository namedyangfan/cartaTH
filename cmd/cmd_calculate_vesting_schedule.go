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

func (c *CmdCalcualteVestingSchedule) PrintSortedVestingResults(vestingResults types.VestedResults) {
	keys := vestingResults.GetSortedKeysByEmployeeIdAwardId()

	for _, key := range keys {
		result := vestingResults[key]
		line := fmt.Sprintf("%s,%s,%s,%v", key.EmployeeId, key.EmployeeName, key.AwardId, result.StringFixed(c.Places))
		fmt.Println(line)
	}
}

func (c *CmdCalcualteVestingSchedule) Run(dataRepoClient data_repo.Client, vestCalculatorClient vest_calcualtor.Client) {
	dataChan := make(chan types.VestingEvent)

	// set up data channel to get vesting event
	go dataRepoClient.GetVestingEvents(dataChan)

	// process the vesting event to get vesting result
	vestingResults := vestCalculatorClient.CalculateTotalVestedShares(dataChan)

	// print out vesting result
	c.PrintSortedVestingResults(vestingResults)
}
