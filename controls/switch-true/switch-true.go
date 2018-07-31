package main

import (
	"fmt"
	"time"
)

func hour(t int) string {
	switch { // quando não tem condição, automaticamente procura um case que retorna true
	case t < 12:
		return "Bom dia!!!"
	case t > 12 && t < 18:
		return "Boa tarde!!!"
	case t > 18:
		return "Boa noite!!!"
	default:
		return "Valor inválido!"
	}
}

func main() {
	t := time.Now()

	fmt.Println(hour(t.Hour()))
}
