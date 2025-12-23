package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func askQuestions() []Answer {
	var answers = []Answer{}

	for _, question := range QUESTIONS {
		fmt.Println("Категория:", question.Category)
		fmt.Println("Вопрос:", question.Text)

		var answer int
		fmt.Scanln(&answer)

		answers = append(answers, Answer{
			Question: question,
			Scale:    answer,
		})
	}
	return answers
}

func saveAnswers(answers []Answer) {
	data, err := json.Marshal(answers)
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
