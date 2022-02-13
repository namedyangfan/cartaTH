package main

import (
	"carta/cmd"
	"carta/internal/data_repo"
	"carta/internal/vest_calcualtor"
	"flag"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func main() {
	flag.Parse()
	if flag.NArg() != 2 && flag.NArg() != 3 {
		zap.S().Panicf("at least two arguments were required but %v were given", flag.NArg())
	}

	fileName := flag.Arg(0)

	targetDate, err := time.Parse("2006-01-02", flag.Arg(1))
	if err != nil {
		zap.S().Panicf("fail to parse the target date %v", err)
	}

	inputPlaces := "2"
	if flag.Arg(2) != "" {
		inputPlaces = flag.Arg(2)
	}

	places, err := strconv.ParseInt(inputPlaces, 10, 32)
	if err != nil || places < 0 || places > 6 {
		zap.S().Panicf("precision need to be between 0 and 6")
	}

	dataRepoClient := data_repo.NewClient(fileName, int32(places))
	vestCalculatorClient := vest_calcualtor.NewClient(targetDate)

	cmdCalculateVestingSchedule := cmd.CmdCalcualteVestingSchedule{Places: int32(places)}
	cmdCalculateVestingSchedule.Run(dataRepoClient, vestCalculatorClient)
}
