package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var x string
	fmt.Print("Escreva uma string: ")
	fmt.Scan(&x)

	_, err := strconv.Atoi(x)
	if err != nil {
		fmt.Printf("A string %s não é uma sequencia numérica", x)
	} else {
		numeros := strings.Split(x, "")
		decrescente := true
		for i := 1; i < len(numeros); i++ {
			anterior := numeros[i-1]
			atual := numeros[i]
			if anterior <= atual {
				decrescente = false
				break
			}
		}
		if decrescente {
			fmt.Println("Está em ordem decrescente")
		} else {
			fmt.Println("Não está em ordem decrescente")
		}
	}
}
