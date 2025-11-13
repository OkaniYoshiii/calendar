package translation

import "time"

func Month(month time.Month) string {
	switch month {
	case time.January:
		return "Janvier"
	case time.February:
		return "Février"
	case time.March:
		return "Mars"
	case time.April:
		return "Avril"
	case time.May:
		return "Mai"
	case time.June:
		return "Juin"
	case time.July:
		return "Juillet"
	case time.August:
		return "Août"
	case time.September:
		return "Septembre"
	case time.October:
		return "Octobre"
	case time.November:
		return "Novembre"
	case time.December:
		return "Décembre"
	default:
		panic("Incorrect month value")
	}
}
