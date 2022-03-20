package passport

import (
	"fmt"
	"os"

	"passport-date/utils"
)

type PassportAPI struct {
	client *utils.Client
}

const (
	Calenders = "iups-api/calendars"
	TimeSlots = "iups-api/timeslots"
)

func NewPassportAPI() (*PassportAPI, error) {
	var client *utils.Client
	client = utils.NewClient(nil, os.Getenv("PASSPORTURL"))

	newPassportAPI := &PassportAPI{
		client: client,
	}
	return newPassportAPI, nil
}

func (p PassportAPI) buildCalendarSlug(url string, addressId int) string {
	return fmt.Sprintf("%s/%d/false", url, addressId)
}

func (p PassportAPI) buildTimeSlotSlug(url string, addressId int, date string) string {
	return fmt.Sprintf("%s/%d/%s/false", url, addressId, date)
}
