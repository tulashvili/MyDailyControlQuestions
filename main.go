package main

import (
	"fmt"
)

func askQuestionsAndSaveAnswers() {
	var answers = []Answer{}

	for _, question := range QUESTIONS {
		fmt.Println(question)
		fmt.Println("Категория:", question.Category)
		fmt.Println("Вопрос:", question.Text)

		var answer int
		fmt.Scanln(&answer)

		answers = append(answers, Answer{
			Question: question,
			Scale:    answer,
		})
	}

	for i := 0; i < len(answers); i++ {
		fmt.Println("Категория:", answers[i].Question.Category)
		fmt.Println("Вопрос:", answers[i].Question.Text)
		fmt.Println("Ответ:", answers[i].Scale)
		fmt.Println()
	}
}

// func saveAnswer(answer []int) {
// 	var answers = []Answer{
// 		{
// 			Scale: 0,
// 		},
// 	}
// 	storage := Storage{
// 		AnswerCategory: "Category",
// 		AnswerQuestion: "Question",
// 		AnswerUser:     answer,
// 	}

// }
func main() {
	askQuestionsAndSaveAnswers()

	// data, err := json.Marshal(answerFile.AnswerUser)
	// if err != nil {
	// 	panic(err)
	// }

	// err = os.WriteFile("answers_output.json", data, 0644)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("--------------------")
}
