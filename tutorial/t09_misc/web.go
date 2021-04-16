package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

type Tour struct {
	Name, Price string
}

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromService(url)

	tours := toursFromJson(content)
	// fmt.Println(tours[0])

	for _, tour := range tours {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero) // base10 with 2-digit precision
		fmt.Printf("%v ($%.2f)\n", tour.Name, price)
	}

	fmt.Printf("There are %d tours\n", len(tours))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromService(url string) string {

	resp, err := http.Get(url)
	checkError(err)

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes)
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}

	return tours
}
