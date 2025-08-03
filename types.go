package main

type location struct {
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type geocodeResponse struct {
	Results []location `json:"results"`
}

type weather struct {
	Current struct {
		Temperature   float64 `json:"temperature_2m"`
		Humidity      int     `json:"relative_humidity_2m"`
		ApparentTemp  float64 `json:"apparent_temperature"`
		WindSpeed     float64 `json:"wind_speed_10m"`
		WindDirection int     `json:"wind_direction_10m"`
		WeatherCode   int     `json:"weather_code"`
		IsDay         int     `json:"is_day"`
	} `json:"current"`
	Daily struct {
		TempMax []float64 `json:"temperature_2m_max"`
		TempMin []float64 `json:"temperature_2m_min"`
		Sunrise []string  `json:"sunrise"`
		Sunset  []string  `json:"sunset"`
	} `json:"daily"`
}

type weatherDisplay struct {
	Description string
	ASCII       []string
}

type displayManager struct {
	weatherMap map[int]weatherDisplay
}
