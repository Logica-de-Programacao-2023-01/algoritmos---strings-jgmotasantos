package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite sua primeira string: ")
	scanner.Scan()
	frase := scanner.Text()
	fmt.Print("Digite sua segunda string: ")
	scanner.Scan()
	frase2 := scanner.Text()
	s := frase + " , " + frase2

	fmt.Print("A concatenação das suas strings é: ", s)

}
