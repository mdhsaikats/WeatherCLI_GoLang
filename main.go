package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const apiKey = "a9aea2d8648441db7864353eca7b78be"

type WeatherResponse struct {
	Name string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
	} `json:"sys"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Cod int `json:"cod"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: weather <city name>")
		return
	}

	city := os.Args[1]

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Network error:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unable to fetch weather data.")
		return
	}
	var weather WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		fmt.Println("Error parsing data:", err)
		return
	}
	if weather.Cod == 404 {
		fmt.Println("City not found.")
		return
	}

	fmt.Println("================================")
	fmt.Printf("City: %s, %s\n", weather.Name, weather.Sys.Country)
	fmt.Printf("Temperature: %.1f°C\n", weather.Main.Temp)
	fmt.Printf("Feels Like: %.1f°C\n", weather.Main.FeelsLike)
	fmt.Printf("Weather: %s\n", weather.Weather[0].Description)
	fmt.Printf("Wind Speed: %.1f m/s\n", weather.Wind.Speed)
	fmt.Println("================================")

}
