package main

import (
	"fmt"
	"log"
	"net/http"
	"sunflow/internal/handlers"
)

func main() {
	// Обработчик для главной страницы
	http.HandleFunc("/", handlers.Handlers)

	// Обработчик для солнечного света
	http.HandleFunc("/sunlight", handlers.SunlightHandler)

	// Статические файлы теперь доступны по пути /static/
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("A:/Others/SunFlow/web/static"))))

	// Запуск сервера
	port := ":8080"
	fmt.Println("Server started on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
