package actions

import (
	"math/rand"
	"mostlikelyto/src/domain"
)

type CreatePoll struct {
	questionRepository domain.QuestionRepository
}

func NewCreatePoll(questionRepository domain.QuestionRepository) *CreatePoll {
	return &CreatePoll{
		questionRepository: questionRepository,
	}
}

func (q CreatePoll) Execute() (domain.Poll, error) {
	questions, err := q.questionRepository.FindAll()

	if err != nil {
		return domain.Poll{}, err
	}

	return domain.Poll{
		Question: questions[rand.Intn(len(questions) - 1)].Title,
	}, nil
}
