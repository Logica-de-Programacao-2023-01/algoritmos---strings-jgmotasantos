package main

import "fmt"

func main() {
	var s string
	fmt.Print("Digite uma palavra ou frase: ")
	fmt.Scan(&s)

	vogais := "aeiouAEIOU"
	result := ""
	for _, char := range s {
		isvogal := false
		for _, vogal := range vogais {
			if char == vogal {
				isvogal = true
				break
			}
		}
		if isvogal {
			result += "*"
		} else {
			result += string(char)
		}
	}

	fmt.Println("Resultado:", result)
}
