package models

import (
	"database/sql"
	"go-echo-mysql/db"
	"go-echo-mysql/helpers"
	"fmt"
)

type User struct {
	Id       int
	Username string
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	conn := db.CreateCon()

	sqlStatement := "select * from users where username = ?"

	err := conn.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Invalid password")
		return false, err
	}

	return true, nil
}
