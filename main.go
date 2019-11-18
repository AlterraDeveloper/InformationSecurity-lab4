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
	publicKey := privateKey.SavePublicKeyToFile("public.txt")

	fmt.Println("Ключи успешно сгенерированы")
	fmt.Printf("Открытый ключ : %v\n", publicKey)
	fmt.Printf("Личный ключ : %+v\n", privateKey)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите текст : ")
	scanner.Scan()
	text := scanner.Text()

	encryptionResult := Encrypt(text, publicKey)
	encryptionResultString := IntSliceToString(encryptionResult)

	fmt.Printf("Зашифрованный текст : %v\n", encryptionResult)
	fmt.Printf("Зашифрованный текст как строка : %v\n", encryptionResultString)
	fmt.Printf("Расшифрованный текст : %v\n", Decrypt(encryptionResult, privateKey))

	fmt.Printf("Зашифрованный текст в числовом виде : %v\n", StringToIntSlice(encryptionResultString))
	//добавить в метод чтения из файла
	// file, _ := os.Open("public.txt")
	// bytes := make([]byte, 2000000, 2000000)
	// bytesRead, _ := file.Read(bytes)
	// fmt.Printf("Bytes = %v\n", string(bytes[:bytesRead]))

}
