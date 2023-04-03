package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite sua string: ")
	scanner.Scan()
	frase := scanner.Text()
	fmt.Println("Sua string contem este numero de caracteres: ", len(frase))
	fmt.Print("Agora, digite um numero n, em que este sera a ordem de caracteres que ficara em caixa alta.Seu n nao pode exceder o numero de caracteres de seu string. ")
	fmt.Scan(&n)
	Alta := (string(frase[:n]))
	resto := (string(frase[n:]))
	caixaAlta := strings.ToUpper(Alta)
	fmt.Print(caixaAlta + resto)

}
