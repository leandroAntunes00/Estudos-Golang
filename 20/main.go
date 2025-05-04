package main

import (
	"fmt"
	"github.com/leandroAntunes00/Estudos-Golang/20/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	d := matematica.Divisao(10,0)
	fmt.Printf("Resultado Soma: %v\n", s)
	fmt.Printf("Resultado Divisão: %v\n", d)
}
