package main

import (
	"github.com/aom31/fibergoapi/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// init http server with fiber
	appServer := fiber.New()

	// use route http
	appServer.Get("/books", handler.GetBooks)
	appServer.Get("/book/:id", handler.GetBookByID)
	appServer.Post("/book", handler.CreateBook)
	appServer.Put("/book", handler.UpdateBookByID)
	appServer.Delete("/book/:id", handler.DeleteBookByID)
	appServer.Post("/uploadfile", handler.UploadFile)

	// start server with port
	appServer.Listen(":8080")
}
