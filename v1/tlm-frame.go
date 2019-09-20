package eddystone

// TlmFrame provides access to parsed Eddystone Telemetry
// data.
type TlmFrame interface {
	Frame

	TlmVersion() byte

	// BatteryVoltage returns the current battery charge in
	// millivolts (mV).
	//
	// Not all beacons support battery voltage measurements.
	// Use SupportsBatteryVoltage() to confirm that the
	// current beacon does before using this value.
	BatteryVoltage() uint16

	// BeaconTemperatureInC returns the current temperature of
	// the beacon in Celsius.
	//
	// Not all beacons support temperature measurements.  Use
	// SupportsTemperature() to confirm that the current
	// beacon does before using this value.
	BeaconTemperatureInC() float32

	// BeaconTemperatureInC returns the current temperature of
	// the beacon in Fahrenheit.
	//
	// Not all beacons support temperature measurements.  Use
	// SupportsTemperature() to confirm that the current
	// beacon does before using this value.
	BeaconTemperatureInF() float32

	// AdvertisementCount returns the number of advertisement
	// packets sent by the beacon since last reboot.
	AdvertisementCount() uint32

	// ActiveTime returns the uptime, or time since last
	// reboot for the broadcasting beacon.
	//
	// Value returned has a precision of 1/10 of a second.
	ActiveTime() float32

	// SupportsBatteryVoltage returns whether or not the
	// beacon that sent this frame sends battery voltage data.
	SupportsBatteryVoltage() bool

	// SupportsTemperature returns whether or not the beacon
	// that sent this frame sends temperature data.
	SupportsTemperature() bool
}
