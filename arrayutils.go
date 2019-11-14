package main

//Sum вычисление суммы элементов массива
func Sum(array []uint64) uint64 {
	var sum uint64 = 0
	for _, val := range array {
		sum += val
	}
	return sum
}
