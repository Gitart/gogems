package utils

import (
	"strconv"
	"strings"
)

// TimeToMinute
// ConvertTime (string) to Minutes int64
// timeStr := "09:30"
// Replace this with your input time in "hh:mm" format
func TimeToMinute(timeStr string) float64 {
	parts := strings.Split(timeStr, ":")
	hours, _ := strconv.Atoi(parts[0])
	minutes, _ := strconv.Atoi(parts[1])
	totalMinutes := hours*60 + minutes
	// fmt.Printf("Total minutes: %d\n", totalMinutes)
	return float64(totalMinutes)
}

// TimeToHours
// timeStr := "09:30"
// Replace this with your input time in "hh:mm" format
func TimeToHours(timeStr string) float64 {

	parts := strings.Split(timeStr, ":")
	hours, _ := strconv.Atoi(parts[0])
	minutes, _ := strconv.Atoi(parts[1])

	totalHours := float64(hours) + float64(minutes)/60
	//fmt.Printf("Total hours: %.2f\n", totalHours)
	return float64(totalHours)
}
