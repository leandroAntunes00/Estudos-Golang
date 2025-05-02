package main

// import (
// 	"fmt"
// )

func main() {
	// memoria -> endereço -> valor
	//* é 0xc000054738 , endereçamento
	// variavel -> ponteiro que tem um endereço na memoria -> valor

	a := 10
	var ponteiro *int = &a

	*ponteiro = 20 + 2
	b := &a
	println(*b + 1)

}
