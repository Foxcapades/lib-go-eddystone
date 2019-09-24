package mock

import (
	"github.com/google/uuid"
	"math/big"
)

type BigIntGetter struct {
	Calls   uint
	Returns *big.Int
}

type ByteSliceGetter struct {
	Calls   uint
	Returns []byte
}

type Int8Getter struct {
	Calls   uint
	Returns int8
}

type StringGetter struct {
	Calls   uint
	Returns string
}

type Uint8Getter struct {
	Calls   uint
	Returns uint8
}

type Uint64Getter struct {
	Calls   uint
	Returns uint64
}

type UuidGetter struct {
	Calls   uint
	Returns uuid.UUID
}
