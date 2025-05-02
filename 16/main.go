package main

import "fmt"

// import (
// 	"fmt"
// )

type x interface {} //interface vazia implementa tudo

func main() {
	var x interface{} = 10
	var y interface{} = "hello world"

	showType(x)
	showType(y)

}

func showType(t interface{}){
	fmt.Printf("O tipo da variave é %T e o valor é %v\n", t, t)
}