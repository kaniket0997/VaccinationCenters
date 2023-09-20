package utils

import "fmt"

func ConvertDayToNumber(day string) int64 {
	switch day {
	case "monday":
		return 0
	case "tuesday":
		return 1
	case "wednesday":
		return 2
	case "thursday":
		return 3
	case "friday":
		return 4
	case "saturday":
		return 5
	case "sunday":
		return 6
	}
	fmt.Printf("invalid day, day=%s", day)
	return -1
}
