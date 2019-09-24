package mock

type UrlFrame struct {
	Frame

	MTxPower Int8Getter
	MUrl     StringGetter
}

func (m *UrlFrame) TxPower() int8 {
	m.MTxPower.Calls++
	return m.MTxPower.Returns
}

func (m *UrlFrame) Url() string {
	m.MUrl.Calls++
	return m.MUrl.Returns
}
