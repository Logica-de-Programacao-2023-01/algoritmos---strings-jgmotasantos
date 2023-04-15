package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)

	freqMap := make(map[rune]int)
	for _, char := range s {
		freqMap[char]++
	}

	var letras []string
	for _, char := range s {
		if freqMap[char] == 1 {
			letras = append(letras, string(char))
		}
	}
	fmt.Println("Letras únicas:", strings.Join(letras, ""))
}
