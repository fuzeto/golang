package main

import (
	"fmt"
)

func main() {
	array1 := [3]int{1, 2, 3}
	// slice1 := []int{}
	var slice1 []int

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) // o copy só trabalha com slice, nao posso fazer copy de um array
	fmt.Println(slice2)
}
