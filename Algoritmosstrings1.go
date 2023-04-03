package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite sua string com letras minusculas: ")
	scanner.Scan()
	frase := scanner.Text()
	caixaAlta := strings.ToUpper(frase)
	fmt.Println("Sua string em caixa alta é: ", caixaAlta)
}
