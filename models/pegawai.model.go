package models

import (
	"go-echo-mysql/db"
	"net/http"
)

type Pegawai struct {
	Id      int
	Nama    string
	Alamat  string
	No_telp string
}

func FetchAllPegawai() (Response, error) {
	var obj Pegawai
	var arrobj []Pegawai
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM pegawai"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.No_telp)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = arrobj

	return res, nil
}

func StorePegawai(nama string, alamat string, no_telp string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT pegawai (nama, alamat, no_telp) values (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, no_telp)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "success"
	res.Data = map[string]int{
		"last_inserted_id": int(lastInsertedId),
	}

	return res, nil

}

func UpdatePegawai(Id int, Nama string, Alamat string, No_telp string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "update pegawai set nama = ?, alamat = ?, no_telp = ? where id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Nama, Alamat, No_telp, Id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]int{
		"rows_affected": int(rowsAffected),
	}

	return res, nil
}

func DeletePegawai(Id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE from pegawai where id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusNoContent
	res.Message = "success"
	res.Data = map[string]int{
		"rows_affected": int(rowsAffected),
	}

	return res, nil
}
