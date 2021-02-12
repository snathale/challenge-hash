package entity

import (
	"time"

	"github.com/snathale/challenge-hash/calculator/src/utc_time"
)

type User struct {
	Id        string    `json:"_key,omitempty"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Birthday  time.Time `json:"date_of_birth"`
}

func NewUser(firstName, lastName string, birthday time.Time) User {
	return User{
		Id:        "",
		FirstName: firstName,
		LastName:  lastName,
		Birthday:  birthday,
	}
}

func (u *User) isBirthday() bool {
	today := utc_time.Now()
	return today.Day() == u.Birthday.Day() && today.Month() == u.Birthday.Month()
}
