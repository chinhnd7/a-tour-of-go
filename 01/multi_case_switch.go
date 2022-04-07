package main

import "fmt"

func main() {
	months := []string{"January", "February", "March", "April", "May",
		"June", "July", "August", "September", "October",
		"November", "December"}

	for i, month := range months {
		fmt.Println(i+1, month, findQuarterYear(month))
	}
}

func findQuarterYear(month string) string {
	switch month {
	case "January", "February", "March":
		return "1st Quarter"
	case "April", "May", "June":
		return "2nd Quarter"
	case "July", "August", "September":
		return "3rd Quarter"
	default:
		return "4th Quarter"
	}
}
