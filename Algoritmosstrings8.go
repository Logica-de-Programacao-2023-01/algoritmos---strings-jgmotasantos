package main

import "fmt"

func main() {
	var s string
	fmt.Println("Digite uma string:")
	fmt.Scan(&s)

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	output := string(runes)
	fmt.Printf("A string invertida é: %s\n", output)
}
