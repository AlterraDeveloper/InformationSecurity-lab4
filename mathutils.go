package main

import (
	"math/big"
)

//NumberIsSimple checks number on simplicity
func NumberIsSimple(number *big.Int) bool {
	if number.Cmp(big.NewInt(1)) == 0 {
		return false
	}

	i := big.NewInt(2)
	sqrt := number.Sqrt(number)

	for i.Cmp(sqrt) == 0 || i.Cmp(sqrt) == -1 {
		modResult := number.Mod(number, i)
		if modResult.Cmp(big.NewInt(0)) == 0 {
			return false
		}
	}

	return true
}
