package calendar

import (
	"fmt"
	"time"
)

func GenerateCalendar() {
	start := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local)
	end := start.AddDate(1, 0, 0)

	days := make([]time.Time, 366)
	next := start
	i := 0
	for end.Compare(next) == 1 {
		days[i] = next
		next = next.AddDate(0, 0, 1)
		i++
	}

	GenerateWeekCalendar(start)
	GenerateWeekCalendar(end)
}

func GenerateWeekCalendar(date time.Time) {
	weekDays := []string{
		"Lun",
		"Mar",
		"Mer",
		"Jeu",
		"Ven",
		"Sam",
		"Dim",
	}

	for i, weekDay := range weekDays {
		fmt.Printf("%-20s", weekDay)
		if i == len(weekDays)-1 {
			fmt.Print("\n")
		} else {
			fmt.Printf(" | ")
		}
	}

	startDay := date.Weekday()
	daysSinceMonday := int(startDay) - int(time.Monday)
	start := date.AddDate(0, 0, -daysSinceMonday)
	daysInWeek := 7
	for i := range daysInWeek {
		day := start.AddDate(0, 0, i)
		date := fmt.Sprintf("%d-%s-%d", day.Year(), day.Month(), day.Day())
		fmt.Printf("%-20s", date)
		if i != daysInWeek-1 {
			fmt.Printf(" | ")
		} else {
			fmt.Printf("\n")
		}
	}
}
