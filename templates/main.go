package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome string
	CargaHoraria int
}

func main(){
	curso := Curso{
		Nome: "Curso de Go",
		CargaHoraria: 40,
	}

	tmp := template.New("cursoTemplate")
	tmp, err := tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	if err != nil {
		panic(err)
	}
	err = tmp.ExecuteTemplate(os.Stdout, "cursoTemplate", curso)
	if err != nil {
		panic(err)
	}
}