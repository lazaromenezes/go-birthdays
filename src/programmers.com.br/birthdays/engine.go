package birthdays

import (
	"fmt"
	"programmers.com.br/birthdays/notifier"
	"programmers.com.br/birthdays/provider"
	"time"
)

type BirthdayNotificationEngine struct {
	Provider provider.BirthdayeeProvider
	Notifier notifier.BirthdayNotifier
}

func (e BirthdayNotificationEngine) SendNotifications(birth_date time.Time) {
	birthdayees := e.Provider.BirthdayeesForDate(birth_date)

	if len(birthdayees) == 0 {
		fmt.Println("No birthdayees for today!")
	}

	for _, birthdayee := range birthdayees {
		e.Notifier.NotifyAbout(birthdayee)
	}
}
