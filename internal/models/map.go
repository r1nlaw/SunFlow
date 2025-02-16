package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SunlightData struct {
	Sunrise   string `json:"sunrise"`
	Sunset    string `json:"sunset"`
	DayLength string `json:"day_length"`
}

const apiURL = "https://api.sunrise-sunset.org/json"

func GetSunLightData(lat, lng string) (*SunlightData, error) {
	url := fmt.Sprintf("%s?lat=%s&lng=%s&formatted=0", apiURL, lat, lng)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if data, ok := result["results"].(map[string]interface{}); ok {
		sunlightData := &SunlightData{
			Sunrise:   data["sunrise"].(string),
			Sunset:    data["sunset"].(string),
			DayLength: data["day_length"].(string),
		}
		return sunlightData, nil
	}

	return nil, fmt.Errorf("unexpected response format")
}
