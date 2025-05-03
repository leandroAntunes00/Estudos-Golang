package main

import "fmt"

func main() {
	var minhaVar interface{} = "Leandro Antunes"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n",res,ok)
	// Type assertion (codigo legado) atualmente utiliza-se generics
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v", res2)
}
