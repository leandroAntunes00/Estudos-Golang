package main

import (
	"fmt"
)

type Endereco struct{
	Logradouro string
	Numero int
	Cidade string
	Estado string
}

type Client struct {
	Nome string
	Idade int
	Ativo bool
	Adress Endereco
}


func (c Client) Desativar(){
	c.Ativo = false
	fmt.Printf("O Cliente %s desativou alguma coisa hehe\n", c.Nome)
}

//closures funçoes anonimas, dentro função tem outra função

func main() {
	leandro := Client {
		Nome: "leandro",
		Idade: 30,
		Ativo: true,
	}

	leandro.Adress.Cidade = "Cascavel"
	leandro.Adress.Cidade = "Toledo"
	leandro.Desativar()

	fmt.Printf("%s tem idade %d e é %t mora na cidade de %s\n", leandro.Nome,leandro.Idade,leandro.Ativo, leandro.Adress.Cidade)

}


