package main

import "fmt"

type Question struct {

	// у вопроса есть принадлежность к одной из секций
	// пока пусть вопрос будет принадлежать к какой-то одной секции, например, Биология
	Section string

	// сам вопрос
	Title string

	// у вопроса есть какой-то по 5-ти бальной шкале,
	// который пользователь должен вводить самостоятельно через stdin
	Answer int
}

// func (q *Question) askQuestion() {
// }

func (q *Question) inputAnswer() {
	var input int
	fmt.Scanln("Введите ответ на вопрос:", &input)
	input = q.Answer
}

func main() {
	question := Question{
		Section: "Биология",
		Title:   "Проснулся(-ась) в одно и то же время, чувствуя себя отдохнувшим? (1-5)",
		Answer:  0,
	}
	question.inputAnswer()
	fmt.Printf("Секция: %s\nВопрос: %s\nОтвет: %d\n", question.Section, question.Title, question.Answer)
}
