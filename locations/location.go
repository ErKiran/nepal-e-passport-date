package locations

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type Location struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Enabled     bool   `json:"enabled"`
	Description string `json:"description"`
}

func (l Location) Load() ([]Location, error) {
	locationFile, err := os.Open("locations/data/locations.json")

	if err != nil {
		return nil, err
	}

	defer locationFile.Close()

	byteValue, err := ioutil.ReadAll(locationFile)

	if err != nil {
		return nil, err
	}

	var possibleLocations []Location
	json.Unmarshal(byteValue, &possibleLocations)

	return possibleLocations, nil
}

func (l Location) Names() ([]string, error) {
	locations, err := l.Load()

	if err != nil {
		return nil, err
	}

	var names []string

	for _, location := range locations {
		names = append(names, location.Name)
	}

	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})

	return names, nil
}

func (l Location) FuzzySearch(address string) (string, error) {
	locations, err := l.Load()
	if err != nil {
		return "", nil
	}
	for _, location := range locations {
		if strings.Contains(strings.ToLower(location.Address), strings.ToLower(address)) {
			return location.Address, nil
		}
	}
	return "", nil
}

func (l Location) GetAddressId(address string) (int, error) {
	locations, err := l.Load()
	if err != nil {
		return 0, err
	}
	for _, location := range locations {
		if location.Name == address {
			return location.ID, nil
		}
	}
	return 0, nil
}
