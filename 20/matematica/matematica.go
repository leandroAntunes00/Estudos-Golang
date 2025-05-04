package matematica

import "fmt"

//soma visivel somente dentro do pacote Soma todos dora do pacote encherga
//maiuscula exportado, minuscula não exportado
func Soma[T int | float64](a, b T) T {
	return a + b
}

func Subtracao[T int | float64](a, b T) T {
	return a - b
}

func Multiplicacao[T int | float64](a, b T) T {
	return a * b
}

func Divisao[T int | float64](a, b T) T {
	if b == 0 {
		fmt.Println("Erro: Divisão por zero")
		var zero T
		return zero
	}
	return a / b
}