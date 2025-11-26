package calendar

import (
	"iter"
	"strconv"
	"time"
)

const DaysInWeek = 7
const MonthsInYear = 12

type Calendar[T any] struct {
	Months [MonthsInYear]Month[T]
}

type Month[T any] struct {
	Label string
	Weeks []Week[T]
}

type Week[T any] struct {
	Days [DaysInWeek]Day[T]
}

type Day[T any] struct {
	time.Time

	Value int
	Label string

	Payload T
}

func (day *Day[T]) Valid() bool {
	return day.Value != 0
}

func (day *Day[T]) IsSame(other time.Time) bool {
	return day.Year() == other.Year() && day.Month() == other.Month() && day.Day() == other.Day()
}

func (day *Day[T]) String() string {
	return strconv.Itoa(day.Value)
}

// Creates a new calendar
//
// payloadFunc is used to add additional payload on each day if wanted
// This can be used to attach events to some days like anniversaries
// or notes.
//
// translationFunc is used to convert a time.Month to a label that you want to
// show to a user for example.
func New[T any](year int, payloadFunc func(*Day[T]), translationFunc func(time.Month) string) Calendar[T] {
	calendar := Calendar[T]{}
	for i := range 12 {
		month := time.Month(i + 1)
		monthStart := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
		nextMonth := monthStart.AddDate(0, 1, 0).Month()
		// Saturday: 6, Monday: 0, Tuesday: 1 ...
		daysSinceMonday := (int(monthStart.Weekday()) + 6) % 7
		duration := time.Hour * time.Duration(24*daysSinceMonday*-1)
		monday := monthStart.Add(duration)

		weeks := []Week[T]{}
		day := monday
		for day.Month() != nextMonth {
			week := Week[T]{}
			for i := range 7 {
				current := Day[T]{
					Time: day,
				}

				if day.Month() == month {
					current.Value = day.Day()
				}

				payloadFunc(&current)

				week.Days[i] = current

				day = day.Add(24 * time.Hour)
			}

			weeks = append(weeks, week)
		}

		calendar.Months[i].Weeks = weeks
		calendar.Months[i].Label = translationFunc(month)
	}

	return calendar
}

func Days[T any](cal Calendar[T]) iter.Seq[Day[T]] {
	return func(yield func(Day[T]) bool) {
		for _, month := range cal.Months {
			for _, week := range month.Weeks {
				for _, day := range week.Days {
					if !yield(day) {
						return
					}
				}
			}
		}
	}
}
