package notifier

import "programmers.com.br/birthdays/model"

type BirthdayNotifier interface {
	NotifyAbout(birthdayee *model.Birthdayee)
}
