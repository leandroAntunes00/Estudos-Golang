package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": ToUpper})
		t = template.Must(t.ParseFiles(templates...))

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
