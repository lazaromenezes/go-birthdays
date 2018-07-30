package provider

import (
	"programmers.com.br/birthdays/model"
	"time"
)

type BirthdayeeProvider interface {
	BirthdayeesForDate(date time.Time) []*model.Birthdayee
}
