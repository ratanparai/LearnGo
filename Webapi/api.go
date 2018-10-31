package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Sample structure to send sample reply

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type EmptyResponse struct {
}

func (a *Account) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Get("/ping", Pong)
	router.Get("/account", sendSingleAccountInfo)
	router.Get("/nocontent", func(w http.ResponseWriter, r *http.Request) {
		render.NoContent(w, r)
	})

	http.ListenAndServe(":3333", router)

}

// Pong return nothing but write Pong! in Response object
func Pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong!"))
	fmt.Println(r)
}

func sendSingleAccountInfo(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &Account{ID: "1", Name: "Ratan"})
}
