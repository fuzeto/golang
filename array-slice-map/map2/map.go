package main

import (
	"fmt"
)

func main() {
	employeeWithSalary := map[string]float64{
		"Jose":     15325.45,
		"Vanessa":  2615.02,
		"Leonardo": 8500.00,
	}

	employeeWithSalary["Pedro José"] = 1200.0
	delete(employeeWithSalary, "Inexistente") //não vai excluir pq não existe, mas tb não dá erro

	for nome, salario := range employeeWithSalary {
		fmt.Println(nome, salario)
	}
}
