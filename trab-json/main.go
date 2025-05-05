package main

import (
	"encoding/json"
	"os"
)


type Conta struct {
	Numero int `json:"n"`
	Saldo int `json:"s"` //`json:"-"`
}

func main(){
	conta := Conta{Numero:1,Saldo:100}
	res,err :=json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	encoder :=json.NewEncoder(os.Stdout)
	encoder.Encode(conta)

	jsonPuro := []byte(`{"n":2,"s":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err!=nil {
		println(err)
	}
	println(contaX.Saldo)
}