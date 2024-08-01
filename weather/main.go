package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Structs to map the JSON response
type WeatherData struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	Id      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

const (
	API_KEY = "8d2a110b6ad468ae1a0e459757cf659d"
	API_URL = "https://api.openweathermap.org/data/2.5/weather"
)

func getWeatherInfo(city, country string) (WeatherData, error) {
	url := fmt.Sprintf("%s?q=%s,%s&appid=%s&units=metric", API_URL, city, country, API_KEY)
	resp, err := http.Get(url)
	if err != nil {
		return WeatherData{}, fmt.Errorf("couldn't fetch weather data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherData{}, fmt.Errorf("couldn't read response body: %w", err)
	}

	var weatherData WeatherData
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return WeatherData{}, fmt.Errorf("couldn't unmarshal JSON: %w", err)
	}

	return weatherData, nil
}

func displayWeather(weatherData WeatherData) {
	fmt.Printf("Weather in %s, %s:\n", weatherData.Name, weatherData.Sys.Country)
	fmt.Printf("Temperature: %.2f°C\n", weatherData.Main.Temp)
	fmt.Printf("Feels Like: %.2f°C\n", weatherData.Main.FeelsLike)
	fmt.Printf("Weather: %s - %s\n", weatherData.Weather[0].Main, weatherData.Weather[0].Description)
	fmt.Printf("Humidity: %d%%\n", weatherData.Main.Humidity)
	fmt.Printf("Pressure: %d hPa\n", weatherData.Main.Pressure)
	fmt.Printf("Wind Speed: %.2f m/s\n", weatherData.Wind.Speed)
	fmt.Printf("Visibility: %d meters\n", weatherData.Visibility)
	fmt.Printf("Coordinates: Lon %.4f, Lat %.4f\n", weatherData.Coord.Lon, weatherData.Coord.Lat)
}

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