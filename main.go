package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%sUsage:%s weather <city>\n", gray, reset)
		fmt.Printf("%sExample:%s weather Moscow\n", gray, reset)
		return
	}

	cityName := strings.Join(os.Args[1:], " ")

	location, err := getLocation(cityName)
	if err != nil {
		fmt.Printf("%sError:%s %v\n", red, reset, err)
		return
	}

	weather, err := getWeather(location.Latitude, location.Longitude)
	if err != nil {
		fmt.Printf("%sError:%s %v\n", red, reset, err)
		return
	}

	display := newWeatherDisplay()
	display.show(weather, location)
}
