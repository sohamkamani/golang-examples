package main

import (
	"fmt"
)

const (
	SPADES = iota
	HEARTS
	DIAMONDS
	CLUBS
)

const (
	ZERO = 10 * iota
	TEN
	TWENTY
	THIRTY
)

const (
	BYTE     = 1 << (iota * 10) // = 1 << 0
	KILOBYTE                    // = 1 << 10
	MEGABYTE                    // = 1 << 20
	GIGABYTE                    // = 1 << 30
)

const (
	ERR_ABC = 1
	ERR_DEF = 2
	ERR_HIJ = 3
	ERR_KLM = 3
	ERR_NOP = 4
	ERR_QRS = 5
)

const (
	SECOND = 1
	MINUTE = 60 * SECOND
	HOUR   = 60 * MINUTE
	DAY    = 24 * HOUR
)

const (
	TOTAL_CARDS  = 52
	TOTAL_JOKERS = 2
)

func main() {
	fmt.Println("SPADES:", SPADES, "HEARTS:", HEARTS, "DIAMONDS:", DIAMONDS, "CLUBS:", CLUBS)
	fmt.Println("ZERO:", ZERO, "TEN:", TEN, "TWENTY:", TWENTY, "THIRTY:", THIRTY)
	fmt.Println("BYTE:", BYTE, "KILOBYTE:", KILOBYTE, "MEGABYTE:", MEGABYTE, "GIGABYTE:", GIGABYTE)
}
