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

func main() {
	question := Question{
		Section: "Биология",
		Title:   "Проснулся(-ась) в одно и то же время, чувствуя себя отдохнувшим? (1-5)",
		Answer:  0,
	}

	sections := []string{"Биология"}

	biologyQuestions := map[int]string{
		1: "Проснулся(-ась) в одно и то же время, чувствуя себя отдохнувшим? (1-5)",
		2: "Избегал простых сахаров и обработанной еды? (1- в питании было много сахара, 5 - сахара, обработанной еды  вообще не было)",
		3: "В рационе было достаточно белка и клетчатки? (1 - не было вообще, 5 - да)клетчатка содержится в фруктах, овощах, бобовых",
	}

	var answer int
	var numQuestion int = 1
	var numSection int = 0

	for {
		question.Section = sections[numSection]
		question.Title = biologyQuestions[numQuestion]

		fmt.Printf("Секция: %s\nВопрос: %s\n", question.Section, question.Title)
		fmt.Printf("Ответ:  ")
		fmt.Scanln(&answer)
		question.Answer = answer
		fmt.Println("----------------")

		if numQuestion == len(biologyQuestions) {
			break
		} else {
			numQuestion++
			continue
		}

	}

}
