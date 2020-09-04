package repository

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"mostlikelyto/src/domain"
)

type mysqlQuestionRepository struct {
	Conn *sql.DB
}

func NewMysqlQuestionRepository(Conn *sql.DB) domain.QuestionRepository {
	return &mysqlQuestionRepository{Conn}
}

func (r *mysqlQuestionRepository) FindAll() (questions []domain.Question, err error) {
	query := "SELECT title FROM question"
	results, err := r.Conn.Query(query)
	if err != nil {
		log.Error("error trying to execute query..", err)
		return nil, err
	}

	questions = make([]domain.Question, 0)
	for results.Next() {
		var question domain.Question
		_ = results.Scan(&question.Title)
		questions = append(questions, question)
	}

	return questions, nil
}
