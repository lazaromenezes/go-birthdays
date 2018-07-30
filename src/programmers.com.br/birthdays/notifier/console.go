package notifier

import (
	"fmt"
	"programmers.com.br/birthdays/model"
)

type ConsoleNotifier struct{}

func (n ConsoleNotifier) NotifyAbout(birthdayee *model.Birthdayee) {
	fmt.Printf("Sending birthday notification about %s.\n", birthdayee.Name)
}
