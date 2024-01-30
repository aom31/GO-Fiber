package main

import (
	"log"
	"os"

	"github.com/aom31/fibergoapi/handler"
	"github.com/aom31/fibergoapi/middleware"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/joho/godotenv"
)

func main() {

	// init http server with fiber
	appServer := fiber.New()

	//load env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	appServer.Post("/login", handler.Login)

	// JWT Middleware
	appServer.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_secret")),
	}))

	//use middleware: if role admin can use API below
	appServer.Use(middleware.ValidateRoleAuthorize)

	// use route http
	appServer.Get("/books", handler.GetBooks)
	appServer.Get("/book/:id", handler.GetBookByID)
	appServer.Post("/book", handler.CreateBook)
	appServer.Put("/book", handler.UpdateBookByID)
	appServer.Delete("/book/:id", handler.DeleteBookByID)
	appServer.Post("/uploadfile", handler.UploadFile)

	appServer.Get("/config", handler.GetEnv)

	// start server with port
	appServer.Listen(":8080")
}
