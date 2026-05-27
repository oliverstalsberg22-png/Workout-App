package main

import (
	"github.com/go-chi/chi/v5"
	//"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	log.Fatal(http.ListenAndServe("PORT", r))
}
