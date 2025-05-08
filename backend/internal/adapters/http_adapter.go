package adapters

import (
	"wholesale-pos/internal/domain"
	"wholesale-pos/internal/ports"

	"github.com/gofiber/fiber/v2"
)

type httpProductHandler struct {
	service ports.ProductService
}

func NewHttpProductHandler(service ports.ProductService) *httpProductHandler {
	return &httpProductHandler{service: service}
}

func (h *httpProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product domain.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.service.CreateProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
	})
}
