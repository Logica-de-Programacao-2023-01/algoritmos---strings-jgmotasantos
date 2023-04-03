package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite sua string: ")
	scanner.Scan()
	frase := scanner.Text()
	fmt.Print("Digite sua letra de ocorrencia string: ")
	scanner.Scan()
	letra := scanner.Text()

	count := strings.Count(frase, letra)
	fmt.Printf("o caractere %c aparece %d vezes na string %c", letra, count, frase)
}
