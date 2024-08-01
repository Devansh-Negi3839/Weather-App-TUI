package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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