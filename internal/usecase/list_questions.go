package usecase

import "github.com/tulashvili/MyDailyControlQuestions/internal/domain"

// Return question list with category
func GetQuestions() []domain.Question {
	return domain.QUESTIONS
}
