package main

import "fmt"

func main() {
	salarios := map[string]float64{
		"leandro": 1000.00,
		"andre":   2000.00,
		"andreia": 3000.00,
	}
	delete(salarios, "andreia")
	salarios["fernanda"] = 4000.00



	fmt.Println(salarios["andre"])
	fmt.Println(salarios["andreia"])
	fmt.Println(salarios["leandro"])
	fmt.Println(salarios["fernanda"])

	// sal := make(map[string]float64, 10)
	// sal1 = map[string]float64{}

	for k, v := range salarios {
		fmt.Printf("O valor do %s é %.2f\n", k, v)
	}
}