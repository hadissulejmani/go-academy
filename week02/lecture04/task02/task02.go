package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cityValues := groupSlices(citiesAndPrices())
	for key, value := range cityValues {
		fmt.Println(key, value)
	}
}

func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 100
	// randomly choose cities
	cities := make([]string, dataPointCount)
	for i := range cities {
		cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}
	prices := make([]int, dataPointCount)
	for i := range prices {
		prices[i] = rand.Intn(100)
	}
	return cities, prices
}

func groupSlices(keySlice []string, valueSlice []int) map[string][]int {
	result := make(map[string][]int, 5)
	for i, k := range keySlice {
		result[k] = append(result[k], valueSlice[i])
	}
	return result
}
