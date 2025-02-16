package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Handlers(w http.ResponseWriter, r *http.Request) {
	// Получаем рабочий каталог
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to get working directory: %v", err)
	}
	fmt.Println("Current working directory:", wd)

	// Путь к index.html
	indexPath := filepath.Join("A:", "Others", "SunFlow", "web", "static", "index.html")

	// Проверка, существует ли файл
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		fmt.Println("File does not exist:", indexPath)
		http.Error(w, "Unable to load index.html", http.StatusInternalServerError)
		return
	}

	// Чтение и отдача файла index.html
	http.ServeFile(w, r, indexPath)
}
