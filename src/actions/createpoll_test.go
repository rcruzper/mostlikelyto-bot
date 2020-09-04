package actions_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"mostlikelyto/src/actions"
	"mostlikelyto/src/domain"
	"mostlikelyto/src/domain/mocks"
	"testing"
)

func TestCreatePoll(t *testing.T) {

	t.Run("should fail if repository returns error", func(t *testing.T) {
		questionRepository := new(mocks.QuestionRepository)
		questionRepository.On("FindAll").Return([]domain.Question{}, errors.New("error"))

		createPoll := actions.NewCreatePoll(questionRepository)

		poll, err := createPoll.Execute()
		assert.Empty(t, poll)
		assert.Equal(t, err, errors.New("error"))
	})

	t.Run("should create a poll", func(t *testing.T) {
		questionRepository := new(mocks.QuestionRepository)
		questions := make([]domain.Question, 0)
		questions = append(questions, domain.Question{Title: "question1"})
		questions = append(questions, domain.Question{Title: "question2"})

		questionRepository.On("FindAll").Return(questions, nil)

		createPoll := actions.NewCreatePoll(questionRepository)

		poll, _ := createPoll.Execute()

		assert.NotEmpty(t, poll.Question)
	})

}
