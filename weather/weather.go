package weather

import (
	"encoding/json"
	"fmt"
	"net/http"

)

type Weather struct {
	Temperature float64 `json:"temp"`
	Description string  `json:"description"`
}
const (
OpenWeatherMapAPIKey = "a120c6943839b623b1a4336fbae8f377"
)

func GetCurrentWeather(location string) (Weather, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", location, OpenWeatherMapAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return Weather{}, err
	}
	defer resp.Body.Close()

	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return Weather{}, err
	}

	weather := Weather{
		Temperature: data.Main.Temp,
		Description: data.Weather[0].Description,
	}

	return weather, nil
}