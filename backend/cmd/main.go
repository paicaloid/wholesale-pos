package main

import (
	"log/slog"
	"wholesale-pos/internal/adapters"
	"wholesale-pos/internal/domain"
)

func main() {
	db := adapters.InitDB()
	db.AutoMigrate(&domain.Category{}, &domain.Product{})

	slog.Info("finished initialize database")
}
