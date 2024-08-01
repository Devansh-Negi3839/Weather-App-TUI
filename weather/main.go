package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	API_KEY = "8d2a110b6ad468ae1a0e459757cf659d"
	API_URL = "https://api.openweathermap.org/data/2.5/weather"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please enter your city and country below:")

		city, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		city = strings.TrimSpace(city) // Remove newline and extra spaces

		country, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		country = strings.TrimSpace(country) // Remove newline and extra spaces

		weatherData, err := getWeatherInfo(city, country)
		if err != nil {
			fmt.Println("Error:", err)
			continue // Continue to prompt for input in case of an error
		}

		displayWeather(weatherData)
	}
}
