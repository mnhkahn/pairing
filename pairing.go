/*
 * @Author: mnhkahn
 * @Wiki: https://en.wikipedia.org/wiki/Pairing_function#Cantor_pairing_function
 */

package pairing

import "math"

// Encode two uint64 numbers by cantor pairing function.
func Encode(k1, k2 uint64) uint64 {
	pair := k1 + k2
	pair = pair * (pair + 1)
	pair = pair / 2
	pair = pair + k2

	return pair
}

// Decode one uint64 pair to two uint64 numbers by cantor pairing function.
func Decode(pair uint64) (uint64, uint64) {
	w := math.Floor((math.Sqrt(float64(8*pair+1)) - 1) / 2)
	t := (w*w + w) / 2

	k2 := pair - uint64(t)
	k1 := uint64(w) - k2
	return k1, k2
}
