package utils

import (
	"errors"
	"time"
)

func GetDateRange(start, end string) ([]string, error) {
	var days []string
	if start == "" || end == "" {
		return days, errors.New("start and end date are required")
	}

	// Change String to Date
	startDate, err := stringToTime(start)
	if err != nil {
		return days, errors.New("start date is invalid")
	}

	endDate, err := stringToTime(end)
	if err != nil {
		return days, errors.New("end date is invalid")
	}

	// find the difference in days between start and end date
	diffDays := endDate.Sub(startDate).Hours() / 24

	// for loop
	// declare a variable with array of string
	for i := 0; i <= int(diffDays); i++ {
		addedDate := startDate.Add(time.Hour * 24 * time.Duration(i)).Format("2006-01-02")
		if addedDate != endDate.Format("2006-01-02") {
			days = append(days, addedDate)
		}
	}
	days = append(days, end)
	return days, nil
}

func stringToTime(str string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func SetDifference(s1 []string, s2 []string) (difer []string) {
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		// If elements present in the hashmap then don't append differ list.
		if !hash[e] {
			difer = append(difer, e)
		}
	}
	return difer
}
