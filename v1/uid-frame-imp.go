package eddystone

import (
	"fmt"
	"github.com/google/uuid"
	"math/big"
	"strings"
)


/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏  Internal Constants                                    ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/


const (
	uidFrameLen   = 31 // TODO: confirm this is how it is reported...
	uidNSpaceLen  = 10
	uidIdenLen    = 6

	uidFormatString = "UidFrame{Type: 0x%x, TxPower: %d, UUID: %s}"
	uidFormatJson   = `{` +
		`"type":"%s",` +
		`"typeId": %d,` +
		`"txPower": %d,` +
		`"uuid": "%s"` +
		`}`

	errUidBadLen  = "invalid UID frame length; expected %d, got %d"
	errUidBadType = "invalid UID frame type; expected 0x%x, got 0x%x"
)


/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏  Internal Implementation                               ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/


type uidFrame struct {
	power  uint8
	nSpace [uidNSpaceLen]byte
	iden   [uidIdenLen]byte
}

func (u uidFrame) String() string {
	return fmt.Sprintf(uidFormatString, u.Type().Id(), u.TxPower(), u.Uuid())
}

func (u uidFrame) MarshalJSON() ([]byte, error) {
	tpe := u.Type()
	return []byte(fmt.Sprintf(
		uidFormatJson,
		tpe.Value(),
		tpe.Id(),
		u.TxPower(),
		u.Uuid(),
	)), nil
}

func (u uidFrame) Type() FrameType {
	return FrameTypeUid
}

func (u *uidFrame) ToBytes() (out []byte) {
	var off offset

	out = make([]byte, uidFrameLen)

	out[off.inc()] = u.Type().Id()

	for i := 0; i < uidNSpaceLen; i++ {
		out[off.inc()] = u.nSpace[i]
	}

	for i := 0; i < uidIdenLen; i++ {
		out[off.inc()] = u.iden[i]
	}

	return
}

func (u *uidFrame) FromBytes(b []byte) error {
	if len(b) != uidFrameLen {
		return fmt.Errorf(errUidBadLen, uidFrameLen, len(b))
	}

	var off offset
	z := off.inc()

	if b[z] != u.Type().Id() {
		return fmt.Errorf(errUidBadType, u.Type().Id(), b[z])
	}

	u.power = b[off.inc()]

	for i := 0; i < uidNSpaceLen; i++ {
		u.nSpace[i] = b[off.inc()]
	}

	for i := 0; i < uidIdenLen; i++ {
		u.iden[i] = b[off.inc()]
	}

	return nil
}

func (u uidFrame) TxPower() int8 {
	return int8(u.power)
}

func (u uidFrame) RangingData() int8 {
	return u.TxPower()
}

func (u *uidFrame) NamespaceBytes() (out [uidNSpaceLen]byte) {
	for i := range u.nSpace {
		out[i] = u.nSpace[i]
	}
	return
}

func (u *uidFrame) NamespaceString() string {
	var sb strings.Builder
	for _, v := range u.nSpace {
		byteToHex(v, &sb)
	}
	return sb.String()
}

func (u *uidFrame) NamespaceInt() *big.Int {
	return new(big.Int).SetBytes(u.nSpace[:])
}

func (u *uidFrame) InstanceBytes() (out [6]byte) {
	for i := range u.iden {
		out[i] = u.iden[i]
	}
	return
}

func (u *uidFrame) InstanceString() string {
	var sb strings.Builder
	for _, v := range u.iden {
		byteToHex(v, &sb)
	}
	return sb.String()
}

func (u *uidFrame) InstanceInt() (out uint64) {
	for i, v := range u.iden {
		out |= uint64(v) << uint(i*8)
	}
	return
}

func (u *uidFrame) Uuid() (o uuid.UUID) {
	size := uidNSpaceLen + uidIdenLen
	buf  := make([]byte, size)
	for _, v := range u.nSpace {
		buf = append(buf, v)
	}
	for _, v := range u.iden {
		buf = append(buf, v)
	}
	o, e := uuid.FromBytes(buf)
	if e != nil {
		panic(e)
	}
	return
}
