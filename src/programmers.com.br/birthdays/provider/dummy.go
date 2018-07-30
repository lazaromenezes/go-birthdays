package provider

import (
	"programmers.com.br/birthdays/model"
	"time"
)

/* Test Dummy provider should be moved to testing suite at some point. */

type DummyBirthdayeeProvider struct{}

func (c DummyBirthdayeeProvider) BirthdayeesForDate(date time.Time) []*model.Birthdayee {
	cadu := model.Birthdayee{
		Name:        "Cadu",
		Email:       "cadu@email-host.com",
		DateOfBirth: time.Date(1974, time.October, 23, 0, 0, 0, 0, time.UTC),
	}

	badu := model.Birthdayee{
		Name:        "Badu",
		Email:       "badu@email-host.com",
		DateOfBirth: time.Date(1982, time.July, 15, 0, 0, 0, 0, time.UTC),
	}

	birthdayees := []*model.Birthdayee{&cadu, &badu}
	return birthdayees
}
