package passport

import (
	"context"
	"net/http"
)

type TimeSlotResponse struct {
	Name        string `json:"name"`
	Status      bool   `json:"status"`
	Capacity    int    `json:"capacity"`
	VipCapacity int    `json:"vipCapacity"`
}

func (p PassportAPI) GetTimeSlot(addressId int, date string) (*[]TimeSlotResponse, error) {
	url := p.buildTimeSlotSlug(TimeSlots, addressId, date)

	req, err := p.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res := &[]TimeSlotResponse{}

	if _, err := p.client.Do(context.Background(), req, res); err != nil {
		return nil, err
	}

	return res, nil
}
