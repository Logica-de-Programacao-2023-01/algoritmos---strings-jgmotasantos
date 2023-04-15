package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&s)

	output := removeVowels(s)

	fmt.Printf("A string sem vogais é: %s\n", output)
}

func removeVowels(str string) string {
	vowels := "aeiouAEIOU"
	result := strings.Builder{}

	for _, char := range str {
		if !strings.ContainsRune(vowels, char) {
			result.WriteRune(char)
		}
	}

	return result.String()
}
