package main

import (
	"fmt"
)

func main() {
	employeeByLetter := map[string]map[string]float64{
		"G": {
			"Gabriel":  1508.41,
			"Giordano": 3151.92,
		},
		"J": {
			"José": 5151.05,
		},
		"M": {
			"Maria": 2415.00,
		},
	}

	delete(employeeByLetter, "M")

	for letra, funcs := range employeeByLetter {
		for nome, salario := range funcs {
			fmt.Printf("(%s) - %s tem o salário de R$ %.2f\n", letra, nome, salario)
		}
	}
}
