package types

import (
	"github.com/shopspring/decimal"
	"sort"
)

type VestedResultKey struct {
	EmployeeId, AwardId, EmployeeName string
}

type VestedResults map[VestedResultKey]decimal.Decimal

func (c *VestedResults) GetSortedKeysByEmployeeIdAwardId() []VestedResultKey {
	var keys []VestedResultKey
	for key := range *c {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		if keys[i] == keys[j] {
			return keys[i].AwardId < keys[j].AwardId
		}
		return keys[i].EmployeeId < keys[j].EmployeeId
	})
	return keys
}
