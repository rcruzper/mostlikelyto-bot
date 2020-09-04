package repository_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"mostlikelyto/src/domain"
	"mostlikelyto/src/infrastructure/persistence/mysql/repository"
	"testing"
)

func TestMysqlQuestionRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	questionRepository := repository.NewMysqlQuestionRepository(db)
	query := "SELECT title FROM question"

	t.Run("should fail when query fails", func(t *testing.T) {
		// given
		mock.ExpectQuery(query).WillReturnError(errors.New("error"))

		// when
		questions, err := questionRepository.FindAll()

		// then
		assert.Empty(t, questions)
		assert.NotEmpty(t, err)
		assert.Equal(t, err, errors.New("error"))
	})

	t.Run("should be empty", func(t *testing.T) {
		// given
		rows := sqlmock.NewRows([]string{"title"})
		mock.ExpectQuery(query).WillReturnRows(rows)

		// when
		questions, err := questionRepository.FindAll()

		// then
		assert.Empty(t, questions)
		assert.NoError(t, err)
	})

	t.Run("should return questions", func(t *testing.T) {
		// given
		rows := sqlmock.NewRows([]string{"title"}).
			AddRow("question1").
			AddRow("question2")
		mock.ExpectQuery(query).WillReturnRows(rows)

		// when
		questions, err := questionRepository.FindAll()

		// then
		assert.NotEmpty(t, questions)
		assert.NoError(t, err)
		assert.Len(t, questions, 2)
		assert.ElementsMatch(t, questions, []domain.Question{
			{Title: "question1"},
			{Title: "question2"},
		}, "uh")
	})
}
