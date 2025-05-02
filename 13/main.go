package main

import (
	"fmt"
)

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Client struct {
	Nome   string
	Idade  int
	Ativo  bool
	Adress Endereco
}

func (c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("O Cliente %s desativou alguma coisa hehe\n", c.Nome)
}

//closures funçoes anonimas, dentro função tem outra função

type Pessoa interface {
	Desativar()
}

type Soma10 interface {
	Adiciona10(valor int) int
}

func Adiciona10(valor int) int {
	return valor + 10
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	leandro := Client{
		Nome:  "leandro",
		Idade: 30,
		Ativo: true,
	}

	leandro.Desativar()

}
