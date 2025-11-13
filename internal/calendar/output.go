package calendar

import (
	"fmt"
	"strings"
	"time"
)

func ToConsole(weekDates [7]time.Time) (string, error) {
	format := "%-20s"
	weekDays := Weekdays()

	buffer := strings.Builder{}
	for i, weekDay := range weekDays {
		translation := "  "
		switch weekDay {
		case time.Monday:
			translation = "Lun"
		case time.Tuesday:
			translation = "Mar"
		case time.Wednesday:
			translation = "Mer"
		case time.Thursday:
			translation = "Jeu"
		case time.Friday:
			translation = "Ven"
		case time.Saturday:
			translation = "Sam"
		case time.Sunday:
			translation = "Dim"
		default:
			panic("Unknown day of the week")
		}

		buffer.WriteString(fmt.Sprintf(format, translation))

		if i != len(weekDays)-1 {
			buffer.WriteString(" | ")
		} else {
			buffer.WriteString("\n")
		}
	}

	for i, weekDate := range weekDates {
		dateStr := fmt.Sprintf("%d-%s-%d", weekDate.Year(), weekDate.Month(), weekDate.Day())

		buffer.WriteString(fmt.Sprintf(format, dateStr))

		if i != len(weekDates)-1 {
			buffer.WriteString(" | ")
		} else {
			buffer.WriteString("\n")
		}
	}

	return buffer.String(), nil
}
