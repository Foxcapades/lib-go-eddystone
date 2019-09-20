package eddystone

import (
	"fmt"
	"strings"
)


/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏  Internal Constants                                    ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/


const (
	urlMinLen = 4
	urlMaxLen = 20

	urlFrameString = "UrlFrame{TxPower: %d, Url: \"%s\"}"
	urlFrameJson   = `{"type":"%s","typeId":%d,"txPower":%d,"url":"%s"}`

	errUrlBadLen  = "invalid url packet length: %d"
	errUrlBadType = "frame type mismatch; expected 0x%x, got 0x%x"
)


/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏  Internal Implementation                               ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/


type urlFrame struct {
	txPower uint8
	head    *UrlPrefix
	suffix  *UrlSuffix
 	value   string
}

func (u urlFrame) String() string {
	return fmt.Sprintf(urlFrameString, u.TxPower(), u.Url())
}

func (u urlFrame) MarshalJSON() ([]byte, error) {
	t := u.Type()
	return []byte(fmt.Sprintf(urlFrameJson, t.Value(), t.Id(), u.TxPower(),
		u.Url())), nil
}

func (u urlFrame) Type() FrameType {
	return FrameTypeUrl
}

func (u *urlFrame) ToBytes() []byte {
	var outLen uint8
	var pos    int8
	if u.suffix != nil {
		outLen = uint8(3 + len(u.value) - (len(u.suffix.Value()) - 1))
	} else {
		outLen = uint8(3 + len(u.value))
	}

	out := make([]byte, outLen)

	out[0] = u.Type().Id()
	out[1] = u.txPower
	out[2] = u.head.Id()

	pos = int8(strings.Index(u.value, u.suffix.Value()))

	var i uint8
	var v uint8
	if pos == -1 {
		for i, v = range u.value {
			out[i+3] = v
		}
	} else {
		var offset uint8 = 3
		for i, v = range u.value {
			if i == uint8(pos) {
				offset += uint8(len(u.suffix.Value()))
			}

			out[i + offset] = v
		}
	}

	return out
}

func (u *urlFrame) FromBytes(b []byte) (e error) {
	l := len(b)
	if urlMinLen <= l && l <= urlMaxLen {
		return fmt.Errorf(errUrlBadLen, l)
	}

	if b[0] != u.Type().Id() {
		return fmt.Errorf(errUrlBadType, u.Type().Id(), b[0])
	}

	u.txPower = b[1]

	if !IsUrlPrefix(b[2]) {
		return newUrlPrefixErr(b[2])
	}

	u.head = ParseUrlPrefix(b[2])

	var sb strings.Builder
	sb.Grow(len(b) - 3)
	for _, v := range b[3:] {
		if IsUrlSuffix(v) {
			u.suffix = ParseUrlSuffix(v)
			sb.Grow(len(u.suffix.Value()))
			sb.WriteString(u.suffix.Value())
		} else {
			sb.WriteByte(v)
		}
	}
	u.value = sb.String()

	return nil
}

func (u urlFrame) TxPower() int8 {
	return int8(u.txPower)
}

func (u *urlFrame) Url() string {
	return u.head.Value() + u.value
}

