package data

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	EmailAddress string
	UserActive   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Password     string
}

func (u *User) PasswordMatches(plainText string) (bool, error) {

	if u.Password != plainText {
		log.Println(u.Password, plainText)
		return false, errors.New("provided password does not match stored password")
	}

	return true, nil
}
