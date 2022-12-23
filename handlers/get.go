package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Gabriel-SBatista/API-POSTGRESQL/models"
	"github.com/go-chi/chi/v5"
)

func GetMoradorId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	morador, err := models.GetMoradorId(int64(id))
	if err != nil {
		log.Printf("Erro ao buscar morador: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(morador)
}

func GetMoradorAp(w http.ResponseWriter, r *http.Request) {
	ap, err := strconv.Atoi(chi.URLParam(r, "ap"))
	if err != nil {
		log.Printf("Erro ao fazer o parse do ap: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	moradores, err := models.GetMoradorAp(int64(ap))
	if err != nil {
		log.Printf("Erro ao buscar moradores: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(moradores)
}

func GetMoradorCpf(w http.ResponseWriter, r *http.Request) {
	cpf := chi.URLParam(r, "cpf")

	morador, err := models.GetMoradorCpf(cpf)
	if err != nil {
		log.Printf("Erro ao buscar morador: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(morador)
}
