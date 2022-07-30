package cron

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func Cron() {
	cron := cron.New()
	cron.AddFunc("@every 1m", func() {
		is, err := IsDateAvailable()
		if err != nil {
			fmt.Println("err checking date", err)
		}
		if is {
			Mailer()
		}
	})
}
