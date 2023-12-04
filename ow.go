package ow

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func get(client *http.Client, lat, lon float64, url, key string) ([]byte, error) {
	resp, err := client.Get(fmt.Sprintf(url, lat, lon, key))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	stream, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return stream, nil
}

func GetAirPollution(client *http.Client, lat, lon float64, key string) (*AirPollution, error) {
	const URL string = "https://api.openweathermap.org/data/2.5/air_pollution?lat=%f&lon=%f&appid=%s"

	data, err := get(client, lat, lon, URL, key)
	if err != nil {
		return nil, err
	}

	var jsonData AirPollution

	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	return &jsonData, nil
}

func GetCurrent(client *http.Client, lat, lon float64, key string) (*Current, error) {
	const URL string = "https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s"

	data, err := get(client, lat, lon, URL, key)
	if err != nil {
		return nil, err
	}

	var jsonData Current

	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	return &jsonData, nil
}

func GetForecast(client *http.Client, lat, lon float64, key string) (*Forecast, error) {
	const URL string = "https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s"

	data, err := get(client, lat, lon, URL, key)
	if err != nil {
		return nil, err
	}

	var jsonData Forecast

	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	return &jsonData, nil
}
