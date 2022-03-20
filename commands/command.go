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

	var interestedDate string
	prompt := &survey.Select{
		Message: "When will you like to see aviable timeslots for the appointment?",
		Options: difference,
	}

	survey.AskOne(prompt, &interestedDate)

	slots, err := passport.GetTimeSlot(addressId, interestedDate)

	if len(*slots) == 0 {
		fmt.Println("Sorry No Dates Avaibale!")
		return
	}

	var availableSlots []string
	for _, slot := range *slots {
		if slot.Capacity != 0 {
			availableSlots = append(availableSlots, slot.Name)
		}
	}

	var interestedSlot string
	timeSlotprompt := &survey.Select{
		Message: "When will you like to see aviable timeslots for the appointment?",
		Options: availableSlots,
	}

	survey.AskOne(timeSlotprompt, &interestedSlot)

	fmt.Println("interestedSlots", interestedSlot)
}
