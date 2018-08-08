package main

import (
	"fmt"
)

func inc1(n int) {
	n++
}

// um ponteiro é representado por *
func inc2(n *int) {
	// * é usado para desreferenciar, ou seja
	// ter acesso ao valor que o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // passagem por valor
	fmt.Println(n)

	inc2(&n) // passagem por parametro, com o &, pois ele referencia o endereço para onde o ponteiro aponta
	fmt.Println(n)
}
