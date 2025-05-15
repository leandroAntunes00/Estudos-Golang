package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started")
	defer log.Println("Request ended")
	select {
	case <-time.After(5 * time.Second):
		//print stdout
		log.Println("Request processada com sucesso")
		// imprime browser
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		//print stdout
		log.Println("Request cancelada pelo cliente")
		// imprime browser
		http.Error(w, "Request cancelada pelo cliente", http.StatusRequestTimeout)

	}
}
