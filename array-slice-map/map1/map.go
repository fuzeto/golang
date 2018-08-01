package main

import (
	"fmt"
)

func main() {
	// não é possível lançar valores dentro de um map vazio
	// var aprovados map[int]string
	aprovados := make(map[int]string)

	aprovados[123] = "Maria"
	aprovados[456] = "José"
	aprovados[789] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[456])
	delete(aprovados, 456)
	fmt.Println(aprovados[456])
}
