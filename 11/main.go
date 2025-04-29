package main

import (
	"fmt"
)


type Client struct {
	Nome string
	Idade int
	Ativo bool
}

//closures funçoes anonimas, dentro função tem outra função

func main() {
	leandro := Client {
		Nome: "leandro",
		Idade: 30,
		Ativo: true,
	}

	fmt.Printf("%s tem idade %d e é %t\n", leandro.Nome,leandro.Idade,leandro.Ativo)

}


