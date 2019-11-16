package main

import (
	"strconv"
)

//IncreaseText adds spaces at the end of the text until the length of the text is divided by a multiplier
func IncreaseText(text string, multiplier int) string {

	runes := []rune(text)

	desiredLength := len(runes) * 32

	for desiredLength%multiplier != 0 {
		desiredLength += 32
	}

	return PadRight(text, desiredLength/32, " ")
}

//RunesToBits converts slice of runes to binary representation
func RunesToBits(runes []rune) string {
	var str string

	for _, value := range runes {
		tmp := strconv.FormatInt(int64(value), 2)
		str += PadLeft(tmp, 32, "0")
	}

	return str
}

//BitsToRunes ...
func BitsToRunes(bits string) []rune {
	var runes []rune

	for i := 0; i < len(bits); i += 32 {
		tmp, _ := strconv.ParseInt(bits[i:i+32], 2, 32)
		runes = append(runes, rune(tmp))
	}

	return runes
}

//Reverse string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//PadRight ...
func PadRight(text string, desiredLength int, placeholder string) string {

	str := ""

	for i := 0; i < (desiredLength - len(text)); i++ {
		str += placeholder
	}
	return text + str
}

//PadLeft ...
func PadLeft(text string, desiredLength int, placeholder string) string {
	str := ""

	for i := 0; i < (desiredLength - len(text)); i++ {
		str += placeholder
	}
	return str + text
}
