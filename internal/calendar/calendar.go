package calendar

import (
	"fmt"
	"strconv"
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

func DaysInYear(year int) []time.Time {
	start := time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local)

	result := [366]time.Time{}
	i := 0
	for day := start; day.Year() < start.Year()+1; day = day.Add(time.Hour * 24) {
		result[i] = day
		i++
	}

	return result[:i]
}

type Calendar struct {
	Months [12]Month
}

type Month struct {
	Label string
	Days  []time.Time
	Weeks []Week
}

type Week struct {
	Days [7]Day
}

type Day struct {
	Value int
	Label string
}

func (day *Day) Valid() bool {
	return day.Value != 0
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
					week.Days[i] = Day{}
				} else {
					week.Days[i] = Day{
						Value: day.Day(),
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
