package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"wb0-app/cache"
	"wb0-app/client"
	"wb0-app/config"
	"wb0-app/http/handlers"
	"wb0-app/logger"
	"wb0-app/models"
	"wb0-app/storage"
)

func main() {
	fmt.Println("App run!")
	config.Load()

	logger.Setup()

	storage := storage.New()

	cache := cache.New()

	subscriber := client.New()
	subscriber.Subscribe(models.Order{}, storage.Save, cache.Save)

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/order", handlers.MakeOrderHandler(cache))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error("server: cannot start", slog.String("error", err.Error()))
	}
	slog.Error("server: down")

}
