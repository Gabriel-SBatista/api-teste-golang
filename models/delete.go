package models

import "github.com/Gabriel-SBatista/API-POSTGRESQL/db"

func DeleteMorador(id int64) (int64, error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM moradores WHERE id_morador=$1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func DeleteVisitante(id int64) (int64, error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM visitantes WHERE id_visitante=$1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
