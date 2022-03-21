package commands

import (
	"fmt"

	"passport-date/locations"
	"passport-date/utils"

	"passport-date/passport"

	"github.com/AlecAivazis/survey/v2"

	"github.com/spf13/cobra"
)

var location locations.Location

var DateCmd = &cobra.Command{
	Use:   "date",
	Short: "Find the possible available Date",
	Long:  `Find the possible available Date If you want to know the available date for a location, you can use this command.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var address string

		if len(args) == 1 {
			matchedAddress, err := location.FuzzySearch(args[0])

			if err != nil {
				fmt.Println("err", err)
				return
			}

			if matchedAddress != "" {
				address = matchedAddress
				GetAddress(address)
				return
			}

			SurveyAddressPrompt(address)
			return
		}
		SurveyAddressPrompt(address)

	},
}

func GetAddress(address string) {
	addressId, err := location.GetAddressId(address)

	if err != nil {
		fmt.Println("err", err)
		return
	}

	passport, err := passport.NewPassportAPI()

	if err != nil {
		fmt.Println("err", err)
		return
	}

	calendarData, err := passport.GetCalender(addressId)

	if err != nil {
		fmt.Println("err", err)
		return
	}

	allPossibleDates, err := utils.GetDateRange(calendarData.MinDate, calendarData.MaxDate)

	if err != nil {
		fmt.Println("err", err)
		return
	}

	difference := utils.SetDifference(calendarData.OffDates, allPossibleDates)

	if len(difference) == 0 {
		fmt.Println("Sorry No Dates Avaibale!")
		return
	}

	interestedDate := SurveyDatePrompt(difference)
	slots, err := passport.GetTimeSlot(addressId, interestedDate)

	if err != nil {
		fmt.Println("err", err)
		return
	}

	SurveyTimeSlotPrompt(slots)

	SurveyOnceMorePrompt(difference, slots)
}

func SurveyAddressPrompt(address string) {
	locations, err := location.Names()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	prompt := &survey.Select{
		Message: "Select a address (or a Place) where you want to set your appointment?",
		Options: locations,
	}

	survey.AskOne(prompt, &address)
	GetAddress(address)
}

func SurveyDatePrompt(difference []string) string {
	var interestedDate string
	prompt := &survey.Select{
		Message: "When will you like to see aviable timeslots for the appointment? 👨‍💻",
		Options: difference,
	}

	survey.AskOne(prompt, &interestedDate)
	return interestedDate
}

func SurveyTimeSlotPrompt(slots *[]passport.TimeSlotResponse) {
	var slot string
	if len(*slots) == 0 {
		fmt.Println("😓😓 Sorry No Dates Avaibale!")
		return
	}

	var availableSlots []string
	for _, slot := range *slots {
		if slot.Capacity != 0 {
			availableSlots = append(availableSlots, slot.Name)
		}
	}

	timeSlotPrompt := &survey.Select{
		Message: "Pick the aviable Timeslot in specific date?",
		Options: availableSlots,
	}

	survey.AskOne(timeSlotPrompt, &slot)
}

func SurveyOnceMorePrompt(difference []string, slots *[]passport.TimeSlotResponse) {
	var onceMore string
	onceMorePrompt := &survey.Select{
		Message: "Do you like to check timeslot for another date?",
		Options: []string{"yes", "no"},
	}

	survey.AskOne(onceMorePrompt, &onceMore)
	if onceMore == "no" {
		return
	}
	SurveyDatePrompt(difference)
	SurveyTimeSlotPrompt(slots)
	SurveyOnceMorePrompt(difference, slots)
}
