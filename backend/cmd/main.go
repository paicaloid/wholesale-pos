package main

import (
	"log/slog"
	"wholesale-pos/internal/adapters"
	"wholesale-pos/internal/application"
	"wholesale-pos/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db := adapters.InitDB()
	db.AutoMigrate(&domain.Category{}, &domain.Product{})

	slog.Info("finished initialize database")

	productRepo := adapters.NewGormProductRepository(db)
	productService := application.NewProductService(productRepo)
	productHandler := adapters.NewHttpProductHandler(productService)

	categoryRepo := adapters.NewGormCategoryRepository(db)
	categoryService := application.NewCategoryService(categoryRepo)
	categoryHandler := adapters.NewHttpCategoryHandler(categoryService)

	app.Post("/product", productHandler.CreateProduct)
	app.Post("/category", categoryHandler.CreateCategory)
	app.Listen(":8080")
}
