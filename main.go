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
	fmt.Println("Открытый ключ сохранен в файле public.txt")
	// fmt.Printf("Открытый ключ : %v\n", publicKey)
	// fmt.Printf("Личный ключ : %+v\n", privateKey)

	for {
		var command string
		fmt.Print("Введите 1 - для шифрования 2 - для дешифрования 0 - выход из программы : ")
		fmt.Scan(&command)

		if command == "0" {
			break
		}
		if command == "1" {
			var text string
			var inputMethod string
			for {
				fmt.Print("Как ввести текст? 1 - из файла 2 - с клавиатуры : ")
				fmt.Scan(&inputMethod)
				switch inputMethod {
				case "1":
					var filename string
					fmt.Print("Введите имя файла : ")
					fmt.Scan(&filename)
					text = getTextFromFile(filename)
				case "2":
					scanner := bufio.NewScanner(os.Stdin)
					fmt.Print("Введите текст : ")
					scanner.Scan()
					text = scanner.Text()
				}
				if len(text) > 0 {
					break
				}
			}

			encryptedText := IntSliceToString(Encrypt(text, publicKey))
			// fmt.Printf("Зашифрованный текст : %v\n", encryptedText)
			fmt.Printf("Зашифрованный текст : %v\n", Encrypt(text, publicKey))
			fmt.Print("Сохранить зашифрованный текст в файл? [Y/N] : ")
			fmt.Scan(&inputMethod)
			if inputMethod == "Y" || inputMethod == "y" {
				var filename string
				fmt.Print("Введите имя файла : ")
				fmt.Scan(&filename)
				saveTextToFile(filename, encryptedText)
				fmt.Printf("Зашифрованный текст успешно сохранен в файл %v\n", filename)
			}
		}

		if command == "2" {
			var filename string
			fmt.Print("Введите имя файла : ")
			fmt.Scan(&filename)
			encryptedText := getTextFromFile(filename)
			fmt.Printf("Расшифрованный текст : %v\n", Decrypt(StringToIntSlice(encryptedText), privateKey))

		}
	}

	// fmt.Printf("Зашифрованный текст : %v\n", encryptionResult)
	// fmt.Printf("Расшифрованный текст : %v\n", Decrypt(encryptionResult, privateKey))

	// fmt.Printf("Зашифрованный текст в числовом виде : %v\n", StringToIntSlice(encryptionResultString))

}

func getTextFromFile(filename string) string {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes := make([]byte, 2000000, 2000000)
	bytesRead, err := file.Read(bytes)
	if err != nil {
		panic(err)
	}

	return string(bytes[:bytesRead])
}

func saveTextToFile(filename, text string) {

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		panic(err)
	}

}
