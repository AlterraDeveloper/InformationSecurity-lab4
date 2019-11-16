package main

import (
	"fmt"
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
	publicKey := privateKey.SavePublicKeyToFile("public.txt")

	fmt.Printf("Публичный ключ : %v\n", publicKey)
	// fmt.Printf("Личный ключ : %+v\n", privateKey)
	fmt.Println("Ключи успешно сгенерированы")

}
