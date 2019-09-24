package eddystone

import (
	"github.com/google/uuid"
	"math/big"
)

type mockUidFrame struct {
	mockFrame
	mTxPower struct {
		calls uint
		rv1   int8
	}
	mRangingData struct {
		calls uint
		rv1   int8
	}
	mNamespaceBytes struct {
		calls uint
		rv1   [10]byte
	}
	mNamespaceString struct {
		calls uint
		rv1   string
	}
	mNamespaceInt struct {
		calls uint
		rv1   *big.Int
	}
	mInstanceBytes struct {
		calls uint
		rv1   [6]byte
	}
	mInstanceString struct {
		calls uint
		rv1   string
	}
	mInstanceInt struct {
		calls uint
		rv1   uint64
	}
	mUuid struct {
		calls uint
		rv1   uuid.UUID
	}
}

func (m *mockUidFrame) TxPower() int8 {
	m.mTxPower.calls++
	return m.mTxPower.rv1
}

func (m *mockUidFrame) RangingData() int8 {
	m.mRangingData.calls++
	return m.mRangingData.rv1
}

func (m *mockUidFrame) NamespaceBytes() [10]byte {
	m.mNamespaceBytes.calls++
	return m.mNamespaceBytes.rv1
}

func (m *mockUidFrame) NamespaceString() string {
	m.mNamespaceString.calls++
	return m.mNamespaceString.rv1
}

func (m *mockUidFrame) NamespaceInt() *big.Int {
	m.mNamespaceInt.calls++
	return m.mNamespaceInt.rv1
}

func (m *mockUidFrame) InstanceBytes() [6]byte {
	m.mInstanceBytes.calls++
	return m.mInstanceBytes.rv1
}

func (m *mockUidFrame) InstanceString() string {
	m.mInstanceString.calls++
	return m.mInstanceString.rv1
}

func (m *mockUidFrame) InstanceInt() uint64 {
	m.mInstanceInt.calls++
	return m.mInstanceInt.rv1
}

func (m *mockUidFrame) Uuid() uuid.UUID {
	m.mUuid.calls++
	return m.mUuid.rv1
}

