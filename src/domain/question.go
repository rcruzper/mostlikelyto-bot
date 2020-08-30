package domain

type Question struct {
	Title string `json:"title" validate:"required"`
}

type QuestionRepository interface {
	FindAll() ([]Question, error)
}
