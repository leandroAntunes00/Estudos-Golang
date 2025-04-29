package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(300,10)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum(a, b int) (int, error) { //se a e b for mesmo tipo podemos simplificar
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
