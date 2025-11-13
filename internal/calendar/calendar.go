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

	str, _ := GenerateWeekCalendar(start, ToConsole)
	fmt.Print(str)
	str2, _ := GenerateWeekCalendar(end, ToConsole)
	fmt.Print(str2)
}

func GenerateWeekCalendar(date time.Time, renderFunc func([7]time.Time) (string, error)) (string, error) {
	datesInWeek := [7]time.Time{}

	startDay := date.Weekday()
	daysSinceMonday := int(startDay) - int(time.Monday)
	start := date.AddDate(0, 0, -daysSinceMonday)
	daysInWeek := 7
	for i := range daysInWeek {
		day := start.AddDate(0, 0, i)
		datesInWeek[i] = day
	}

	return renderFunc(datesInWeek)
}

func Weekdays() [7]time.Weekday {
	weekDays := [7]time.Weekday{
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
		time.Sunday,
	}

	return weekDays
}
