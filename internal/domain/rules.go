package domain

import "errors"

func ValidateAnswer(answer int) error {
	if answer < 1 || answer > 5 {
		return errors.New("answer must be between 1 and 5")
	}
	return nil
}
