package main

import (
	"log"
	"net/http"

	"GroceryListOrganizer/internal/handler"
	"GroceryListOrganizer/internal/service"
)

func main() {
	// Загружаем синонимы из файла (файл synonyms.txt должен быть в корневой папке проекта)
	if err := service.LoadSynonyms("synonyms.txt"); err != nil {
		log.Fatalf("Ошибка загрузки синонимов: %v", err)
	}

	http.HandleFunc("/organize", handler.OrganizeHandler)
	log.Println("Сервер запущен на порту 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
