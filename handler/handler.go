package handler

import (
	"os"
	"strconv"
	"time"

	"github.com/aom31/fibergoapi/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	//1. check user logic is valid

	//instance user from struct User
	user := new(models.User)

	// parse body request to user
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Email != models.MockDataUserLogic.Email || user.Password != models.MockDataUserLogic.Password {
		return fiber.ErrUnauthorized
	}

	token, _ := createTokenPattern(c, user)

	return c.JSON(fiber.Map{
		"message": "login success",
		"token":   token,
	})
}

func createTokenPattern(c *fiber.Ctx, user *models.User) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_secret")))
	if err != nil {
		return "", c.SendStatus(fiber.StatusInternalServerError)
	}

	return t, nil

}

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

func CreateBook(c *fiber.Ctx) error {
	// instance book for get data request
	book := new(models.Book)

	// parse data request to struct
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	dataBook := models.InitDataBook()
	dataBook = append(dataBook, *book)

	return c.JSON(dataBook)
}

func UpdateBookByID(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	dataBook := models.InitDataBook()

	bookUpdate := new(models.Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for ind, book := range dataBook {
		if id == book.ID {
			dataBook[ind].Title = bookUpdate.Title
			dataBook[ind].Author = bookUpdate.Author
			return c.JSON(dataBook[ind])
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func DeleteBookByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	dataBook := models.InitDataBook()

	for ind, book := range dataBook {
		if id == book.ID {
			// delete book[ind] from slice
			dataBook = append(dataBook[:ind], dataBook[ind+1:]...)

			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func UploadFile(c *fiber.Ctx) error {
	key := "image"
	file, err := c.FormFile(key)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}

	//save file to folder local
	if err := c.SaveFile(file, "./uploads/"+file.Filename); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString("file upload complete")
}

func GetEnv(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"DBNAME": os.Getenv("DBNAME"),
	})
}
