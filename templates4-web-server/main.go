package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Cursos{
			{Nome: "Curso de Javascript", CargaHoraria: 40},
			{Nome: "Curso de Go", CargaHoraria: 40},
			{Nome: "Curso de Python", CargaHoraria: 30},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":8080", nil)

	// Output: Curso: Curso de Go - Carga Horária: 40
	// Curso: Curso de Javascript - Carga Horária: 40
	// Curso: Curso de Python - Carga Horária: 30
}
