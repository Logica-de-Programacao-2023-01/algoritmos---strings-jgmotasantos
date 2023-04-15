package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Print("Escreva uma string: ")
	fmt.Scan(&x)

	if string(x[0]) == strings.ToLower(string(x[0])) {
		contador := 1
		for _, letra := range x {
			if string(letra) == strings.ToUpper(string(letra)) {
				contador++
			}
		}
		fmt.Printf("A string %s possui %d palavras", x, contador)
	} else {
		fmt.Println("Não está em camelCase")
	}
}
