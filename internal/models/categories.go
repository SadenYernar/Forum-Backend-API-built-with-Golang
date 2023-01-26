package models

type Category int64

const (
	Coding  = 0b0000001
	Music   = 0b0000010
	Art     = 0b0000100
	Sports  = 0b0001000
	Cooking = 0b0010000
	Other   = 0b0100000
	All     = 0b1000000
)
