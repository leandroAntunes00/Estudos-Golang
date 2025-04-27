package main

import "fmt"

const frase = "Hello"

type ID int // criação tipos

var (
	b bool    = true
	c int     = 10
	d string  = "leandro"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de frase é %T", f)
}
