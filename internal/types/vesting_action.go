package types

type VestingActionType string

const (
	VestAction   VestingActionType = "VEST"
	CancelAction VestingActionType = "CANCEL"
)
