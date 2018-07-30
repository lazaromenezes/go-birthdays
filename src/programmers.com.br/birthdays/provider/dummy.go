package provider

import (
	"programmers.com.br/birthdays/model"
	"time"
)

type DummyBirthdayeeProvider struct{}

func (c DummyBirthdayeeProvider) BirthdayeesForDate(date time.Time) []*model.Birthdayee {
	cadu := model.Birthdayee{
		Name:        "Cadu",
		Email:       "cadu@oscardoso.com.br",
		DateOfBirth: time.Date(1979, time.October, 12, 0, 0, 0, 0, time.UTC),
	}

	carol := model.Birthdayee{
		Name:        "Carol",
		Email:       "carolinacapel@gmail.com",
		DateOfBirth: time.Date(1979, time.July, 28, 0, 0, 0, 0, time.UTC),
	}

	birthdayees := []*model.Birthdayee{&cadu, &carol}
	return birthdayees
}
