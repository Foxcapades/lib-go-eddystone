package eddystone

import (
	"encoding/binary"
	"fmt"
)

/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏  Internal Constants                                    ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/


const (
	tlmSize = 14
	tlmVoltageDisabled uint16 = 0x0000
	tlmTempDisabled uint16 = 0x8000
)

const (
	jsonTlmFormat = `{` +
		`"type":"%s",` +
		`"typeId":"%d",` +
		`"voltage":"%d",` +
		`"beaconTempC":%.8f,` +
		`"advertisements":%d,` +
		`"uptime":%.8f` +
		`}`
	tlmFormatString = "TlmFrame{Type: 0x%x, Version: 0x%x, Voltage: %d, " +
		"Temp: %.8fC, AdCount: %d, Seconds: %.1f}"
)

const (
	errTlmPacketSize = "invalid tlm packet size; expected %d bytes, got %d"
	errTlmPacketType = "invalid tlm frame type id; expected 0x%x, got 0x%x"
)


/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏  Exported Functions                                    ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/

func IsTlmPacket(b []byte) bool {
	return len(b) > 0 && FrameTypeTlm.Id() == b[0]
}


/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏  Internal Helpers                                      ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/

func newErrTlmPacketSize(exp, act int) error {
	return fmt.Errorf(errTlmPacketSize, exp, act)
}
func newErrTlmPacketType(act byte) error {
	return fmt.Errorf(errTlmPacketType, FrameTypeTlm.Id(), act)
}

/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏  Internal Implementation                               ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/


type tlmFrame struct {
	version byte
	voltage uint16
	temp    uint16
	adCount uint32
	seconds uint32
}

func (t tlmFrame) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(
		jsonTlmFormat,
		FrameTypeTlm.Value(),
		FrameTypeTlm.Id(),
		t.BatteryVoltage(),
		t.BeaconTemperatureInC(),
		t.AdvertisementCount(),
		t.ActiveTime(),
	)), nil
}

func (t tlmFrame) String() string {
	return fmt.Sprintf(
		tlmFormatString,
		t.Type().Id(),
		t.TlmVersion(),
		t.BatteryVoltage(),
		t.BeaconTemperatureInC(),
		t.AdvertisementCount(),
		t.ActiveTime(),
	)
}

func (t tlmFrame) Type() FrameType {
	return FrameTypeTlm
}

func (t tlmFrame) ToBytes() (out []byte) {
	return []byte{
		t.Type().Id(),
		t.version,
		byte(t.voltage >> 8),
		u16ToByte(t.voltage),
		byte(t.temp >> 8),
		u16ToByte(t.temp),
		u32ToByte(t.adCount >> 24),
		u32ToByte(t.adCount >> 16),
		u32ToByte(t.adCount >> 8),
		u32ToByte(t.adCount),
		u32ToByte(t.seconds >> 24),
		u32ToByte(t.seconds >> 16),
		u32ToByte(t.seconds >> 8),
		u32ToByte(t.seconds),
	}
}

func (t *tlmFrame) FromBytes(b []byte) error {
	if len(b) != tlmSize {
		return newErrTlmPacketSize(tlmSize, len(b))
	}

	if b[0] != t.Type().Id() {
		return newErrTlmPacketType(b[0])
	}

	t.version = b[1]
	t.voltage = binary.BigEndian.Uint16(b[2:4])
	t.temp    = binary.BigEndian.Uint16(b[4:6])
	t.adCount = binary.BigEndian.Uint32(b[6:10])
	t.seconds = binary.BigEndian.Uint32(b[10:14])
	return nil
}

func (t tlmFrame) TlmVersion() byte {
	return t.version
}

func (t tlmFrame) BatteryVoltage() uint16 {
	return t.voltage
}

func (t tlmFrame) BeaconTemperatureInC() float32 {
	return floatFrom88(int16(t.temp))
}

func (t tlmFrame) BeaconTemperatureInF() float32 {
	return cToF(t.BeaconTemperatureInC())
}

func (t tlmFrame) AdvertisementCount() uint32 {
	return t.adCount
}

func (t tlmFrame) ActiveTime() float32 {
	return float32(t.seconds) / float32(10)
}

func (t tlmFrame) SupportsBatteryVoltage() bool {
	return t.voltage != tlmVoltageDisabled
}

func (t tlmFrame) SupportsTemperature() bool {
	return t.temp != tlmTempDisabled
}

