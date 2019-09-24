package eddystone

type mockUrlFrame struct {
	mockFrame
	mTxPower struct {
		calls uint
		rv1   int8
	}
	mUrl struct {
		calls uint
		rv1   string
	}
}

func (m *mockUrlFrame) TxPower() int8 {
	m.mTxPower.calls++
	return m.mTxPower.rv1
}

func (m *mockUrlFrame) Url() string {
	m.mUrl.calls++
	return m.mUrl.rv1
}

