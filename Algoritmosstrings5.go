package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite sua string: ")
	scanner.Scan()
	frase := scanner.Text()

	i, err := strconv.ParseFloat(frase, 641)
	if err != nil {

		fmt.Printf("O valor %v não é um valor em ponto flutuante\n", frase)

	} else {

		fmt.Printf("O valor em ponto flutuante é valido \n", i)

	}

}
