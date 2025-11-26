package calendar

import (
	"strconv"
	"time"
)

type Calendar struct {
	Months [12]Month
}

type Month struct {
	Label string
	Weeks []Week
}

type Week struct {
	Days [7]Day
}

type Day struct {
	time.Time

	Value int
	Label string
}

func (day *Day) Valid() bool {
	return day.Value != 0
}

func (day *Day) IsSame(other time.Time) bool {
	return day.Year() == other.Year() && day.Month() == other.Month() && day.Day() == other.Day()
}

func (day *Day) String() string {
	return strconv.Itoa(day.Value)
}

func New(year int) Calendar {
	calendar := Calendar{}
	for i := range 12 {
		month := time.Month(i + 1)
		monthStart := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
		nextMonth := monthStart.AddDate(0, 1, 0).Month()
		// Saturday: 6, Monday: 0, Tuesday: 1 ...
		daysSinceMonday := (int(monthStart.Weekday()) + 6) % 7
		duration := time.Hour * time.Duration(24*daysSinceMonday*-1)
		monday := monthStart.Add(duration)

		weeks := []Week{}
		day := monday
		for day.Month() != nextMonth {
			week := Week{}
			for i := range 7 {
				if day.Month() != month {
					week.Days[i] = Day{
						Time: day,
					}
				} else {
					week.Days[i] = Day{
						Value: day.Day(),
						Time:  day,
					}
				}

				day = day.Add(24 * time.Hour)
			}

			weeks = append(weeks, week)
		}

		calendar.Months[i].Weeks = weeks
	}

	return calendar
}
