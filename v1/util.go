package eddystone

import (
	"fmt"
	"math"
	"strings"
)

const maxByte = math.MaxUint8

func u32ToByte(i uint32) byte {
	return byte(i & maxByte)
}

func u16ToByte(i uint16) byte {
	return byte(i & maxByte)
}

func cToF(c float32) (f float32) {
	return c * 1.8 + 32
}

func floatFrom88(i int16) (o float32) {
	return float32(i) / float32(256)
}

type offset uint32
func (o *offset) inc() (out uint32) {
	out = uint32(*o)
	*o  = offset(out + 1)
	return
}

func byteToHex(i byte, sb *strings.Builder) {
	sb.WriteByte(digitToHex(i / 16))
	sb.WriteByte(digitToHex(i % 16))
}

func digitToHex(i byte) byte {
	if i < 10 {
		return '0' + i
	} else if i < 16 {
		return 'a' + (i - 10)
	}
	panic("cannot convert values greater than 15")
}

type byteValue struct {
	id  byte
	val string
}

func (b byteValue) Id() byte {
	return b.id
}

func (b byteValue) Value() string {
	return b.val
}

func (b byteValue) String() string {
	return fmt.Sprintf("0x%x: %s", b.id, b.val)
}
