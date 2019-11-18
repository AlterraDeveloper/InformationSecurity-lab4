package main

import (
	"math"
)

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
func XGCDIterative(a, b int64) (int64, int64, int64) {

	var q, r, x1, x2, y1, y2 int64

	if b == 0 {
		return 1, 0, a
	}

	x1 = 0
	x2 = 1
	y1 = 1
	y2 = 0
	q = 0
	r = 0

	var x, y int64
	x = 0
	y = 0

	for b > 0 {
		q = a / b
		r = a - q*b
		x = x2 - q*x1
		y = y2 - q*y1
		a = b
		b = r
		x2 = x1
		x1 = x
		y2 = y1
		y1 = y
	}
	return x2, y2, a
}

//InverseByMod ...
func InverseByMod(number, module uint64) uint64 {
	a, _, gcd := XGCDIterative(int64(number), int64(module))
	if gcd == 1 {
		for a < 0 {
			a += int64(module)
		}
		return uint64(a)
	}
	return 0
}
