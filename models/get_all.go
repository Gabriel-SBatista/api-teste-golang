package models

import "github.com/Gabriel-SBatista/API-POSTGRESQL/db"

func GetAllMoradores() (moradores []Morador, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM moradores`)
	if err != nil {
		return
	}

	var morador Morador
	for rows.Next() {
		err = rows.Scan(&morador.ID, &morador.Nome, &morador.CPF, &morador.EmCasa, &morador.IDApartamento)
		if err != nil {
			continue
		}

		moradores = append(moradores, morador)
	}

	return
}

func GetAllVisitantes() (visitantes []Visitante, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM visitantes`)

	var visitante Visitante

	for rows.Next() {
		err = rows.Scan(&visitante.ID, &visitante.Nome, &visitante.CPF, &visitante.IDApartamento)
		if err != nil {
			continue
		}

		visitantes = append(visitantes, visitante)
	}

	return
}
