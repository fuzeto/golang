package main

import (
	"fmt"
)

func comprar(work1, work2 bool) (bool, bool, bool) {
	comprarTv50 := work1 == work2
	comprarTv32 := work1 != work2
	comprarSorvete := work1 || work2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := comprar(true, true)

	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
