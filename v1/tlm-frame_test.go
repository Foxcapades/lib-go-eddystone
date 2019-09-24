package eddystone

type mockTlmFrame struct {
	mockFrame
	mTlmVersion struct {
		calls uint
		rv1   byte
	}
	mBatteryVoltage struct {
		calls uint
		rv1   uint16
	}
	mBeaconTemperatureInC struct {
		calls uint
		rv1   float32
	}
	mBeaconTemperatureInF struct {
		calls uint
		rv1   float32
	}
	mAdvertisementCount struct {
		calls uint
		rv1   uint32
	}
	mActiveTime struct {
		calls uint
		rv1   float32
	}
	mSupportsBatteryVoltage struct {
		calls uint
		rv1   bool
	}
	mSupportsTemperature struct {
		calls uint
		rv1   bool
	}
}

func (m *mockTlmFrame) TlmVersion() byte {
	m.mTlmVersion.calls++
	return m.mTlmVersion.rv1
}

func (m *mockTlmFrame) BatteryVoltage() uint16 {
	m.mBatteryVoltage.calls++
	return m.mBatteryVoltage.rv1
}

func (m *mockTlmFrame) BeaconTemperatureInC() float32 {
	m.mBeaconTemperatureInC.calls++
	return m.mBeaconTemperatureInC.rv1
}

func (m *mockTlmFrame) BeaconTemperatureInF() float32 {
	m.mBeaconTemperatureInF.calls++
	return m.mBeaconTemperatureInF.rv1
}

func (m *mockTlmFrame) AdvertisementCount() uint32 {
	m.mAdvertisementCount.calls++
	return m.mAdvertisementCount.rv1
}

func (m *mockTlmFrame) ActiveTime() float32 {
	m.mActiveTime.calls++
	return m.mActiveTime.rv1
}

func (m *mockTlmFrame) SupportsBatteryVoltage() bool {
	m.mSupportsBatteryVoltage.calls++
	return m.mSupportsBatteryVoltage.rv1
}

func (m *mockTlmFrame) SupportsTemperature() bool {
	m.mSupportsTemperature.calls++
	return m.mSupportsTemperature.rv1
}
