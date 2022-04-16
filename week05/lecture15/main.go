package main

import (
	"fmt"
	"sort"
	"time"
)

func sortDates(format string, dates ...string) ([]string, error) {
	var formatedDates []time.Time
	var output []string
	const layout = "Jan-2-2006"

	for _, v := range dates {
		date, err := time.Parse(layout, v)
		if err == nil {
			formatedDates = append(formatedDates, date)
		}
	}

	sort.Slice(formatedDates, func(i, j int) bool {
		return formatedDates[i].Before(formatedDates[j])
	})

	for _, v := range formatedDates {
		output = append(output, v.String())
	}

	return output, nil
}

func main() {
	const layout = "Jan-2-2006"
	sortedDates, err := sortDates(layout, "Sep-14-2008", "Dec-03-2021", "Mar-18-2022")

	if err != nil {
		fmt.Println(err)
	}
	for _, v := range sortedDates {
		fmt.Println(v)
	}
}
