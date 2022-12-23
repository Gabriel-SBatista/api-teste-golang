package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Gabriel-SBatista/API-POSTGRESQL/models"
)

func CreateMorador(w http.ResponseWriter, r *http.Request) {

	var morador models.Morador

	err := json.NewDecoder(r.Body).Decode(&morador)
	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.InsertMorador(morador)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Morador inserida com sucesso! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func CreateVisitante(w http.ResponseWriter, r *http.Request) {

	var visitante models.Visitante

	err := json.NewDecoder(r.Body).Decode(&visitante)
	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.InsertVisitante(visitante)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Error ao tentar inserir um visitante: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Visitante inserido com sucesso! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	return
}
