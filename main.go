package main

import "fmt"

type Category string

// категории
const (
	Biology    Category = "Биология"
	Psychology Category = "Психология"
)

type Question struct {

	// у вопроса есть принадлежность к одной из секций
	// пока пусть вопрос будет принадлежать к какой-то одной секции, например, Биология
	Category

	// сам вопрос
	Text string
}

type Answer struct {
	// ответ
	Scale int
}

// список вопросов и категорий
var QUESTIONS = []Question{
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
