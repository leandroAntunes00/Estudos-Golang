package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(1,4,6,78,8,5,4))
}

func sum(numeros ...int) int { 
	total :=0
	for _, numero := range numeros{
		total += numero
	}
	return total
}
