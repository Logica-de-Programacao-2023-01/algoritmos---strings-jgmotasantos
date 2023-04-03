package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Escreva um programa que receba uma string e substitua todas as ocorrências desse caractere na string por outro caractere especificado pelo usuário. Informe ao usuário o resultado.
	var letra, letrasubstituir string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite sua string: ")
	scanner.Scan()
	frase := scanner.Text()
	fmt.Print("Agora,digite a letra que deseja substituir: ")
	fmt.Scan(&letra)
	fmt.Print("Agora, digite a lera que voce deseja colocar no lugar da substituida: ")
	fmt.Scan(&letrasubstituir)
	espaçovazio := strings.ReplaceAll(frase, letra, letrasubstituir)
	fmt.Print("Sua string sem espaços vazios fiacaria: ", espaçovazio)

}
