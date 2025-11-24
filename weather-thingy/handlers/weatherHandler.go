package handlers

import (
	"fmt"
	"net/http"

	"github.com/Anna-Koudelkova/Ready-set-GO/weather-thingy/apilogic"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/weatherpage" {
		http.Error(w, "404 not found, as usual", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	city := r.FormValue("city")

	temperature, err := apilogic.GetTemperature(city)
	
	if err != nil {
		http.Error(w, "Failed normally to fetch the city data", http.StatusInternalServerError)
		return
	}

	fahrenheit := (temperature * 1.8)+32

	fmt.Fprintf(w, "The temperature in your desired city %v is: %.2f in Celsius /%.2f in Fahrenheit", city, temperature, fahrenheit)

}
