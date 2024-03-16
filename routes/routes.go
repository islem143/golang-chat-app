package routes

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/islem143/go-chat/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func extractKeyFromErrorMessage(message string) string {
	parts := strings.Split(message, "{ : \"")
	if len(parts) < 2 {
		return "" // unable to extract
	}
	key := strings.Split(parts[1], "\"")[0]
	return key
}
func IsDup(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				fmt.Println(extractKeyFromErrorMessage(we.Details.String()))
				return true
			}
		}
	}
	return false
}

func Setup(app *fiber.App) {
	api := app.Group("/")

	api.Get("/", func(c *fiber.Ctx) error {

		//user := models.User{Name: "islem", Email: "test@test.com"}
		user, err := models.FindUserByEmail("test@ŧest.com")
		if err != nil {
			return err

		}
		return c.JSON(user)

	})
	api.Post("/", func(c *fiber.Ctx) error {

		user := new(models.User)
		if err := c.BodyParser(user); err != nil {
			return err
		}
		err := models.InsertUser(user)
		if err != nil {
			return err
		}

		return c.JSON(user)

	})

}
