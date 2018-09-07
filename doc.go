/*
pairing is package can encode and decode integers, decoded integers can keep sorts.
It base on cantor pairing function.
Wiki: https://en.wikipedia.org/wiki/Pairing_function#Cantor_pairing_function.

Chinese detail: http://blog.cyeam.com//golang/2017/02/07/go-optimize-pair?utm_source=peanut.

Benchmark

	BenchmarkEncode-2       2000000000               0.37 ns/op            0 B/op          0 allocs/op
    BenchmarkDecode-2       50000000                26.9 ns/op             0 B/op          0 allocs/op
    BenchmarkEncodeBit-2    2000000000               0.37 ns/op            0 B/op          0 allocs/op
    BenchmarkDecodeBit-2    2000000000               0.37 ns/op            0 B/op          0 allocs/op
    BenchmarkEncodeStr-2     5000000               271 ns/op              32 B/op          3 allocs/op


 */
// Package pairing
package pairing
