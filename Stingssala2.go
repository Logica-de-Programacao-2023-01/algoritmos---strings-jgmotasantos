package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite sua string: ")
	scanner.Scan()
	frase := scanner.Text()
	fmt.Println("Sua string contem este numero de caracteres: ", len(frase))
}
