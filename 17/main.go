package main

import (
	"fmt"
)

type Cliente struct{
	nome string
}

func (c Cliente) andou() {
	c.nome = "Leandro Antunes"
	fmt.Printf("O cliente %s andou\n", c.nome)
}


func main() {
	leandro := Cliente{
		nome:"leandro",
	}
	leandro.andou()
	fmt.Printf("O valor da struct com nome %v\n", leandro.nome)
}
