package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Структура для получения данных из API
type SunriseSunsetAPIResponse struct {
	Results struct {
		Sunrise string `json:"sunrise"`
		Sunset  string `json:"sunset"`
	} `json:"results"`
	Status string `json:"status"`
}

// Функция для получения данных о солнечном свете для заданных координат
func GetSunlightData(latitude, longitude float64) (*SunriseSunsetAPIResponse, error) {
	// Формируем URL для запроса к API
	apiURL := fmt.Sprintf("https://api.sunrise-sunset.org/json?lat=%f&lng=%f&formatted=0", latitude, longitude)

	// Выполняем HTTP-запрос
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching sunrise/sunset data: %v", err)
	}
	defer resp.Body.Close()

	// Декодируем JSON-ответ
	var result SunriseSunsetAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding sunrise/sunset data: %v", err)
	}

	// Возвращаем данные о времени восхода и заката
	if result.Status != "OK" {
		return nil, fmt.Errorf("error in API response: %s", result.Status)
	}

	return &result, nil
}
