package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	f, err := os.Open("./USA_cars_datasets.csv")
	checkError(err)
	defer f.Close()

	rdr := csv.NewReader(f)
	rows, err := rdr.ReadAll()
	checkError(err)

	fmt.Printf("Mean price: $%.2f\n", mean(rows, 1))
	fmt.Printf("Median price: $%.2f\n", median(rows, 1))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func mean(rows [][]string, idx int) float64 {
	var total float64
	for i, row := range rows {
		if i != 0 {
			item, _ := strconv.ParseFloat(row[idx], 64)
			total += item
		}
	}
	return total / float64(len(rows)-1)
}

func median(rows [][]string, idx int) float64 {

	var sorted []float64
	for i, row := range rows {
		if i != 0 {
			val, _ := strconv.ParseFloat(row[idx], 64)
			sorted = append(sorted, val)
		}
	}

	// Sort the slice
	sort.Float64s(sorted)

	middle := len(sorted) / 2
	if isEven(sorted) {
		high := sorted[middle]
		low := sorted[middle-1]
		return (high + low) / 2
	}
	return sorted[middle]
}

func isEven(arr []float64) bool {
	if len(arr)%2 == 0 {
		return true
	}
	return false
}
