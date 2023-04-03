package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Escreva um programa que receba uma string e remova todos os espaços em branco. Informe ao usuário o resultado.
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite sua string utilizando a barra de espaço: ")
	scanner.Scan()
	frase := scanner.Text()
	espaçovazio := strings.ReplaceAll(frase, " ", "")
	fmt.Print("Sua string sem espaços vazios fiacaria: ", espaçovazio)

}
