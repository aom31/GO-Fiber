package handler

import (
	"strconv"

	"github.com/aom31/fibergoapi/models"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	return c.JSON(models.InitDataBook())
}

func GetBookByID(c *fiber.Ctx) error {
	bookIDReq := c.Params("id")
	bookData := models.InitDataBook()

	for _, book := range bookData {
		id, _ := strconv.Atoi(bookIDReq)
		if id == book.ID {
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)

}
