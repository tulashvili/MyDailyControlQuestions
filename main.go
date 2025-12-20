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

	biologyQuestions := map[int]string{
		1: "Проснулся(-ась) в одно и то же время, чувствуя себя отдохнувшим? (1-5)",
		2: "Избегал простых сахаров и обработанной еды? (1- в питании было много сахара, 5 - сахара, обработанной еды  вообще не было)",
		3: "В рационе было достаточно белка и клетчатки? (1 - не было вообще, 5 - да)клетчатка содержится в фруктах, овощах, бобовых",
	}

	var answer int
	var countQuestion int = 1

	for {
		fmt.Println(question.Section) // Секция

		question.Title = biologyQuestions[countQuestion]
		fmt.Println(question.Title) // Вопрос

		fmt.Scanln(&answer) // Ответ
		question.Answer = answer

		if countQuestion == len(biologyQuestions) {
			break
		} else {
			countQuestion++
			continue
		}

	}

}
