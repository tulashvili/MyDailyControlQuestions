package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tulashvili/MyDailyControlQuestions/internal/model"
)

func askQuestions() []model.Entry {
	var answers = []model.Entry{}

	for _, question := range model.QUESTIONS {
		fmt.Println("Категория:", question.Category)
		fmt.Println("Вопрос:", question.Text)

		var answer int
		fmt.Scanln(&answer)

		answers = append(answers, model.Entry{
			Question: question,
			Answer:   model.Answer{answer},
		})
	}
	return answers
}

func saveAnswers(answers []model.Entry) {
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
