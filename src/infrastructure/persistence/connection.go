package persistence

import "database/sql"

func CreateDBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost)/mostlikelyto")

	if err != nil {
		panic(err.Error())
	}

	return db
}
