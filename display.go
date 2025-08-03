package main

import (
	"fmt"
	"time"
)

func newWeatherDisplay() *displayManager {
	return &displayManager{
		weatherMap: initWeatherMap(),
	}
}

func (d *displayManager) show(w *weather, loc *location) {
	wd := d.getWeatherDisplay(w.Current.WeatherCode, w.Current.IsDay == 1)

	d.printLocation(loc)
	d.printWeatherArt(wd)
	d.printTemperature(w)
	d.printCondition(wd.Description)
	d.printMetrics(w)
	d.printSunTimes(w)
}

func (d *displayManager) printLocation(loc *location) {
	fmt.Printf("  \n%s%s%s\n", dim, gray, reset)
	fmt.Printf("  %s%s%s, %s%s\n",
		bold, white, loc.Name, loc.Country, reset)
	fmt.Printf("  %s%s%s\n\n",
		gray, time.Now().Format("Monday, January 2, 2006 • 15:04"), reset)
}

func (d *displayManager) printWeatherArt(wd weatherDisplay) {
	for _, line := range wd.ASCII {
		fmt.Printf("   %s\n", line)
	}
	fmt.Println()
}

func (d *displayManager) printTemperature(w *weather) {
	tempColor := d.getTempColor(w.Current.Temperature)
	fmt.Printf("   %s%s%.0f°%s", tempColor, bold, w.Current.Temperature, reset)

	feelsLike := w.Current.ApparentTemp
	if feelsLike != w.Current.Temperature {
		fmt.Printf(" %s(feels like %.0f°)%s", gray, feelsLike, reset)
	}
	fmt.Println()
}

func (d *displayManager) printCondition(description string) {
	fmt.Printf("   %s%s%s\n\n", gray, description, reset)
}

func (d *displayManager) printMetrics(w *weather) {
	fmt.Printf("   %s%s%s %d%%",
		gray, "Humidity", reset, w.Current.Humidity)

	fmt.Printf("   %s%s%s %.0f km/h %s",
		gray, "Wind", reset, w.Current.WindSpeed,
		d.getWindDirection(w.Current.WindDirection))

	if len(w.Daily.TempMax) > 0 && len(w.Daily.TempMin) > 0 {
		fmt.Printf("   %s%s%s %.0f° / %.0f°",
			gray, "Range", reset, w.Daily.TempMin[0], w.Daily.TempMax[0])
	}
	fmt.Println()
}

func (d *displayManager) printSunTimes(w *weather) {
	if len(w.Daily.Sunrise) == 0 || len(w.Daily.Sunset) == 0 {
		return
	}

	sunriseStr := w.Daily.Sunrise[0]
	sunsetStr := w.Daily.Sunset[0]

	sunrise, err1 := time.Parse("2006-01-02T15:04", sunriseStr)
	if err1 != nil {
		sunrise, err1 = time.Parse(time.RFC3339, sunriseStr)
	}

	sunset, err2 := time.Parse("2006-01-02T15:04", sunsetStr)
	if err2 != nil {
		sunset, err2 = time.Parse(time.RFC3339, sunsetStr)
	}

	if err1 == nil && err2 == nil {
		fmt.Printf("\n   %s%s%s %s",
			gray, "Sunrise", reset, sunrise.Format("15:04"))
		fmt.Printf("   %s%s%s %s\n",
			gray, "Sunset", reset, sunset.Format("15:04"))
	}
	fmt.Println()
}

func (d *displayManager) getTempColor(temp float64) string {
	switch {
	case temp >= 30:
		return red
	case temp >= 20:
		return yellow
	case temp >= 10:
		return green
	case temp >= 0:
		return brightBlue
	default:
		return blue
	}
}

func (d *displayManager) getWindDirection(degrees int) string {
	directions := []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
	index := int((float64(degrees)+22.5)/45.0) % 8
	return directions[index]
}

func (d *displayManager) getWeatherDisplay(code int, isDay bool) weatherDisplay {
	if wd, exists := d.weatherMap[code]; exists {
		if !isDay && code == 0 {
			return d.weatherMap[100]
		}
		return wd
	}
	return d.weatherMap[999]
}
