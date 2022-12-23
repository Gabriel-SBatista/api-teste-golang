package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Gabriel-SBatista/API-POSTGRESQL/models"
	"github.com/go-chi/chi/v5"
)

func DeleteMorador(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.DeleteMorador(int64(id))
	if err != nil {
		log.Printf("Erro ao deletar o registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram deletados %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Dados removidos com sucesso",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func DeleteVisitante(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer o parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.DeleteVisitante(int64(id))
	if err != nil {
		log.Printf("Erro ao tentar deletar um visitante: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram deletados: %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Dados removidos com sucesso",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
