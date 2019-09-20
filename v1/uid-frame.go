package eddystone

import (
	"github.com/google/uuid"
	"math/big"
)

type UidFrame interface {
	Frame

	// TxPower returns the advertised TX power measured in dBm
	TxPower() int8

	// RangingData returns the advertised TX power measured in
	// dBm
	//
	// RangingData is an alias of TxPower based on the naming
	// scheme from the Eddystone specification.
	RangingData() int8

	// NamespaceBytes returns the raw byte value of the UID
	// namespace for the source device.
	NamespaceBytes() [10]byte

	// NamespaceString returns the stringified hex value of
	// the UID namespace for the source device.
	NamespaceString() string

	// NamespaceInt returns the parsed big int value of the
	// UID namespace for the source device.
	NamespaceInt() *big.Int

	// InstanceBytes returns the raw byte value of the UID
	// instance identifier for the source device.
	InstanceBytes() [6]byte

	// InstanceString returns the stringified hex value of the
	// UID instance identifier for the source device.
	InstanceString() string

	// InstanceInt returns the parsed uint64 value of the UID
	// instance identifier for the source device.
	InstanceInt() uint64

	// Returns the full identifier value for the source device
	// as a UUID.
	Uuid() uuid.UUID
}
