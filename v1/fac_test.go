package eddystone

import (
	"github.com/Foxcapades/lib-go-eddystone/v1/mock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFrameFactory_NewTlmFrame(t *testing.T) {
	Convey("Method: FrameFactory.NewTlmFrame", t, func() {
		Convey("Given any input", func() {
			Convey("It should call the backing TLM factory function", func() {
				ff := NewFrameFactory()
				res := false
				ff.TlmFrameFactory(func() TlmFrame {
					res = true
					return nil
				})
				_, _ = ff.NewTlmFrame(nil)
				So(res, ShouldBeTrue)
			})
		})

		Convey("Given a nil input", func() {
			Convey("It should return a nil bytes error", func() {
				_, e := NewFrameFactory().NewTlmFrame(nil)
				So(e, ShouldNotBeNil)
				So(e.Error(), ShouldEqual, errNilInput)
			})
		})

		Convey("Given a non nil byte array", func() {
			Convey("It should call the backing frame FromBytes method", func() {
				ff := NewFrameFactory()
				mock := mockTlmFrame{}
				in := []byte{1, 2, 3}

				ff.TlmFrameFactory(func() TlmFrame { return &mock })
				_, _ = ff.NewTlmFrame(in)

				So(mock.mFromBytes.calls, ShouldEqual, 1)
			})

			Convey("It should pass up the frame value from the Tlm factory", func() {
				// TODO
			})

			Convey("It should pass up the error value from the Tlm factory", func() {
				// TODO
			})
		})
	})
}

func TestFrameFactory_NewUrlFrame(t *testing.T) {
	Convey("Method: FrameFactory.NewUrlFrame", t, func() {
		Convey("Given any input", func() {
			Convey("It should call the backing TLM factory function", func() {
				ff := NewFrameFactory()
				res := false
				ff.UrlFrameFactory(func() UrlFrame {
					res = true
					return nil
				})
				_, _ = ff.NewUrlFrame(nil)
				So(res, ShouldBeTrue)
			})
		})

		Convey("Given a nil input", func() {
			Convey("It should return a nil bytes error", func() {
				_, e := NewFrameFactory().NewUrlFrame(nil)
				So(e, ShouldNotBeNil)
				So(e.Error(), ShouldEqual, errNilInput)
			})
		})

		Convey("Given a non nil byte array", func() {
			Convey("It should call the backing frame FromBytes method", func() {
				ff := NewFrameFactory()
				mock := mockUrlFrame{}
				in := []byte{1, 2, 3}

				ff.UrlFrameFactory(func() UrlFrame { return &mock })
				_, _ = ff.NewUrlFrame(in)

				So(mock.mFromBytes.calls, ShouldEqual, 1)
			})

			Convey("It should pass up the frame value from the Url factory", func() {
				// TODO
			})

			Convey("It should pass up the error value from the Url factory", func() {
				// TODO
			})
		})
	})
}

func TestFrameFactory_NewUidFrame(t *testing.T) {
	Convey("Method: FrameFactory.NewUidFrame", t, func() {
		Convey("Given any input", func() {
			Convey("It should call the backing TLM factory function", func() {
				ff := NewFrameFactory()
				res := false
				ff.UidFrameFactory(func() UidFrame {
					res = true
					return nil
				})
				_, _ = ff.NewUidFrame(nil)
				So(res, ShouldBeTrue)
			})
		})

		Convey("Given a nil input", func() {
			Convey("It should return a nil bytes error", func() {
				_, e := NewFrameFactory().NewUidFrame(nil)
				So(e, ShouldNotBeNil)
				So(e.Error(), ShouldEqual, errNilInput)
			})
		})

		Convey("Given a non nil byte array", func() {
			Convey("It should call the backing frame FromBytes method", func() {
				ff := NewFrameFactory()
				mck := mock.UidFrame{}
				in := []byte{1, 2, 3}

				ff.UidFrameFactory(func() UidFrame { return &mck })
				_, _ = ff.NewUidFrame(in)

				So(mck.MFromBytes.calls, ShouldEqual, 1)
			})

			Convey("It should pass up the frame value from the Uid factory", func() {
				// TODO
			})

			Convey("It should pass up the error value from the Uid factory", func() {
				// TODO
			})
		})
	})
}

func TestFrameFactory_TlmFrameFactory(t *testing.T) {
	Convey("Method: FrameFactory.TlmFrameFactory", t, func() {})
}

func TestFrameFactory_UrlFrameFactory(t *testing.T) {
	Convey("Method: FrameFactory.UrlFrameFactory", t, func() {})
}

func TestFrameFactory_UuidFrameFactory(t *testing.T) {
	Convey("Method: FrameFactory.UidFrameFactory", t, func() {})
}


