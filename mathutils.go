package main

import "math"

//NumberIsSimple checks number on simplicity
func NumberIsSimple(number uint64) bool {
	if number == 1 {
		return false
	}

	for i := uint64(2); i <= uint64(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

//GCDIterative made with Euclid's Algorithm
func GCDIterative(u, v uint64) uint64 {
	var t uint64
	for u > 0 {
		if u < v {
			t = u
			u = v
			v = t
		}
		u = u - v
	}
	return v
}

//XGCDIterative made with Extended Euclid's Algorithm
func XGCDIterative(a, b int64) int64 {
	var aa [2]int64 = [2]int64{1, 0}
	var bb [2]int64 = [2]int64{0, 1}
	var q int64

	for {
		q = a / b
		a = a % b
		aa[0] = aa[0] - q*aa[1]
		bb[0] = bb[0] - q*bb[1]

		if a == 0 {
			return int64(math.Max(float64(aa[1]), float64(bb[1])))
		}

		q = b / a
		b = b % a
		aa[1] = aa[1] - q*aa[0]
		bb[1] = bb[1] - q*bb[0]

		if b == 0 {
			return int64(math.Max(float64(aa[0]), float64(bb[0])))
		}
	}
}
