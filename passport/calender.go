package passport

import (
	"context"
	"net/http"
)

type CalenderResponse struct {
	MinDate              string   `json:"minDate"`
	MaxDate              string   `json:"maxDate"`
	OffDates             []string `json:"offDates"`
	WeeklyOffDaysIndexes []int    `json:"weeklyOffDaysIndexes"`
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
