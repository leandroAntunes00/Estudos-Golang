package main

import "net/http"


func main(){
	//primeiro exemplo
	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, world!"))
	// })
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", &blog{title: "My Blog"})
	http.ListenAndServe(":8080", mux)
	//segundo exemplo
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Leandro!"))
	})
	http.ListenAndServe(":8081", mux2)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

type blog struct {
	title string
}

func (b *blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {	
	w.Write([]byte(b.title))
}