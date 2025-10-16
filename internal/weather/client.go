package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

func (c *Client) getWeatherByCoordinates(latitude, longitude float64) error {
	params := url.Values{}
	params.Add("lat", fmt.Sprintf("%f", latitude))
	params.Add("lon", fmt.Sprintf("%f", longitude))
	params.Add("appid", c.apiKey)
	params.Add("units", unitsType)
	params.Add("lang", languageCode)

	fullUrl := baseUrl + "?" + params.Encode()

	resp, err := http.Get(fullUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("weather API error: %s", resp.Status)
	}
	fmt.Println(resp)
	var response Response

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}
