package main

import (
	"fmt"
	rand "math/rand"
)

func main() {

	var keyLength int32 = 0
	fmt.Print("Enter number : ")
	fmt.Scan(&keyLength)

	var openKey = make([]int32, keyLength, keyLength)

	for index := range openKey {
		openKey[index] = rand.Int31n(1000)
	}

	for _, value := range openKey {
		fmt.Print(value, " ")
	}
	fmt.Println()

	var privateKey = getPrivateKey(int(keyLength))

	for _, value := range privateKey {
		fmt.Print(value, " ")
	}

}

func getPrivateKey(size int) []int32 {
	var privateKey = make([]int32, size, size)
	privateKey[0] = 1
	privateKey[1] = 2

	for i := 2; i < size; i++ {
		var tmp int32 = 0
		for {
			sum := sum(privateKey[:i])
			tmp = rand.Int31n(sum + 10)
			if tmp > sum {
				break
			}
		}
		privateKey[i] = tmp
	}
	return privateKey
}

func sum(array []int32) int32 {
	var sum int32 = 0
	for _, value := range array {
		sum += value
	}
	return sum
}
