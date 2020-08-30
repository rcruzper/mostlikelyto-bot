package main

import (
	"mostlikelyto/src/infrastructure/persistence"
	mysqlQuestionRepository "mostlikelyto/src/infrastructure/repository/question/mysql"
)

func main() {
	db := persistence.CreateDBConnection()
	defer db.Close()

	persistence.MigrateDatabase(db)

	_ = mysqlQuestionRepository.NewMysqlQuestionRepository(db)
}
