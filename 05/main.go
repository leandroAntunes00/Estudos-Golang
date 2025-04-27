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
	var meuArray [3]int
	meuArray[0] = 11
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(meuArray[len(meuArray)-1])

	for i, v:= range meuArray {
		fmt.Printf("o valor do indice é %d e o valor é %d\n", i, v)
	}
}
