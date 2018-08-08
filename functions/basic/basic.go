package main

import (
	"fmt"
)

// mais de um retorno na função
func manyReturns() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	fmt.Println(manyReturns())
}
