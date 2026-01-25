package usecase

import "fmt"

//  repository должен не напрямую обращаться к sqlite_repository.go, 
// а обращаться должен к методам answer_repository.go , 
// которые реализовывается в sqlite_repository.go

func GetStatAnswers() {
	rows, err := repository.SelectRows(db, period)
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		formatedPrintResult(row)
	}
}

func formatedPrintResult(userAnswer repository.UserAnswerRow) {
	fmt.Println("---------------------")
	fmt.Println("ID:", userAnswer.ID)
	fmt.Println("Дата:", userAnswer.AnsweredAt)
	fmt.Println("Категория:", userAnswer.Category)
	fmt.Println("Вопрос:", userAnswer.Question)
	fmt.Println("Ответ:", userAnswer.Answer)
}
