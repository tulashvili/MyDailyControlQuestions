package service

import (
	"fmt"
	"time"

	"github.com/tulashvili/MyDailyControlQuestions/internal/models"
)

const (
	Biology    models.Category = "Биология"
	Psychology models.Category = "Психология"
)

var QUESTIONS = []models.Question{
	{
		Category: Psychology,
		Text:     "Насколько высокий уровень стресса сегодня? (1 - очень стрессовый день, 5 - стресса почти не было)",
	},
	{
		Category: Psychology,
		Text:     "Продвинулся ли ты по целям, которые ты себе наметил хотя бы на чуть-чуть?",
	},
	{
		Category: Biology,
		Text:     "Проснулся(-ась) в одно и то же время, чувствуя себя отдохнувшим? (1-5)",
	},
	{
		Category: Biology,
		Text:     "Избегал простых сахаров и обработанной еды? (1 - в питании было много сахара, 5 - сахара не было)",
	},
	{
		Category: Biology,
		Text:     "В рационе было достаточно белка и клетчатки? (1 - не было, 5 - да)",
	},
}

// Return question list with category
func GetQuestions() []models.Question {
	return QUESTIONS
}

// Validate user answer
func ValidateAnswer(answer int) error {
	if answer < 1 || answer > 5 {
		return fmt.Errorf("answer must be between 1 and 5")
	}
	return nil
}

// Init user answer
func CreateUserAnswer(question models.Question, answer int) models.UserAnswer {
	now := time.Now()
	

	return models.UserAnswer{
		QuestionCategory: string(question.Category),
		QuestionText:     question.Text,
		Answer:           answer,
		AnsweredAt:       &now,
	}
}




