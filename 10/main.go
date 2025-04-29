package main

import (
	"fmt"
)

//closures funçoes anonimas, dentro função tem outra função

func main() {

	total := func() int{
		return sum(4,5,6,7) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int { 
	total :=0
	for _, numero := range numeros{
		total += numero
	}
	return total
}
