package main

import (
	"fmt"
	"log"
	"net/http"

	_ "leaflet-back/docs" // Этот импорт необходим для swag
	"leaflet-back/internal/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Leaflet Tile Server API
// @version 1.0
// @description Это сервер для раздачи тайлов карт и управления ими.
// @host localhost:8080
// @BasePath /
func main() {
	// Создаем новый маршрутизатор
	mux := http.NewServeMux()

	// Регистрируем обработчики
	mux.HandleFunc("/maps/info", handlers.GetMapInfo)
	mux.HandleFunc("/upload", handlers.UploadMap)
	mux.HandleFunc("/tiles/", handlers.GetTile) // Обратите внимание на слеш в конце

	// Обработчик для Swagger UI
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// Приветственное сообщение
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the tile server! Visit /swagger/index.html for API documentation.")
	})

	fmt.Println("Starting server on :8080")
	fmt.Println("Swagger UI is available at http://localhost:8080/swagger/index.html")
	
	// Запускаем сервер с нашим маршрутизатором
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
