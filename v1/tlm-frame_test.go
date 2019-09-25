package eddystone

import bean "github.com/Foxcapades/Beanies/v2"

type mockTlmFrame struct {
	mockFrame

	MTlmVersion             bean.ByteGetter
	MBatteryVoltage         bean.Uint16Getter
	MBeaconTemperatureInC   bean.Float32Getter
	MBeaconTemperatureInF   bean.Float32Getter
	MAdvertisementCount     bean.Uint32Getter
	MActiveTime             bean.Float32Getter
	MSupportsBatteryVoltage bean.BoolGetter
	MSupportsTemperature    bean.BoolGetter
}

func (t *mockTlmFrame) TlmVersion() byte {
	return t.MTlmVersion.Get()
}

func (t *mockTlmFrame) BatteryVoltage() uint16 {
	return t.MBatteryVoltage.Get()
}

func (t *mockTlmFrame) BeaconTemperatureInC() float32 {
	return t.MBeaconTemperatureInC.Get()
}

func (t *mockTlmFrame) BeaconTemperatureInF() float32 {
	return t.MBeaconTemperatureInF.Get()
}

func (t *mockTlmFrame) AdvertisementCount() uint32 {
	return t.MAdvertisementCount.Get()
}

func (t *mockTlmFrame) ActiveTime() float32 {
	return t.MActiveTime.Get()
}

func (t *mockTlmFrame) SupportsBatteryVoltage() bool {
	return t.MSupportsBatteryVoltage.Get()
}

func (t *mockTlmFrame) SupportsTemperature() bool {
	return t.MSupportsTemperature.Get()
}
