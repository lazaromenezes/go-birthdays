package main

import (
	"programmers.com.br/birthdays"
	"programmers.com.br/birthdays/notifier"
	"programmers.com.br/birthdays/provider"
	"time"
)

func main() {
	notification_engine := birthdays.BirthdayNotificationEngine{
		Provider: provider.DummyBirthdayeeProvider{},
		Notifier: notifier.ConsoleNotifier{},
	}

	notification_engine.SendNotifications(time.Now())
}
