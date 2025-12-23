package main

// тот же package, потому на текущий момент реализации логики
// questions не может существовать без main.go

type Category string
type Text string

// категории вопросов
const (
	Biology    Category = "Биология"
	Psychology Category = "Психология"
)

type Question struct {
	Category
	Text
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
	// {
	// 	Category: Biology,
	// 	Text:     "Проснулся(-ась) в одно и то же время, чувствуя себя отдохнувшим? (1-5)",
	// },
	// {
	// 	Category: Biology,
	// 	Text:     "Избегал простых сахаров и обработанной еды? (1 - в питании было много сахара, 5 - сахара не было)",
	// },
	// {
	// 	Category: Biology,
	// 	Text:     "В рационе было достаточно белка и клетчатки? (1 - не было, 5 - да)",
	// },
}
