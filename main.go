package main

import "fmt"

func main() {
	for _, questions := range QUESTIONS {

		fmt.Println("Категория:", questions.Category)
		fmt.Println("Вопрос:", questions.Text)

		// для каждого вопроса мы "создаем" свой ответ
		var initAnswer = Answer{
			Scale: 0,
		}
		fmt.Scanln(&initAnswer.Scale)
		fmt.Println("Ответ:", initAnswer.Scale)
		fmt.Println("--------------------")

	}
}
