package main

import (
	"math/rand"
	"os"
	"strconv"
)

//PrivateKey закрытый ключ
type PrivateKey struct {
	W []uint64
	M uint64
	X uint64
}

//Generate private key
func (key *PrivateKey) Generate(keyLength int) {
	key.W = generateW(keyLength)
	key.M = generateM(Sum(key.W))
	key.X = uint64(rand.Int63n(int64(key.M/2))) + 2
}

func (key PrivateKey) generatePublicKey() []uint64 {
	var publicKey []uint64

	for _, w := range key.W {
		publicKey = append(publicKey, (w*key.X)%key.M)
	}

	return publicKey
}

//SavePublicKeyToFile saves generated public key to file and return it
func (key PrivateKey) SavePublicKeyToFile(fileName string) []uint64 {

	file, _ := os.Create(fileName)

	publicKey := key.generatePublicKey()

	for _, value := range publicKey {
		var buffer []byte
		buffer = strconv.AppendUint(buffer, value, 10)
		_, _ = file.Write(buffer)
		file.WriteString(" ")

	}
	return publicKey
}

//Encrypt text
func Encrypt(originalText string, publicKey []uint64) []uint64 {

	// var encryptedText string

	keyLength := len(publicKey)

	increasedText := IncreaseText(originalText, keyLength)

	bitsStream := RunesToBits([]rune(increasedText))

	var encryptedNums []uint64

	for i := 0; i < len(bitsStream); i += keyLength {

		block := bitsStream[i : i+keyLength]

		var sum uint64 = 0
		for i, bit := range block {
			if bit == rune('1') {
				sum += publicKey[i]
			}
		}

		encryptedNums = append(encryptedNums, sum)
	}

	return encryptedNums
}

//Decrypt nums to text
func Decrypt(nums []uint64, privateKey PrivateKey) string {

	var bitsStream string

	y := InverseByMod(privateKey.X, privateKey.M)

	if (privateKey.X*y)%privateKey.M == 1 {
		for _, value := range nums {
			var bits string
			tmp := (value * y) % privateKey.M
			for i := range privateKey.W {
				current := privateKey.W[len(privateKey.W)-1-i]
				if tmp >= current {
					bits += "1"
					tmp -= current
				} else {
					bits += "0"
				}
			}
			bits = Reverse(bits)
			bitsStream += bits
		}

		return string(BitsToRunes(bitsStream))
	}
	return "Не удалось расшифровать..."
}

func generateW(length int) []uint64 {
	var superIncreasingSequence []uint64
	superIncreasingSequence = append(superIncreasingSequence, 1, 2)

	for i := 2; i < length; i++ {
		sum := Sum(superIncreasingSequence[:i])
		next := uint64(rand.Int31n(10)) + sum + uint64(1)
		superIncreasingSequence = append(superIncreasingSequence, next)
	}
	return superIncreasingSequence
}

func generateM(sumOfSequence uint64) uint64 {
	for {
		sumOfSequence++
		if NumberIsSimple(sumOfSequence) {
			break
		}
	}
	return sumOfSequence
}
