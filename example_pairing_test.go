// Package pairing_test

package pairing_test

import (
	"github.com/mnhkahn/pairing"
	"math"
)

func Example() {
	// Encode two numbers
	pair := pairing.Encode(1, 2)
	// Decode to two numbers
	a, b := pairing.Decode(pair)
}

func ExampleEncode() {
	pairing.Encode(1, 2)
}

func ExampleDecode() {
	pairing.Decode(3)
}

func ExampleEncodeBit() {
	pairing.EncodeBit(1,2)
}

func ExampleDecodeBit() {
	pairing.DecodeBit(math.MaxUint64)
}