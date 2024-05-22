package cast

import (
	"time"

	"github.com/google/uuid"
)

type numberType uint

const (
	numberTypeInteger = numberType(iota)
	numberTypeUnsignedInteger
	numberTypeFloat
)

type numbers interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}
type ptrNumbers interface {
	*int | *int8 | *int16 | *int32 | *int64 |
		*uint | *uint8 | *uint16 | *uint32 | *uint64 |
		*float32 | *float64
}

type supported interface {
	numbers | ptrNumbers |
		bool | string |
		*bool | *string |
		time.Time | *time.Time |
		uuid.UUID | *uuid.UUID
}
