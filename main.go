package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {

	const KeyMinLen int = 4
	const KeyMaxLen int = 100

	var keyLengthInput string
	var keyLength int = 0

	for {
		fmt.Printf("Введите длину ключа[%v-%v]: ", KeyMinLen, KeyMaxLen)
		fmt.Scan(&keyLengthInput)
		_keyLength, err := strconv.Atoi(keyLengthInput)
		if err == nil && _keyLength >= KeyMinLen && _keyLength <= KeyMaxLen {
			keyLength = _keyLength
			break
		}
	}

	fmt.Println("Генерация ключей...")
	var privateKey PrivateKey
	privateKey.Generate(keyLength)
	fmt.Println("Ключи успешно сгенерированы")
	fmt.Printf("Открытый ключ : %s\n", ToString(privateKey.PublicKey))
	privateKey.SavePublicKeyToFile("public.txt")
	fmt.Println("Открытый ключ сохранен в файле public.txt")

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
					fmt.Print("Введите имя файла содержимое которого хотите зашифровать : ")
					fmt.Scan(&filename)
					text = getTextFromFile(filename)
				case "2":
					scanner := bufio.NewScanner(os.Stdin)
					fmt.Print("Введите текст который хотите зашифровать : ")
					scanner.Scan()
					text = scanner.Text()
				}
				if len(text) > 0 {
					break
				}
			}

			encryptedSequence := ToString(Encrypt(text, privateKey.PublicKey))
			fmt.Printf("Зашифрованный текст : %v\n", encryptedSequence)
			fmt.Print("Сохранить зашифрованный текст в файл? [Y/N] : ")
			fmt.Scan(&inputMethod)
			if inputMethod == "Y" || inputMethod == "y" {
				var filename string
				fmt.Print("Введите имя файла : ")
				fmt.Scan(&filename)
				saveTextToFile(filename, encryptedSequence)
				fmt.Printf("Зашифрованный текст успешно сохранен в файл %v\n", filename)
			}
		}

		if command == "2" {
			var filename string
			fmt.Print("Введите имя файла с зашифрованными данными : ")
			fmt.Scan(&filename)
			encryptedTSequence := StringToBigIntSlice(getTextFromFile(filename))
			fmt.Printf("Расшифрованный текст :\n===================================\n%v\n===================================\n", Decrypt(encryptedTSequence, privateKey))

		}
	}
}

func getTextFromFile(filename string) string {

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(bytes)
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
