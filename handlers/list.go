package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Gabriel-SBatista/API-POSTGRESQL/models"
)

func ListMorador(w http.ResponseWriter, r *http.Request) {
	moradores, err := models.GetAllMoradores()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(moradores)
}

func ListVisitante(w http.ResponseWriter, r *http.Request) {
	visitantes, err := models.GetAllVisitantes()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(visitantes)
}
