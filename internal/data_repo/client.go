package data_repo

import (
	"carta/internal/types"
	"encoding/csv"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"io"
	"os"
	"time"
)

type Client interface {
	// GetVestingEvents reads VestingEvents and put it into the dataChan channel
	GetVestingEvents(dataChan chan types.VestingEvent)
}

type client struct {
	fileName string
	places   int32
}

func NewClient(fileName string, places int32) Client {
	return &client{fileName: fileName, places: places}
}

func (c *client) GetVestingEvents(dataChan chan types.VestingEvent) {
	defer close(dataChan)

	csvFile, err := os.Open(c.fileName)
	if err != nil {
		zap.S().Panic(err)
	}

	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			zap.S().Panic(err)
		}
	}(csvFile)

	reader := csv.NewReader(csvFile)

	for {
		// Read each record from csv
		record, err := reader.Read()
		if err == io.EOF {
			zap.S().Info("end of line")
			return
		}
		if err != nil {
			zap.S().Errorf("fail to read line with error: %v", err)
		}

		date, err := time.Parse("2006-01-02", record[4])
		if err != nil {
			zap.S().Errorf("failed to parse the date with error: %v", err)
		}

		quantity, err := decimal.NewFromString(record[5])
		if err != nil {
			zap.S().Errorf("failed to parse the quantity with error: %v", err)
		}

		event := types.VestingEvent{
			VestingAction: types.VestingActionType(record[0]),
			EmployeeId:    record[1],
			EmployeeName:  record[2],
			AwardId:       record[3],
			Date:          date,
			Quantity:      quantity.RoundDown(c.places),
		}
		dataChan <- event
	}
}
