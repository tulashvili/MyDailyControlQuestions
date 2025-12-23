package main

type Storage struct {
	AnswerCategory string `json:"Category"`
	AnswerQuestion string `json:"Question"`
	AnswerUser     []int  `json:"Answer"`
}
