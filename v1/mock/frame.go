package mock

import (
	e "github.com/Foxcapades/lib-go-eddystone/v1"
)

type Frame struct {
	MString StringGetter

	MMarshalJSON struct {
		calls uint
		rv1   []byte
		rv2   error
	}

	MType struct {
		calls uint
		rv1   e.FrameType
	}

	MToBytes ByteSliceGetter

	MFromBytes struct {
		calls uint
		in1   [][]byte
		rv1   error
	}
}

func (m *Frame) String() string {
	m.MString.calls++
	return m.MString.rv1
}

func (m *Frame) MarshalJSON() ([]byte, error) {
	m.MMarshalJSON.calls++
	return m.MMarshalJSON.rv1, m.MMarshalJSON.rv2
}

func (m *Frame) Type() e.FrameType {
	m.MType.calls++
	return m.MType.rv1
}

func (m *Frame) ToBytes() []byte {
	m.MToBytes.calls++
	return m.MToBytes.rv1
}

func (m *Frame) FromBytes(b []byte) error {
	m.MFromBytes.calls++
	m.MFromBytes.in1 = append(m.MFromBytes.in1, b)
	return m.MFromBytes.rv1
}
