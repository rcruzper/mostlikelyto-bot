package main

import (
	"mostlikelyto/src/infrastructure/persistence"
)

func main() {
	db := persistence.CreateDBConnection()
	defer db.Close()

	persistence.MigrateDatabase(db)

}
