package eddystone

type mockFrame struct {
	mString struct {
		calls uint
		rv1   string
	}

	mMarshalJSON struct {
		calls uint
		rv1   []byte
		rv2   error
	}

	mType struct {
		calls uint
		rv1   FrameType
	}

	mToBytes struct {
		calls uint
		rv1   []byte
	}

	mFromBytes struct {
		calls uint
		in1   [][]byte
		rv1   error
	}
}

func (m *mockFrame) String() string {
	m.mString.calls++
	return m.mString.rv1
}

func (m *mockFrame) MarshalJSON() ([]byte, error) {
	m.mMarshalJSON.calls++
	return m.mMarshalJSON.rv1, m.mMarshalJSON.rv2
}

func (m *mockFrame) Type() FrameType {
	m.mType.calls++
	return m.mType.rv1
}

func (m *mockFrame) ToBytes() []byte {
	m.mToBytes.calls++
	return m.mToBytes.rv1
}

func (m *mockFrame) FromBytes(b []byte) error {
	m.mFromBytes.calls++
	m.mFromBytes.in1 = append(m.mFromBytes.in1, b)
	return m.mFromBytes.rv1
}

