package main

import "fmt"

func weekDay(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid day"
	}
}

func weekDay2(number int) string {
	var weekday string

	switch {
	case number == 1:
		weekday = "Sunday"
		// fallthrough is used to skip the rest of the cases
	case number == 2:
		weekday = "Monday"
	case number == 3:
		weekday = "Tuesday"
	case number == 4:
		weekday = "Wednesday"
	case number == 5:
		weekday = "Thursday"
	case number == 6:
		weekday = "Friday"
	case number == 7:
		weekday = "Saturday"
	default:
		weekday = "Invalid day"
	}
	return weekday
}

func main() {
	day := weekDay(1)

	fmt.Println(day)

	day = weekDay2(1)
	fmt.Println(day)
}
