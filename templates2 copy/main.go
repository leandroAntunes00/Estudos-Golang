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

	t := template.Must(template.New("cursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
	// Output: Curso: Curso de Go - Carga Horária: 40
}