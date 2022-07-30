package cron

import (
	"encoding/json"
	"fmt"
	"passport-date/locations"
	"passport-date/passport"
	"passport-date/utils"
)

var location locations.Location

func IsDateAvailable() (bool, error) {
	district, err := location.GetAddressId("Kavrepalanchok")
	if err != nil {
		fmt.Println("err getting district id", err)
		return false, err
	}
	fmt.Println("district", district)
	passport, err := passport.NewPassportAPI()

	if err != nil {
		fmt.Println("err loading passport api", err)
		return false, err
	}

	calendarData, err := passport.GetCalender(district)

	if err != nil {
		fmt.Println("err getting calender date", err)
		return false, err
	}
	js, err := json.MarshalIndent(calendarData, "", "  ")
	fmt.Println("js", string(js))

	allPossibleDates, err := utils.GetDateRange(calendarData.MinDate, calendarData.MaxDate)

	if err != nil {
		fmt.Println("err geting possible dates", err)
		return false, err
	}

	fmt.Println("allPossibleDates", allPossibleDates)

	difference := utils.SetDifference(calendarData.OffDates, allPossibleDates)
	if len(difference) != 0 {
		fmt.Printf("Bingoo!! Date Found")
		return true, nil
	}
	fmt.Printf("No Date Found")
	return false, nil
}
