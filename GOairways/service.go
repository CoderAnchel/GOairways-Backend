package main

import "encoding/json"

func parseJSON(jsonData []byte) (*RouteCreator, error) {
	var routeCreator RouteCreator
	err := json.Unmarshal(jsonData, &routeCreator)
	if err != nil {
		return nil, err
	}
	return &routeCreator, nil
}

func isSelected(day string, selectedDays []string) int {
	for _, d := range selectedDays {
		if d == day {
			return 1
		}
	}
	return 0
}
