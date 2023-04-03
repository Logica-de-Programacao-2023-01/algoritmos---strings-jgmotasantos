package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite sua string :")
	scanner.Scan()
	frase := scanner.Text()

	palavrasFormadas := strings.Fields(frase)
	numPalavrasFormadas := len(palavrasFormadas)
	fmt.Printf("Sua string possui %d palavras\n", numPalavrasFormadas)

}
