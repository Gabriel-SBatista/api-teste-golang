package models

import (
	"errors"

	"github.com/Gabriel-SBatista/API-POSTGRESQL/db"
)

func InsertMorador(morador Morador) (id int64, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO moradores (nome, cpf, id_apartamento) VALUES ($1, $2, $3) RETURNING id_morador`

	err = conn.QueryRow(sql, morador.Nome, morador.CPF, morador.IDApartamento).Scan(&id)

	return
}

func InsertVisitante(visitante Visitante) (id int64, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	moradores, err := GetMoradorAp(visitante.IDApartamento)
	if err != nil {
		return
	}

	for _, v := range moradores {
		if v.EmCasa {
			sql := `INSERT INTO visitantes (nome, cpf, id_apartamento) VALUES ($1, $2, $3) RETURNING id_visitante`

			err = conn.QueryRow(sql, visitante.Nome, visitante.CPF, visitante.IDApartamento).Scan(&id)

			return
		}
	}

	err = errors.New("Não há moradores no momento")
	return
}
