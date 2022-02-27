package main

import "fmt"

func main() {

	fmt.Println(daysInMonth(10, 2012))
}

func daysInMonth(month int, year int) (int, bool) {
	leapYear := false
	var daysInMonth int

	if year%4 == 0 {
		if month == 2 {
			return 29, true
		}
		leapYear = true
	}

	switch {
	case month%2 != 0:
		daysInMonth = 31
	case month%2 == 0 && month != 2:
		daysInMonth = 30
	case month == 2:
		daysInMonth = 28
	}
	return daysInMonth, leapYear
}
