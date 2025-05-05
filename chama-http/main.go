package main

import (
	"io"
	"net/http"
)



func main(){
	request,err := http.Get("https://www.google.com")
	if err !=nil {
		panic(err)
	}
	
	defer request.Body.Close()
	result,err := io.ReadAll(request.Body)
	if err != nil{
		panic(err)
	}
	println(string(result))
	//sempre lembrar de fechar
}