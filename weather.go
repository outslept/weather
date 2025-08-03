package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	weatherAPI = "https://api.open-meteo.com/v1/forecast"
	geocodeAPI = "https://geocoding-api.open-meteo.com/v1/search"
)

func getLocation(cityName string) (*location, error) {
	encodedCity := url.QueryEscape(cityName)
	requestURL := fmt.Sprintf("%s?name=%s&count=1&language=en&format=json", geocodeAPI, encodedCity)

	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var geoResp geocodeResponse
	if err := json.Unmarshal(body, &geoResp); err != nil {
		return nil, err
	}

	if len(geoResp.Results) == 0 {
		return nil, fmt.Errorf("city '%s' not found", cityName)
	}

	return &geoResp.Results[0], nil
}

func getWeather(lat, lon float64) (*weather, error) {
	requestURL := fmt.Sprintf("%s?latitude=%.2f&longitude=%.2f&current=temperature_2m,relative_humidity_2m,apparent_temperature,wind_speed_10m,wind_direction_10m,weather_code,is_day&daily=temperature_2m_max,temperature_2m_min,sunrise,sunset&timezone=auto",
		weatherAPI, lat, lon)

	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var w weather
	if err := json.Unmarshal(body, &w); err != nil {
		return nil, err
	}

	return &w, nil
}
