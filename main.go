package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"time"
)

type Child struct {
	Birthday time.Time
	Name     string
}

type HistoryEntry struct {
	Child Child
	Date  time.Time
}

func main() {
	childs := make([]Child, 5)
	childs[0] = Child{Name: "Imen", Birthday: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local)}
	childs[1] = Child{Name: "Noah", Birthday: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local)}
	childs[2] = Child{Name: "Joel", Birthday: time.Date(2000, time.February, 12, 0, 0, 0, 0, time.Local)}
	childs[3] = Child{Name: "Ellie", Birthday: time.Date(2000, time.January, 16, 0, 0, 0, 0, time.Local)}
	childs[4] = Child{Name: "Sebastian", Birthday: time.Date(2000, time.April, 2, 0, 0, 0, 0, time.Local)}

	scanner := bufio.NewScanner(os.Stdin)
	history := []HistoryEntry{}

	fallback := time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	start, err := AskDate("Rentrer une date de départ (YYYY-MM-DD): ", scanner, fallback)
	if err != nil {
		log.Fatal(err)
	}

	// Redéfinit la date de départ au début de la semaine (Lundi)
	start = start.AddDate(0, 0, -int(start.Weekday())+1)

	fmt.Println("Modification de la date de départ pour commencer en début de semaine : ")
	fmt.Println(start)

	end := time.Date(start.Year()+1, start.Month(), start.Day(), start.Hour(), start.Minute(), start.Second(), start.Nanosecond(), start.Location())

	fmt.Println(end, history)

	for {
		date := start

		// Pour chaque semaine de l'année entre start et end
		childsIndexes := Indexes(childs)
		for date.Compare(end) <= 0 {
			fmt.Println("")
			fmt.Println("--- Debut semaine : ", date, " ---")

			closest := 0
			for i, index := range childsIndexes {
				child := childs[index]
				birthday := child.Birthday

				isBirthdayInWeek := IsInWeekIgnoringYear(date, birthday)

				if !isBirthdayInWeek {
					continue
				}

				fmt.Println(">>> Anniversaire de ", child.Name, " dans la semaine du ", date)
				closest = i
			}

			index := childsIndexes[closest]
			fmt.Println(">>> Peluche attribuée à : ", childs[index].Name)

			if len(childsIndexes)-1 <= 0 {
				fmt.Println(">>>> Pool d'enfant vidé")
				childsIndexes = Indexes(childs)
			} else {
				temp := childsIndexes[closest]
				childsIndexes[closest] = childsIndexes[len(childsIndexes)-1]
				childsIndexes[len(childsIndexes)-1] = temp
				childsIndexes = childsIndexes[:len(childsIndexes)-1]
			}

			date = date.AddDate(0, 0, 7)
			fmt.Println("-- Fin semaine : ", date, " --")
		}

		break
	}
}

func Indexes[T any](slice []T) []int {
	indexes := make([]int, len(slice))
	for i := range len(slice) {
		indexes[i] = i
	}

	return indexes
}

func IsInWeekIgnoringYear(base, compared time.Time) bool {
	same := compared.AddDate(base.Year()-compared.Year(), 0, 0)

	// Test de l'année d'avant et l'année d'après
	candidates := []time.Time{
		same.AddDate(-1, 0, 0),
		same,
		same.AddDate(1, 0, 0),
	}

	baseYear, baseWeek := base.ISOWeek()

	for _, candidate := range candidates {
		year, week := candidate.ISOWeek()
		if year == baseYear && week == baseWeek {
			return true
		}
	}

	return false
}

func DifferenceIgnoringYear(base time.Time, compared time.Time) time.Duration {
	same := compared.AddDate(base.Year()-compared.Year(), 0, 0)
	before := same.AddDate(-1, 0, 0)
	after := same.AddDate(1, 0, 0)

	differences := []time.Duration{
		base.Sub(before).Abs(),
		base.Sub(same).Abs(),
		base.Sub(after).Abs(),
	}

	min := slices.Min(differences)

	return min
}

func AskDate(message string, scanner *bufio.Scanner, fallback time.Time) (time.Time, error) {
	fmt.Println(message)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return time.Time{}, err
	}
	result := string(scanner.Bytes())

	if result == "" {
		return fallback, nil
	}

	date, err := time.Parse(time.DateOnly, result)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}
