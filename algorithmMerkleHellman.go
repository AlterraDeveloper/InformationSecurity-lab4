package main

import (
	"math/rand"
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
	key.X = uint64(rand.Int63n(int64(key.M/2)) + 2)
}

//GeneratePublicKey returns public key for private key
func (key PrivateKey) GeneratePublicKey() []uint64 {
	var publicKey []uint64

	for _, w := range key.W {
		publicKey = append(publicKey, (w*key.X)%key.M)
	}

	return publicKey
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
