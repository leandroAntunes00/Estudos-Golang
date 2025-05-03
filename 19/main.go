package main

type MyNumber int


type Number interface {
	~int | ~float64
}

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaInterfaceVazia(m map[string]interface{}) float64 {
	var soma float64
	for _, v := range m {
		switch value := v.(type) {
		case int:
			soma += float64(value)
		case float64:
			soma += value
		case int32:
			soma += float64(value)
		case int64:
			soma += float64(value)
		case float32:
			soma += float64(value)
		}
	}
	return soma
}

// SomaGenerics sums the values of a map using generics.
// It accepts a map with string keys and values of type T,
// where T is a type that can be used in numeric operations.
func SomaGenerics[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}


func Compara[T comparable] (a T, b T) bool{
	if a==b{
		return true
	}
	return false
}

func main() {
	mInt := map[string]int{
		"item1": 10,
		"item2": 20,
		"item3": 30,
		"item4": 40,
	}

	mMyInt := map[string]MyNumber{
		"item1": 10,
		"item2": 20,
		"item3": 30,
		"item4": 40,
	}

	mFloat := map[string]float64{
		"item1": 10.5,
		"item2": 20.5,
		"item3": 30.5,
		"item4": 40.5,
	}
	mTypeAssertion := map[string]interface{}{
		"item1": 10,
		"item2": 20.5,
		"item3": int64(30),
		"item4": float32(40.5),
	}

	//

	mGenerics := map[string]int{
		"item1": 10,
		"item2": 20,
		"item3": 30,
		"item4": 40,
	}

	mGenericsFloat := map[string]float64{
		"item1": 10.5,
		"item2": 20.5,
		"item3": 30.5,
		"item4": 40.5,
	}

	println("Soma generics int:", SomaGenerics(mGenerics))
	println("Soma generics float:", SomaGenerics(mGenericsFloat))
	println("Soma generics MyNumber:", int(SomaGenerics(mMyInt)))

	println("Soma interface:", SomaInterfaceVazia(mTypeAssertion))
	println("Soma int:", SomaInteiro(mInt))
	println("Soma float:", SomaFloat(mFloat))
	println("Soma typeAssertion:", SomaInterfaceVazia(mTypeAssertion))

	println(Compara(10,10.8))
}
