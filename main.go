package main

import (
	"fmt"
)

func main() {

	var keyLength int = 0
	fmt.Print("Enter number : ")
	fmt.Scan(&keyLength)

	var privateKey PrivateKey
	fmt.Println("Генерация ключей...")
	privateKey.Generate(keyLength)
	fmt.Printf("Публичный ключ : %v\n", privateKey.GeneratePublicKey())
	fmt.Printf("Личный ключ : %+v\n", privateKey)
	fmt.Println("Ключи успешно сгенерированы")

}
