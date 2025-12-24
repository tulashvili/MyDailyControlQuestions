package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func askQuestions() []Entry {
	var answers = []Entry{}

	for _, question := range QUESTIONS {
		fmt.Println("Категория:", question.Category)
		fmt.Println("Вопрос:", question.Text)

		var answer int
		fmt.Scanln(&answer)

		answers = append(answers, Entry{
			Question: question,
			Answer:   Answer{answer},
		})
	}
	return answers
}

func saveAnswers(answers []Entry) {
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
