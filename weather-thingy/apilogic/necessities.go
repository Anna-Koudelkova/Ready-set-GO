package apilogic

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type APIConfig struct {
	APIKey string
}

type WeatherResponse struct {
	CityName string `json:"name"`
	Main     struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func loadConfig() (APIConfig, error) {
	var config APIConfig

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed loading .env file: %v", err)
		return config, err
	}

	key, exists := os.LookupEnv("OPENWEATHER_API_KEY")
	if !exists {
		log.Fatal("The API key is missing in env")
		return config, errors.New("Missing API key")
	}

	config = APIConfig{APIKey: key}
	return config, nil
}

func GetTemperature(city string) (float64, error) {
	config, err := loadConfig()
	if err != nil {
		return 0, fmt.Errorf("Failed terribly to load API key: %v", err)
	}

	apiURL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?appid=%s&q=%s", config.APIKey, city)
	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, fmt.Errorf("Failed horribly to make a request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("Failed miserably to read the response body: %v", err)
	}

	var weatherData WeatherResponse
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return 0, fmt.Errorf("Completely failed to parse JSON: %v", err)
	}

	celsius := weatherData.Main.Kelvin - 273.15
	return celsius, nil
}
