package main

const frase = "Hello"
type ID int // criação tipos

var (
	b bool = true
	c int = 10
	d string = "Wesley"
	e float64 = 1.2
	f ID = 1
)

func main() {
	a := "X" //string somente quando a variavel é declarada pela primeira vez
	a = "Y"
	println(a)
	println(frase)
	println(f)
}


