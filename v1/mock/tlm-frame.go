package mock

type  TlmFrame struct {
	Frame

	MTlmVersion struct {
		calls uint
		rv1   byte
	}

	MBatteryVoltage struct {
		calls uint
		rv1   uint16
	}

	MBeaconTemperatureInC struct {
		calls uint
		rv1   float32
	}

	MBeaconTemperatureInF struct {
		calls uint
		rv1   float32
	}

	MAdvertisementCount struct {
		calls uint
		rv1   uint32
	}

	MActiveTime struct {
		calls uint
		rv1   float32
	}

	MSupportsBatteryVoltage struct {
		calls uint
		rv1   bool
	}

	MSupportsTemperature struct {
		calls uint
		rv1   bool
	}
}

func (m *TlmFrame) TlmVersion() byte {
	m.MTlmVersion.calls++
	return m.MTlmVersion.rv1
}

func (m *TlmFrame) BatteryVoltage() uint16 {
	m.MBatteryVoltage.calls++
	return m.MBatteryVoltage.rv1
}

func (m *TlmFrame) BeaconTemperatureInC() float32 {
	m.MBeaconTemperatureInC.calls++
	return m.MBeaconTemperatureInC.rv1
}

func (m *TlmFrame) BeaconTemperatureInF() float32 {
	m.MBeaconTemperatureInF.calls++
	return m.MBeaconTemperatureInF.rv1
}

func (m *TlmFrame) AdvertisementCount() uint32 {
	m.MAdvertisementCount.calls++
	return m.MAdvertisementCount.rv1
}

func (m *TlmFrame) ActiveTime() float32 {
	m.MActiveTime.calls++
	return m.MActiveTime.rv1
}

func (m *TlmFrame) SupportsBatteryVoltage() bool {
	m.MSupportsBatteryVoltage.calls++
	return m.MSupportsBatteryVoltage.rv1
}

func (m *TlmFrame) SupportsTemperature() bool {
	m.MSupportsTemperature.calls++
	return m.MSupportsTemperature.rv1
}
