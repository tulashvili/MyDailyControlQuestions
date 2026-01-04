package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tulashvili/MyDailyControlQuestions/internal/model"
	"github.com/tulashvili/MyDailyControlQuestions/internal/storage/sqlite3"
)

const answersFile = "qa.json"

// func askQuestions() []model.Entry {
// 	var answers = []model.Entry{}

// 	for _, question := range model.QUESTIONS {
// 		fmt.Println("Категория:", question.Category)
// 		fmt.Println("Вопрос:", question.Text)
// 		var answer int
// 		fmt.Scanln(&answer)

// 		if answer > 5 {
// 			for answer < 1 || answer > 5 {
// 				fmt.Println("Ответ должен быть в диапазоне от 1 до 5")
// 				fmt.Println("Введи значение заново")
// 				fmt.Scanln(&answer)
// 			}
// 		} else {
// 			answers = append(answers, model.Entry{
// 				Question: question,
// 				Answer:   model.Answer{answer},
// 			})
// 		}
// 	}
// 	return answers
// }

func NewDay(t time.Time) model.Day {
	y, m, d := t.Date()
	return model.Day{
		Date: time.Date(y, m, d, 0, 0, 0, 0, t.Location()),
	}
}

// func checkAnswersDay() {
// 	data, err := os.ReadFile(answersFile)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var dailyLog model.DailyLog
// 	err = json.Unmarshal(data, &dailyLog)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// fmt.Println(dailyLog.Date, dailyLog.Entries)
// 	// Output
// 	// {2025-12-30 00:00:00 +0300 MSK} [{{Психология Насколько высокий уровень стресса сегодня? (1 - очень стрессовый день, 5 - стресса почти не было)} {4}}]
// }

func saveAnswers(answers []model.Entry) {
	var date = NewDay(time.Now())

	var dailyLog = model.DailyLog{
		Date:    date,
		Entries: answers,
	}

	data, err := json.MarshalIndent(dailyLog, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(answersFile, data, 0644)
	if err != nil {
		panic(err)
	}

}

func main() {
	// answers := askQuestions()
	// checkAnswersDay()
	pathToDB := sqlite3.SqliteConf{PathToDB: "sqlite3.db"}
	conn, err := sqlite3.OpenConnection(pathToDB)
	if err != nil {
		log.Fatal(err)
	}

	if err := sqlite3.CreateTable(conn); err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	data := model.Entry{
		QuestionCategory: "Bio",
		QuestionText:     "Как ты себя сегодня чувствуешь?",
		Answer:           3,
		CompletedAt:      &now,
	}
	if err := sqlite3.InsertRows(conn, data); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✅ Данные успешно добавлены в таблицу daily_log")
	// saveAnswers(answers)

}
