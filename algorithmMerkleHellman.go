package main

import (
	"math/big"
	"math/rand"
	"os"
)

//PrivateKey закрытый ключ
type PrivateKey struct {
	W         []big.Int
	M         big.Int
	X         big.Int
	PublicKey []big.Int
}

//Generate private key
func (key *PrivateKey) Generate(keyLength int) {
	key.W = generateW(keyLength)
	key.M = *(Sum(key.W))
	key.X = generateX(key.M)
	key.PublicKey = key.generatePublicKey()
}

func (key PrivateKey) generatePublicKey() []big.Int {

	var publicKey []big.Int

	for _, w := range key.W {
		tmp := big.NewInt(0)
		tmp = tmp.Mul(&w, &(key.X))
		tmp = tmp.Mod(tmp, &(key.M))
		publicKey = append(publicKey, *tmp)
	}

	return publicKey
}

//SavePublicKeyToFile saves generated public key to file and return it
func (key PrivateKey) SavePublicKeyToFile(fileName string) {

	file, _ := os.Create(fileName)
	defer file.Close()

	file.WriteString(ToString(key.PublicKey))
}

//Encrypt text
func Encrypt(originalText string, publicKey []big.Int) []big.Int {

	keyLength := len(publicKey)

	increasedText := IncreaseText(originalText, keyLength)

	bitsStream := RunesToBits([]rune(increasedText))

	var encryptedNums []big.Int

	for i := 0; i < len(bitsStream); i += keyLength {

		block := bitsStream[i : i+keyLength]

		sum := big.NewInt(0)
		for i, bit := range block {
			if bit == rune('1') {
				sum = big.NewInt(0).Add(sum, &(publicKey[i]))
			}
		}

		encryptedNums = append(encryptedNums, *sum)
	}

	return encryptedNums
}

//Decrypt nums to text
func Decrypt(nums []big.Int, privateKey PrivateKey) string {

	var bitsStream string

	y := big.NewInt(0).ModInverse(&(privateKey.X), &(privateKey.M))

	if y != nil {
		for _, value := range nums {
			var bits string
			tmp := big.NewInt(0).Mul(&value, y)
			tmp = big.NewInt(0).Mod(tmp, &(privateKey.M))
			for i := len(privateKey.W) - 1; i >= 0; i-- {
				current := privateKey.W[i]
				if tmp.Cmp(&current) == 0 || tmp.Cmp(&current) == 1 {
					bits += "1"
					tmp = big.NewInt(0).Sub(tmp, &current)
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

func generateW(length int) []big.Int {

	var superIncreasingSequence []big.Int

	for i := 0; i < length; i++ {
		next := big.NewInt(1).Lsh(big.NewInt(1), uint(i))
		superIncreasingSequence = append(superIncreasingSequence, *next)
	}
	return superIncreasingSequence
}

func generateX(keyM big.Int) big.Int {
	genX := big.NewInt(0)
	rand := rand.New(rand.NewSource(24111998))
	divResult := genX.Div(&keyM, big.NewInt(2))
	genX = genX.Rand(rand, divResult)
	return *(genX.Add(genX, big.NewInt(2)))
}

//ToString represents private key as string
func (key PrivateKey) ToString() string {
	var output string

	output += "M : "
	output += key.M.String() + "\n"
	output += "X : "
	output += key.X.String() + "\n"
	output += "W : "
	for _, val := range key.W {
		output += val.Text(10)
		output += " "
	}

	return output
}
