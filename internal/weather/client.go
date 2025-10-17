package weather

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"tgWeatherBot/internal/router"
)

const (
	baseUrl = "https://api.openweathermap.org/data/2.5/weather?"
)

const (
	unitsType    = "metric"
	languageCode = "ru"
)

type Client struct {
	apiKey string
}

func MakeClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

func (c *Client) GetWeatherByCoordinates(lat, lon string) (router.WeatherData, error) {
	params := url.Values{}
	params.Add("lat", lat)
	params.Add("lon", lon)
	params.Add("appid", c.apiKey)
	params.Add("units", unitsType)
	params.Add("lang", languageCode)

	fullUrl := baseUrl + params.Encode()
	resp, err := http.Get(fullUrl)
	if err != nil {
		return router.WeatherData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return router.WeatherData{},
			fmt.Errorf("weather API error: %s", resp.Status)
	}
	var response Response

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return router.WeatherData{}, err
	}

	weatherData := convertResponseWeatherToWeatherData(response)
	return weatherData, nil
}

func convertResponseWeatherToWeatherData(r Response) router.WeatherData {
	return router.WeatherData{
		Type:      r.Weather[0].Description,
		Temp:      int(math.Round(r.Main.Temp)),
		FeelsLike: int(math.Round(r.Main.FeelsLike)),
		WindSpeed: int(math.Round(r.Main.FeelsLike)),
		Name:      r.Name,
	}
}
