package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite sua string: ")
	scanner.Scan()
	temNumero := false
	palavras := scanner.Text()
	for _, ch := range palavras {
		if unicode.IsDigit(ch) {
			temNumero = true
			break

		}

	}
	if temNumero {
		fmt.Print("Sua string contem um caractere numerico")
	} else {
		fmt.Print("Sua string nao possui um caractere numerico")
	}

}
