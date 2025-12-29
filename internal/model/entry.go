package model

type Entry struct {
	Question Question `json:"Question"`
	Answer   Answer   `json:"Answer"`
}
