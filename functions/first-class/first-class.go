package main

import (
	"fmt"
)

// função anonima, atribuindo ela à uma variável
var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	// posso tambem criar uma função local (dentro da main)
	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))
}
