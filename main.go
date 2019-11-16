package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var keyLengthInput string
	var keyLength int = 0

	for {
		fmt.Print("Введите длину ключа[4-56]: ")
		fmt.Scan(&keyLengthInput)
		_keyLength, err := strconv.Atoi(keyLengthInput)
		if err == nil && _keyLength >= 4 && _keyLength <= 56 {
			keyLength = _keyLength
			break
		}
	}

	fmt.Println("Генерация ключей...")

	var privateKey PrivateKey
	privateKey.Generate(keyLength)
	// privateKey = PrivateKey{W: []uint64{2, 7, 11, 21, 42, 89, 180, 354}, M: 881, X: 588}
	// privateKey = PrivateKey{W: []uint64{1, 2, 4, 9, 20, 38, 75, 150}, M: 311, X: 52}
	publicKey := privateKey.SavePublicKeyToFile("public.txt")

	fmt.Println("Ключи успешно сгенерированы")
	fmt.Printf("Открытый ключ : %v\n", publicKey)
	fmt.Printf("Личный ключ : %+v\n", privateKey)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите текст : ")
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	text := scanner.Text()

	encryptionResult := Encrypt(text, publicKey)
	fmt.Printf("Зашифрованный текст : %v\n", encryptionResult)
	fmt.Printf("Расшифрованный текст : %v\n", Decrypt(encryptionResult, privateKey))

}
