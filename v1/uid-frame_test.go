package eddystone

import (
	bean "github.com/Foxcapades/Beanies/v2"
	"github.com/google/uuid"
	"math/big"
)

type mockUidFrame struct {
	mockFrame

	MNamespaceBytes struct {
		Calls   uint
		Returns [10]byte
	}

	MInstanceBytes struct {
		Calls   uint
		Returns [6]byte
	}

	MNamespaceInt struct {
		Calls   uint
		Returns *big.Int
	}

	MUuid struct {
		Calls   uint
		Returns uuid.UUID
	}

	MTxPower         bean.Int8Getter
	MRangingData     bean.Int8Getter
	MNamespaceString bean.StringGetter
	MInstanceString  bean.StringGetter
	MInstanceInt     bean.Uint64Getter
}

func (m *mockUidFrame) TxPower() int8 {
	return m.MTxPower.Get()
}

func (m *mockUidFrame) RangingData() int8 {
	return m.MRangingData.Get()
}

func (m *mockUidFrame) NamespaceBytes() [10]byte {
	m.MNamespaceBytes.Calls++
	return m.MNamespaceBytes.Returns
}

func (m *mockUidFrame) NamespaceString() string {
	return m.MNamespaceString.Get()
}

func (m *mockUidFrame) NamespaceInt() *big.Int {
	m.MNamespaceInt.Calls++
	return m.MNamespaceInt.Returns
}

func (m *mockUidFrame) InstanceBytes() [6]byte {
	m.MInstanceBytes.Calls++
	return m.MInstanceBytes.Returns
}

func (m *mockUidFrame) InstanceString() string {
	return m.MInstanceString.Get()
}

func (m *mockUidFrame) InstanceInt() uint64 {
	return m.MInstanceInt.Get()
}

func (m *mockUidFrame) Uuid() uuid.UUID {
	m.MUuid.Calls++
	return m.MUuid.Returns
}
