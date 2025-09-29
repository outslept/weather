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

	city := strings.Join(os.Args[1:], " ")

	loc, err := getLocation(city)
	if err != nil {
		fmt.Printf("%sError:%s %v\n", red, reset, err)
		return
	}

	w, err := getWeather(loc.Latitude, loc.Longitude)
	if err != nil {
		fmt.Printf("%sError:%s %v\n", red, reset, err)
		return
	}

	showWeather(w, loc)
}