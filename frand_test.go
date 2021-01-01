package frand

import (
	"testing"
)

func TestGetN(t *testing.T) {
	samples := []uint32{0, 10, 100, 1 << 31}

	for i := 0; i < len(samples); i++ {
		s := samples[i]
		x := GetN(s)
		if x > s {
			t.Errorf("GetN(%v) is greater than %v", s, s)
		}
	}
}

func TestGetNWith(t *testing.T) {
	samples := []uint32{0, 10, 100, 1 << 31}

	for i := 0; i < len(samples); i++ {
		s := samples[i]
		x := GetNWith(s, [3]int{7,25,20})
		if x > s {
			t.Errorf("GetN(%v) is greater than %v", s, s)
		}
	}
}

func BenchmarkGetN(b *testing.B) {
	f := uint32(1)
	for i := 0; i < b.N; i++ {
		f *= GetN(100)
	}
	_ = f
}
