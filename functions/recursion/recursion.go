package main

import "fmt"

// uint faz com que seja apenas inteiros
func fact(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(4))
}
