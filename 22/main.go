package main

import (
	"fmt"
)

func main() {
	//apenas for

	for i := 0; i < 10; i++ {
		//fmt.Println(i)
	}

	// numeros := []string{"um", "dois", "três"}
	// for k, v := range numeros {
	// 	fmt.Println(k, v)
	// }

	indice:=0
	for indice < 8 {
		fmt.Println(indice)
		indice++
	}
	// loop infinito
	// for {
	// 	fmt.Println("Hello")
	// }
}
