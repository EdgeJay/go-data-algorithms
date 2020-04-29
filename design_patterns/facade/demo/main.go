package main

import (
	"fmt"

	"github.com/edgejay/go-data-algorithms/design_patterns/facade/weatherlib"
)

func main() {
	city := "Singapore"
	countryCode := "SG"
	weatherRetriever := weatherlib.CurrentWeatherData{APIKey: "6863b8402f18d37102d9ff01bd7cc6f0"}
	data, err := weatherRetriever.GetByCityAndCountryCode(city, countryCode)

	if err != nil {
		fmt.Printf("Query failed with error: %s\n", err.Error())
		return
	}

	fmt.Printf("Weather at %s, %s is %s. Temperature is %fdeg celcius.\n", city, countryCode, data.Weather[0].Main, data.Main.Temp)
}
