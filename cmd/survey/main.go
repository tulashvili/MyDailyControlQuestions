package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tulashvili/MyDailyControlQuestions/internal/domain"
)

func askQuestions() []domain.Entry {
	var answers = []domain.Entry{}

	for _, question := range domain.QUESTIONS {
		fmt.Println("Категория:", question.Category)
		fmt.Println("Вопрос:", question.Text)

		var answer int
		fmt.Scanln(&answer)

		answers = append(answers, domain.Entry{
			Question: question,
			Answer:   domain.Answer{answer},
		})
	}
	return answers
}

func saveAnswers(answers []domain.Entry) {
	data, err := json.MarshalIndent(answers, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("answers_out.json", data, 0644)
	if err != nil {
		panic(err)
	}

}
func main() {
	answers := askQuestions()
	saveAnswers(answers)
}
