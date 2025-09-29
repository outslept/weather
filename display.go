package main

import (
	"fmt"
	"math"
	"time"
)

func showWeather(w *Weather, loc *Location) {
	icon := pickIcon(w.Current.WeatherCode, w.Current.IsDay == 1)

	printHeader(loc, w.Timezone)
	printIcon(icon)
	printTemperature(w)
	printCondition(icon.Description)
	printMetrics(w)
	printSunTimes(w)
}

func printHeader(loc *Location, tz string) {
	fmt.Println()

	now := time.Now()
	if tz != "" {
		if l, err := time.LoadLocation(tz); err == nil {
			now = now.In(l)
		}
	}

	fmt.Printf("  %s%s%s, %s%s\n", bold, white, loc.Name, loc.Country, reset)
	fmt.Printf("  %s%s%s\n\n", gray, now.Format("Monday, January 2, 2006 • 15:04 MST"), reset)
}

func printIcon(icon WeatherIcon) {
	for _, line := range icon.ASCII {
		fmt.Printf("   %s\n", line)
	}
	fmt.Println()
}

func printTemperature(w *Weather) {
	color := tempColor(w.Current.Temperature)
	fmt.Printf("   %s%s%.0f°%s", color, bold, w.Current.Temperature, reset)

	if math.Abs(w.Current.ApparentTemp-w.Current.Temperature) >= 1.0 {
		fmt.Printf(" %s(feels like %.0f°)%s", gray, w.Current.ApparentTemp, reset)
	}
	fmt.Println()
}

func printCondition(desc string) {
	fmt.Printf("   %s%s%s\n\n", gray, desc, reset)
}

func printMetrics(w *Weather) {
	fmt.Printf("   %sHumidity%s %d%%", gray, reset, w.Current.Humidity)
	fmt.Printf("   %sWind%s %.0f km/h %s", gray, reset, w.Current.WindSpeed, windDirection(w.Current.WindDirection))

	if len(w.Daily.TempMax) > 0 && len(w.Daily.TempMin) > 0 {
		fmt.Printf("   %sHi / Lo%s %.0f° / %.0f°", gray, reset, w.Daily.TempMax[0], w.Daily.TempMin[0])
	}
	fmt.Println()
}

func printSunTimes(w *Weather) {
	if len(w.Daily.Sunrise) == 0 || len(w.Daily.Sunset) == 0 {
		return
	}

	loc := time.Local
	if w.Timezone != "" {
		if l, err := time.LoadLocation(w.Timezone); err == nil {
			loc = l
		}
	}

	// Open-Meteo daily sunrise/sunset come as local time "2006-01-02T15:04"
	layout := "2006-01-02T15:04"
	sunrise, err1 := time.ParseInLocation(layout, w.Daily.Sunrise[0], loc)
	sunset, err2 := time.ParseInLocation(layout, w.Daily.Sunset[0], loc)
	if err1 == nil && err2 == nil {
		fmt.Printf("\n   %sSunrise%s %s", gray, reset, sunrise.Format("15:04"))
		fmt.Printf("   %sSunset%s %s\n\n", gray, reset, sunset.Format("15:04"))
	}
}

func tempColor(t float64) string {
	switch {
	case t >= 30:
		return red
	case t >= 20:
		return yellow
	case t >= 10:
		return green
	case t >= 0:
		return brightBlue
	default:
		return blue
	}
}

func windDirection(deg int) string {
	dirs := []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
	i := int((float64(deg)+22.5)/45.0) % 8
	return dirs[i]
}

func pickIcon(code int, isDay bool) WeatherIcon {
	if code == 0 && !isDay {
		return iconMap[100]
	}
	if icon, ok := iconMap[code]; ok {
		return icon
	}
	return iconMap[999]
}