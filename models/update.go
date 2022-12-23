package models

import "github.com/Gabriel-SBatista/API-POSTGRESQL/db"

func UpdateEmCasa(id int64) (int64, error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	morador, err := GetMoradorId(id)
	if err != nil {
		return 0, err
	}

	sql := `UPDATE moradores SET em_casa=$1 WHERE id_morador=$2`

	if morador.EmCasa == true {
		res, err := conn.Exec(sql, false, id)
		if err != nil {
			return 0, err
		}

		return res.RowsAffected()

	} else {
		res, err := conn.Exec(sql, true, id)
		if err != nil {
			return 0, err
		}

		return res.RowsAffected()
	}
}
