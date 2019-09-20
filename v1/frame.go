package eddystone

import (
	"encoding/json"
	"fmt"
)

// Frame is the base type for all Eddystone frame data, on
// for development use, see TlmFrame, UrlFrame,
// or UidFrame.
type Frame interface {
	// Extensions

	fmt.Stringer
	json.Marshaler

	// Returns the type of the current frame.
	Type() FrameType

	// ToBytes returns this frame represented as a slice of
	// bytes compatible with the Eddystone TLM specification.
	ToBytes() []byte

	// FromBytes parses the given byte array and writes the
	// parsed values to the internal store for this Frame
	FromBytes(b []byte) error
}
