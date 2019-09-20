package eddystone

import "fmt"

type UrlPrefix = byteValue

const (
	errInvalidUrlPrefix = "unrecognized url prefix byte 0x%x"
)

var (
	UrlPrefixHttpFull   = UrlPrefix{0x00, "http://www."}
	UrlPrefixHttpsFull  = UrlPrefix{0x01, "https://www."}
	UrlPrefixHttpShort  = UrlPrefix{0x02, "http://"}
	UrlPrefixHttpsShort = UrlPrefix{0x03, "https://"}
)

var prefixMap = [...]*UrlPrefix{
	&UrlPrefixHttpFull,
	&UrlPrefixHttpsFull,
	&UrlPrefixHttpShort,
	&UrlPrefixHttpsShort,
}

func IsUrlPrefix(b byte) bool {
	return b <= 0x03
}

func ParseUrlPrefix(b byte) *UrlPrefix {
	return prefixMap[b]
}


func newUrlPrefixErr(b byte) error {
	return fmt.Errorf(errInvalidUrlPrefix, b)
}
