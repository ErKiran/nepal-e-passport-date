package commands

import (
	"fmt"

	"passport-date/locations"

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
			}

			if matchedAddress == "" {
				SurveyAddressPrompt(address)
			}
		}

		if len(args) == 0 {
			SurveyAddressPrompt(address)
		}
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
	}

	fmt.Println("addressId", addressId)

	fmt.Println(`address`, address)

	passport, err := passport.NewPassportAPI()

	if err != nil {
		fmt.Println("err", err)
	}

	calendarData, err := passport.GetCalender(addressId)

	fmt.Println("calendarData", calendarData)

	slots, _ := passport.GetTimeSlot(addressId, calendarData.OffDates[0])

	fmt.Println("slots", slots)
}
