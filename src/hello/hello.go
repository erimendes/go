package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Douglas"
	var versao = 1.1
	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa está na versão: ", versao)

	fmt.Println("O tipo da variável versão é: ", reflect.TypeOf(versao))
}
