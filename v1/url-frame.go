package eddystone

type UrlFrame interface {
	Frame

	TxPower() int8

	Url() string
}
