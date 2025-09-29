package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	weatherAPI = "https://api.open-meteo.com/v1/forecast"
	geocodeAPI = "https://geocoding-api.open-meteo.com/v1/search"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func getLocation(cityName string) (*Location, error) {
	encodedCity := url.QueryEscape(cityName)
	requestURL := fmt.Sprintf("%s?name=%s&count=1&language=en&format=json", geocodeAPI, encodedCity)

	resp, err := httpClient.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("geocoding: %s: %s", resp.Status, string(b))
	}

	var geo GeocodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&geo); err != nil {
		return nil, err
	}
	if len(geo.Results) == 0 {
		return nil, fmt.Errorf("city %q not found", cityName)
	}

	return &geo.Results[0], nil
}

func getWeather(lat, lon float64) (*Weather, error) {
	requestURL := fmt.Sprintf(
		"%s?latitude=%.5f&longitude=%.5f&current=temperature_2m,relative_humidity_2m,apparent_temperature,wind_speed_10m,wind_direction_10m,weather_code,is_day&daily=temperature_2m_max,temperature_2m_min,sunrise,sunset&timezone=auto",
		weatherAPI, lat, lon,
	)

	resp, err := httpClient.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("weather api: %s: %s", resp.Status, string(b))
	}

	var w Weather
	if err := json.NewDecoder(resp.Body).Decode(&w); err != nil {
		return nil, err
	}

	return &w, nil
}
