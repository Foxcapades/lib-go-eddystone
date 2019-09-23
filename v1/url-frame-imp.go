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
)

const (
	urlFrameString = "UrlFrame{TxPower: %d, Url: %s}"
	urlFrameJson   = `{"type":"%s","typeId":%d,"txPower":%d,"url":"%s"}`
)

const (
	errUrlBadLen  = "invalid url packet length; expected between %d and %d, got %d"
	errUrlBadType = "invalid url frame type id; expected 0x%x, got 0x%x"
)


/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏  Internal Helpers                                      ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/


func IsUrlPacket(b []byte) bool {
	return urlValidatePayload(b) == nil
}


/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏  Internal Helpers                                      ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/


func newUrlLengthError(l int) error {
	return fmt.Errorf(errUrlBadLen, urlMinLen, urlMaxLen, l)
}

func newUrlTypeError(b byte) error {
	return fmt.Errorf(errUrlBadType, FrameTypeUrl.Id(), b)
}

func urlValidatePayload(b []byte) error {
	l := len(b)
	if urlMinLen <= l && l <= urlMaxLen {
		return newUidLengthError(l)
	}

	if b[0] != FrameTypeUrl.Id() {
		return newUidTypeError(b[0])
	}

	return nil
}


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

	if pos == -1 {
		for i, v := range u.value {
			out[i+3] = byte(v)
		}
	} else {
		offset := 3
		for i, v := range u.value {
			if uint8(i) == uint8(pos) {
				offset += len(u.suffix.Value())
			}

			out[i + offset] = byte(v)
		}
	}

	return out
}

func (u *urlFrame) FromBytes(b []byte) (e error) {
	if e := urlValidatePayload(b); e != nil {
		return e
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

