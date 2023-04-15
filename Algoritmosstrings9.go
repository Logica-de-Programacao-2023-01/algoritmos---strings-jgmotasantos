package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, a, b string
	fmt.Print("Escreva uma string: ")
	fmt.Scan(&x)
	fmt.Print("Escreva uma letra para ser substituida: ")
	fmt.Scan(&a)
	fmt.Print("Escreva uma letra para substituir: ")
	fmt.Scan(&b)

	resultado := strings.ReplaceAll(x, a, b)
	fmt.Println(resultado)
}
