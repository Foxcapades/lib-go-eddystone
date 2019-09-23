package eddystone

import (
	"fmt"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"strings"
	"testing"
)

func Test_digitToHex(t *testing.T) {
	Convey("Function: digitToHex", t, func() {

		Convey("Given an input byte in the set [0..15]", func() {
			Convey("The output digit should align with fmt.Sprintf", func() {
				for i := byte(0); i < 16; i++ {
					So(fmt.Sprintf("%x", i)[0], ShouldEqual, digitToHex(i))
				}
			})
		})

		Convey("Given an input byte in the set [16..255]", func() {
			Convey("The function should panic", func() {
				for i := 16; i < 256; i++ {
					So(func() {digitToHex(byte(i))}, ShouldPanic)
				}
			})
		})

	})
}

func BenchmarkHexString(b *testing.B) {
	tests := []struct{
		name string
		fn func(byte, *strings.Builder)
	} {
		{"eddystone.byteToHex",
			func(i byte, s *strings.Builder) {byteToHex(i, s)}},

		{"fmt.Fprintf",
			func(i byte, s *strings.Builder) {_, _ = fmt.Fprintf(s, "%02x", i)}},
	}

	for _, val := range tests {
		b.Run(val.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var act strings.Builder
				for i := 0; i < math.MaxUint8; i++ {
					val.fn(byte(i), &act)
				}
			}
		})
	}
}

func Test_byteToHex(t *testing.T) {
	Convey("Function: byteToHex", t, func() {
		Convey("Given any byte value as a first param", func() {

			Convey("When the second param is nil", func() {
				Convey("It should panic", func() {
					So(func() {byteToHex(0, nil)}, ShouldPanic)
				})
			})

			Convey("When the second param is a string builder", func() {
				Convey("It should build the same string as fmt.Sprintf(\"%02x\")", func() {
					var exp strings.Builder
					var act strings.Builder

					for i := 0; i < math.MaxUint8; i++ {
						_, _ = fmt.Fprintf(&exp, "%02x", i)
						byteToHex(byte(i), &act)
					}

					So(exp.String(), ShouldEqual, act.String())
				})
			})
		})
	})
}

func Test_u32ToByte(t *testing.T) {
	Convey("Function: u32ToByte", t, func() {
		Convey("Given any uint 32 input", func() {
			Convey("When that value is within the set [0..255]", func() {
				Convey("It should return the same int value", func() {
					for i := uint32(0); i < 256; i++ {
						So(byte(i), ShouldEqual, u32ToByte(i))
					}
				})
			})

			Convey("When that value is greater than 255", func() {
				Convey("It should return the least significant byte", func() {
					tests := []struct{
						input  uint32
						expect byte
					} {
						{0xffffff00, 0},
						{0xff0000ff, 255},
						{0x1b25b2df, 223},
					}

					for _, test := range tests {
						So(test.expect, ShouldEqual, u32ToByte(test.input))
					}
				})
			})
		})
	})
}

func Test_u16ToByte(t *testing.T) {
	Convey("Function: u16ToByte", t, func() {
		Convey("Given any uint 16 input", func() {
			Convey("When that value is within the set [0..255]", func() {
				Convey("It should return the same int value", func() {
					for i := uint16(0); i < 256; i++ {
						So(byte(i), ShouldEqual, u16ToByte(i))
					}
				})
			})

			Convey("When that value is greater than 255", func() {
				Convey("It should return the least significant byte", func() {
					tests := []struct{
						input  uint16
						expect byte
					} {
						{0xff00, 0},
						{0xffff, 255},
						{0xb2df, 223},
					}

					for _, test := range tests {
						So(test.expect, ShouldEqual, u16ToByte(test.input))
					}
				})
			})
		})
	})
}

func Test_cToF(t *testing.T) {
	Convey("Function: cToF", t, func() {
		Convey("When given an input in celsius", func() {
			Convey("It should convert that input to fahrenheit", func() {
				tests := []struct{
					c float32
					f float32
				} {
					{0, 32},
					{1, 33.8},
					{10, 50},
					{13.5, 56.3},
					{52.8, 127.04},
				}

				for _, test := range tests {
					So(test.f, ShouldAlmostEqual, cToF(test.c), 0.001)
				}
			})
		})
	})
}

func Test_floatFrom88(t *testing.T) {
	Convey("Function: floatFrom88", t, func() {
		Convey("Given an encoded int input", func() {
			Convey("It should decode the value into a float value", func() {
				tests := []struct{
					i int16
					o float32
				} {
					{256, 1},
					{12, 0.046875},
					{-512, -2},
					{-58, -0.2265625},
				}

				for _, test := range tests {
					So(test.o, ShouldAlmostEqual, floatFrom88(test.i), 0.001)
				}
			})
		})
	})
}

func TestByteValue_Id(t *testing.T) {
	Convey("Method: byteValue.Id", t, func() {
		Convey("Given an internal byte id", func() {
			Convey("It should return that internal byte id", func() {
				for i := 0; i < 256; i++ {
					So(i, ShouldEqual, byteValue{byte(i), ""}.Id())
				}
			})
		})
	})
}

func TestByteValue_Value(t *testing.T) {
	Convey("Method: byteValue.Value", t, func() {
		Convey("Given an internal string value", func() {
			Convey("It should return that internal string value", func() {
				if v, e := uuid.NewUUID(); e != nil {
					panic(e)
				} else {
					So(v.String(), ShouldEqual, byteValue{0, v.String()}.Value())
				}
			})
		})
	})
}

func TestByteValue_String(t *testing.T) {
	Convey("Method: byteValue.String", t, func() {
		Convey("Given any internal field values", func() {
			Convey("It should return those values in the expected string form", func() {
				tests := []struct{
					v1  byte
					v2  string
					out string
				} {
					{0x03, "foo", "0x03: foo"},
					{0xA6, "bar", "0xa6: bar"},
					{0x65, "fiz", "0x65: fiz"},
				}

				for _, test := range tests {
					So(test.out, ShouldEqual, byteValue{test.v1, test.v2}.String())
				}
			})
		})
	})
}

func TestOffset_inc(t *testing.T) {
	Convey("Method: offset.inc", t, func() {
		Convey("Given a starting uint value", func() {
			Convey("It should increase that value by one", func() {
				var i offset
				i.inc()
				So(1, ShouldEqual, i)
			})

			Convey("It should return the previous value", func() {
				var i offset
				So(0, ShouldEqual, i.inc())
			})
		})
	})
}
