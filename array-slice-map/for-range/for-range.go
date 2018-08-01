package main

import (
	"fmt"
)

func main() {
	numbers := [...]int{1, 2, 3, 4, 5} // o compilador que vai contar o size do array

	for i, num := range numbers {
		fmt.Printf("%d) %d\n", i, num)
	}

	for _, num := range numbers { // O '_' ignora o index do array, considerando apenas o valor (number)
		fmt.Println(num)
	}
}
