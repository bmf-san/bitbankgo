package types

// TypeStatus is a type for status.
type TypeStatus string

const (
	// Normal is normal status.
	Normal = "NORMAL"
	// Busy is busy status.
	Busy = "BUSY"
	// VeryBusy is very busy status.
	VeryBusy = "VERY_BUSY"
	// Halt is halt status.
	Halt = "HALT"
)
