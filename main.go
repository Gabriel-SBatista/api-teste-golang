package main

import (
	"fmt"
	"net/http"

	"github.com/Gabriel-SBatista/API-POSTGRESQL/configs"
	"github.com/Gabriel-SBatista/API-POSTGRESQL/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/morador", handlers.CreateMorador)
	r.Get("/morador", handlers.ListMorador)
	r.Get("/morador/ap/{ap}", handlers.GetMoradorAp)
	r.Get("/morador/id/{id}", handlers.GetMoradorId)
	r.Get("/morador/cpf/{cpf}", handlers.GetMoradorCpf)
	r.Delete("/morador/{id}", handlers.DeleteMorador)
	r.Post("/visitante", handlers.CreateVisitante)
	r.Get("/visitante", handlers.ListVisitante)
	r.Delete("/visitante/{id}", handlers.DeleteVisitante)
	r.Put("/{id}", handlers.UpdateEmCasa)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
