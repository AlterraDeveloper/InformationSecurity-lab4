package main

import (
	"math/big"
	"strings"
)

//Sum вычисление суммы элементов массива
func Sum(array []big.Int) *big.Int {
	sum := big.NewInt(0)
	for _, val := range array {
		sum = big.NewInt(0).Add(sum, &val)
	}
	return sum
}

//ToString represents []big.Int as string
func ToString(array []big.Int) string {
	var output string
	for _, val := range array {
		output += val.Text(10)
		output += " "
	}
	return output
}

//StringToBigIntSlice ...
func StringToBigIntSlice(str string) []big.Int {
	var sequence []big.Int

	numbers := strings.Fields(str)

	for _, number := range numbers {
		tmp, _ := big.NewInt(0).SetString(number, 10)
		if tmp != nil {
			sequence = append(sequence, *tmp)
		}
	}

	return sequence
}
