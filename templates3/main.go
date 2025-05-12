package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome string
	CargaHoraria int
}

type Cursos []Curso

func main(){
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{ Nome: "Curso de Javascript", CargaHoraria: 40},
		{Nome: "Curso de Go", CargaHoraria: 40},
		{Nome: "Curso de Python", CargaHoraria: 30},
	})
	if err != nil {
		panic(err)
	}
	// Output: Curso: Curso de Go - Carga Horária: 40
	// Curso: Curso de Javascript - Carga Horária: 40
	// Curso: Curso de Python - Carga Horária: 30
}