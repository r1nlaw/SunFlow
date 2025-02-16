package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SunriseSunsetAPIResponse struct {
	Results struct {
		Sunrise string `json:"sunrise"`
		Sunset  string `json:"sunset"`
	} `json:"results"`
	Status string `json:"status"`
}

func GetSunriseSunsetInfo(latitude, longitude float64) (*SunriseSunsetAPIResponse, error) {
	apiURL := fmt.Sprintf("https://api.sunrise-sunset.org/json?lat=%f&lng=%f&formatted=0", latitude, longitude)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching sunrise/sunset data: %v", err)
	}
	defer resp.Body.Close()

	var result SunriseSunsetAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding sunrise/sunset data: %v", err)
	}

	if result.Status != "OK" {
		return nil, fmt.Errorf("error in API response: %s", resp.Status)
	}

	return &result, nil
}

func CalculateTimeToSunrise(latitude, longitude float64) (string, error) {
	data, err := GetSunriseSunsetInfo(latitude, longitude)
	if err != nil {
		return "", err
	}

	sunrise, err := time.Parse(time.RFC3339, data.Results.Sunrise)
	if err != nil {
		return "", fmt.Errorf("error parsing sunrise time: %v", err)
	}

	now := time.Now()

	duration := sunrise.Sub(now)

	if duration < 0 {
		return "", fmt.Errorf("sunrise has already passed")
	}

	return fmt.Sprintf("Time to sunrise: %v", duration), nil
}
