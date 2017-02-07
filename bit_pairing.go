package pairing

func EncodeBit(k1, k2 uint32) uint64 {
	pair := uint64(k1)<<32 | uint64(k2)
	return pair
}

func DecodeBit(pair uint64) (uint32, uint32) {
	k1 := uint32(pair >> 32)
	k2 := uint32(pair) & 0xFFFFFFFF
	return k1, k2
}
