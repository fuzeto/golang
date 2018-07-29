package main

import (
	"fmt"
)

func main() {
	i := 1

	var p *int = nil
	p = &i // pegando o endereço de memória da variável
	*p++   // desreferenciando, ou seja, pegando o valor
	i++

	// Go não tem aritmética de ponteiros
	// ex: p++

	fmt.Println(p, *p, i, &i)
}
