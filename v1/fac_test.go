package eddystone

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFrameFactory_NewTlmFrame(t *testing.T) {
	Convey("Method: FrameFactory.NewTlmFrame", t, func() {
		Convey("Given any input", func() {
			Convey("It should call the backing TlmFrame factory function", func() {
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

			Convey("It should pass up the frame value from the Tlm factory", func() {
				ff := NewFrameFactory()
				mck := mockTlmFrame{}
				in := []byte{1, 2, 3}

				ff.TlmFrameFactory(func() TlmFrame { return &mck })
				out, _ := ff.NewTlmFrame(in)

				So(out, ShouldResemble, &mck)
			})

			Convey("It should call the backing frame FromBytes method", func() {
				ff := NewFrameFactory()
				mck := mockTlmFrame{}
				in := []byte{1, 2, 3}

				ff.TlmFrameFactory(func() TlmFrame { return &mck })
				_, _ = ff.NewTlmFrame(in)

				So(mck.MFromBytes.CallCount(), ShouldEqual, 1)
			})

			Convey("It should pass the byte slice to the Frame's FromBytes method", func() {
				ff := NewFrameFactory()
				mk := mockTlmFrame{}
				in := []byte{1, 2, 3}

				ff.TlmFrameFactory(func() TlmFrame {return &mk})
				_, _ = ff.NewTlmFrame(in)

				So(mk.MFromBytes.Inputs[0], ShouldResemble, in)
			})

			Convey("It should pass up the error value from the TlmFrame FromBytes method", func() {
				ff := NewFrameFactory()
				ex := errors.New("butts")
				mk := mockTlmFrame{}
				in := []byte{1, 2, 3}

				mk.MFromBytes.SetError(ex)
				ff.TlmFrameFactory(func() TlmFrame {return &mk})
				_, out := ff.NewTlmFrame(in)

				So(out, ShouldEqual, ex)
			})
		})
	})
}

func TestFrameFactory_NewUrlFrame(t *testing.T) {
	Convey("Method: FrameFactory.NewUrlFrame", t, func() {
		Convey("Given any input", func() {
			Convey("It should call the backing UrlFrame factory function", func() {
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

			Convey("It should pass up the frame value from the UrlFrame factory", func() {
				ff := NewFrameFactory()
				mck := mockUrlFrame{}
				in := []byte{1, 2, 3}

				ff.UrlFrameFactory(func() UrlFrame { return &mck })
				out, _ := ff.NewUrlFrame(in)

				So(out, ShouldResemble, &mck)
			})

			Convey("It should call the backing frame FromBytes method", func() {
				ff := NewFrameFactory()
				mck := mockUrlFrame{}
				in := []byte{1, 2, 3}

				ff.UrlFrameFactory(func() UrlFrame { return &mck })
				_, _ = ff.NewUrlFrame(in)

				So(mck.MFromBytes.CallCount(), ShouldEqual, 1)
			})

			Convey("It should pass the byte slice to the Frame's FromBytes method", func() {
				ff := NewFrameFactory()
				mk := mockUrlFrame{}
				in := []byte{1, 2, 3}

				ff.UrlFrameFactory(func() UrlFrame {return &mk})
				_, _ = ff.NewUrlFrame(in)

				So(mk.MFromBytes.Inputs[0], ShouldResemble, in)
			})

			Convey("It should pass up the error value from the UrlFrame FromBytes method", func() {
				ff := NewFrameFactory()
				ex := errors.New("butts")
				mk := mockUrlFrame{}
				in := []byte{1, 2, 3}

				mk.MFromBytes.SetError(ex)
				ff.UrlFrameFactory(func() UrlFrame {return &mk})
				_, out := ff.NewUrlFrame(in)

				So(out, ShouldEqual, ex)
			})
		})
	})
}

func TestFrameFactory_NewUidFrame(t *testing.T) {
	Convey("Method: FrameFactory.NewUidFrame", t, func() {
		Convey("Given any input", func() {
			Convey("It should call the backing UidFrame factory function", func() {
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

			Convey("It should pass up the frame value from the UidFrame factory", func() {
				ff := NewFrameFactory()
				mck := mockUidFrame{}
				in := []byte{1, 2, 3}

				ff.UidFrameFactory(func() UidFrame { return &mck })
				out, _ := ff.NewUidFrame(in)

				So(out, ShouldResemble, &mck)
			})

			Convey("It should call the backing frame FromBytes method", func() {
				ff := NewFrameFactory()
				mck := mockUidFrame{}
				in := []byte{1, 2, 3}

				ff.UidFrameFactory(func() UidFrame { return &mck })
				_, _ = ff.NewUidFrame(in)

				So(mck.MFromBytes.CallCount(), ShouldEqual, 1)
			})

			Convey("It should pass the byte slice to the Frame's FromBytes method", func() {
				ff := NewFrameFactory()
				mk := mockUidFrame{}
				in := []byte{1, 2, 3}

				ff.UidFrameFactory(func() UidFrame {return &mk})
				_, _ = ff.NewUidFrame(in)

				So(mk.MFromBytes.Inputs[0], ShouldResemble, in)
			})

			Convey("It should pass up the error value from the UidFrame FromBytes method", func() {
				ff := NewFrameFactory()
				ex := errors.New("butts")
				mk := mockUidFrame{}
				in := []byte{1, 2, 3}

				mk.MFromBytes.SetError(ex)
				ff.UidFrameFactory(func() UidFrame {return &mk})
				_, out := ff.NewUidFrame(in)

				So(out, ShouldEqual, ex)
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


