package mock

import (
	"github.com/google/uuid"
	"math/big"
)

type UidFrame struct {
	Frame

	MNamespaceBytes struct {
		Calls   uint
		Returns [10]byte
	}

	MInstanceBytes struct {
		Calls   uint
		Returns [6]byte
	}

	MTxPower         Int8Getter
	MRangingData     Int8Getter
	MNamespaceString StringGetter
	MNamespaceInt    BigIntGetter
	MInstanceString  StringGetter
	MInstanceInt     Uint64Getter
	MUuid            UuidGetter
}

func (m *UidFrame) TxPower() int8 {
	m.MTxPower.Calls++
	return m.MTxPower.Returns
}

func (m *UidFrame) RangingData() int8 {
	m.MRangingData.Calls++
	return m.MRangingData.Returns
}

func (m *UidFrame) NamespaceBytes() [10]byte {
	m.MNamespaceBytes.Calls++
	return m.MNamespaceBytes.Returns
}

func (m *UidFrame) NamespaceString() string {
	m.MNamespaceString.Calls++
	return m.MNamespaceString.Returns
}

func (m *UidFrame) NamespaceInt() *big.Int {
	m.MNamespaceInt.Calls++
	return m.MNamespaceInt.Returns
}

func (m *UidFrame) InstanceBytes() [6]byte {
	m.MInstanceBytes.Calls++
	return m.MInstanceBytes.Returns
}

func (m *UidFrame) InstanceString() string {
	m.MInstanceString.Calls++
	return m.MInstanceString.Returns
}

func (m *UidFrame) InstanceInt() uint64 {
	m.MInstanceInt.Calls++
	return m.MInstanceInt.Returns
}

func (m *UidFrame) Uuid() uuid.UUID {
	m.MUuid.Calls++
	return m.MUuid.Returns
}
