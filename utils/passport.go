package utils

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type PassportAPI struct {
	client *Client
}

const (
	Calenders = "calendars"
)

type CalenderResponse struct {
	MinDate              string   `json:"minDate"`
	MaxDate              string   `json:"maxDate"`
	OffDates             []string `json:"offDates"`
	WeeklyOffDaysIndexes []int    `json:"weeklyOffDaysIndexes"`
}

func NewPassportAPI() (*PassportAPI, error) {
	var client *Client
	client = NewClient(nil, os.Getenv("PASSPORTURL"))

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

func (p PassportAPI) GetCalender(addressId int) (*CalenderResponse, error) {
	url := p.buildCalendarSlug(Calenders, addressId)

	req, err := p.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res := &CalenderResponse{}

	if _, err := p.client.Do(context.Background(), req, res); err != nil {
		return nil, err
	}

	return res, nil
}
