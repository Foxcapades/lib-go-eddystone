package eddystone

import (
	"github.com/Foxcapades/Beanies/v2"
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏   Unit Tests                                           ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/

func TestUrlFrame_FromBytes(t *testing.T) {
	Convey("Method: UrlFrame.FromBytes", t, func() {
		// TODO
	})
}

func TestUrlFrame_MarshalJSON(t *testing.T) {
	Convey("Method: UrlFrame.MarshalJSON", t, func() {
		// TODO
	})
}

func TestUrlFrame_String(t *testing.T) {
	Convey("Method: UrlFrame.String", t, func() {
		// TODO
	})
}

func TestUrlFrame_ToBytes(t *testing.T) {
	Convey("Method: UrlFrame.ToBytes", t, func() {
		// TODO
	})
}

func TestUrlFrame_TxPower(t *testing.T) {
	Convey("Method: UrlFrame.TxPower", t, func() {
		// TODO
	})
}

func TestUrlFrame_Type(t *testing.T) {
	Convey("Method: UrlFrame.Type", t, func() {
		// TODO
	})
}

func TestUrlFrame_Url(t *testing.T) {
	Convey("Method: UrlFrame.Url", t, func() {
		// TODO
	})
}

/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏  Mock Types                                            ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/

type mockUrlFrame struct {
	mockFrame

	MTxPower bean.Int8Getter
	MUrl     bean.StringGetter
}

func (m *mockUrlFrame) TxPower() int8 {
	return m.MTxPower.Get()
}

func (m *mockUrlFrame) Url() string {
	return m.MUrl.Get()
}
