package models

import "github.com/Gabriel-SBatista/API-POSTGRESQL/db"

func GetMoradorId(id int64) (morador Morador, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	err = conn.QueryRow(`SELECT * FROM moradores WHERE id_morador=$1`, id).Scan(&morador.ID, &morador.Nome, &morador.CPF, &morador.EmCasa, &morador.IDApartamento)

	return
}

func GetMoradorAp(ap int64) (moradores []Morador, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM moradores WHERE id_apartamento=$1`, ap)
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

func GetMoradorCpf(cpf string) (morador Morador, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	err = conn.QueryRow(`SELECT * FROM moradores WHERE cpf=$1`, cpf).Scan(&morador.ID, &morador.Nome, &morador.CPF, &morador.EmCasa, &morador.IDApartamento)

	return
}
