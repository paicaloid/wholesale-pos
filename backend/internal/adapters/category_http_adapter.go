package adapters

import (
	"wholesale-pos/internal/domain"
	"wholesale-pos/internal/ports"

	"github.com/gofiber/fiber/v2"
)

type httpCategoryHandler struct {
	service ports.CategoryService
}

func NewHttpCategoryHandler(service ports.CategoryService) *httpCategoryHandler {
	return &httpCategoryHandler{service: service}
}

func (h *httpCategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var category domain.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.service.CreateCategory(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Category created successfully",
	})

}
