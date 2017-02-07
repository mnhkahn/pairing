package pairing

import (
	"fmt"
	"testing"
)

var TEST_PAIRS = [][]uint64{
	[]uint64{0, 0},
	[]uint64{0, 1},
	[]uint64{1, 0},
}

var TEST_RESs = []uint64{
	0,
	2,
	1,
}

var TEST_RESBits = []uint64{
	0,
	1,
	4294967296,
}

func TestPair(t *testing.T) {
	for i, p := range TEST_PAIRS {
		a, b := p[0], p[1]
		if pair := Encode(a, b); pair != TEST_RESs[i] {
			t.Error(a, b, pair)
		}
	}

	for i, p := range TEST_PAIRS {
		a, b := p[0], p[1]
		pair := TEST_RESs[i]
		if x, y := Decode(pair); x != a || y != b {
			t.Error(a, b, pair)
		}
	}

	fmt.Println(Encode(559, 83792))

	for i, p := range TEST_PAIRS {
		a, b := uint32(p[0]), uint32(p[1])
		if pair := EncodeBit(a, b); pair != TEST_RESBits[i] {
			t.Error(a, b, pair)
		}
	}

	for i, p := range TEST_PAIRS {
		a, b := uint32(p[0]), uint32(p[1])
		pair := TEST_RESBits[i]
		if x, y := DecodeBit(pair); x != a || y != b {
			t.Error(a, b, pair)
		}
	}

	fmt.Println(EncodeBit(559, 83792))
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(559, 83792)
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode(3557671568)
	}
}

func BenchmarkEncodeBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeBit(559, 83792)
	}
}

func BenchmarkDecodeBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DecodeBit(2400886802256)
	}
}

func BenchmarkEncodeStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d_%d", 559, 83792)
	}
}

/*
go test -bench=. -benchmem
BenchmarkEncode-2       2000000000               0.37 ns/op            0 B/op          0 allocs/op
BenchmarkDecode-2       50000000                26.9 ns/op             0 B/op          0 allocs/op
BenchmarkEncodeBit-2    2000000000               0.37 ns/op            0 B/op          0 allocs/op
BenchmarkDecodeBit-2    2000000000               0.37 ns/op            0 B/op          0 allocs/op
BenchmarkEncodeStr-2     5000000               271 ns/op              32 B/op          3 allocs/op
*/
