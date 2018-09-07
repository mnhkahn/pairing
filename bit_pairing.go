package pairing

// EncodeBit two uint32 numbers by bit.
// First 32 bits store k1 and last 32 bits store k2.
func EncodeBit(k1, k2 uint32) uint64 {
	pair := uint64(k1)<<32 | uint64(k2)
	return pair
}

// DecodeBit one uint64 pair to two uint32 numbers by bit.
func DecodeBit(pair uint64) (uint32, uint32) {
	k1 := uint32(pair >> 32)
	k2 := uint32(pair) & 0xFFFFFFFF
	return k1, k2
}
