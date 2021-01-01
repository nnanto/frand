// frand generates fast random unsigned integers (non-cryptographic).
// Uses Marsaglia's XOR shift similar to the frand function in go runtime
// Ref: Xorshift paper: https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf
package frand

import "time"

var state = uint32(time.Now().Unix() & ((1 << 31)-1))

// Get returns a random number with Marsaglia's favorite shift triplet [13,17,5].
// Ref: Xorshift paper: https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf
func Get() uint32 {
	state ^= (state << 13) ^ (state >> 17) ^ (state << 5)
	return state
}

// GetWith returns a random number with custom shift triplet.
// Ref: https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf
func GetWith(triplets [3]int) uint32 {
	state ^= (state << triplets[0]) ^ (state >> triplets[1]) ^ (state << triplets[2])
	return state
}

// GetN returns a random number between [0,n). if n == 0, returns 0.
// Ref: https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
func GetN(n uint32) uint32 {
	return uint32(uint64(Get()) * uint64(n) >> 32)
}

// GetNWith returns a random number between [0,n) with custom shift triplet.
// Ref: https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
func GetNWith(n uint32, triplets [3]int) uint32 {
	return uint32(uint64(GetWith(triplets)) * uint64(n) >> 32)
}
