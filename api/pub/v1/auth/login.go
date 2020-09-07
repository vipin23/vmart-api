package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/vipin23/vmart-api/models"
	"github.com/vipin23/vmart-api/rest_models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

func Login(c *fiber.Ctx) {
	reqUser := models.User{}
	err := c.BodyParser(&reqUser)
	if err != nil {
		res := rest_models.BaseResponse{Code: http.StatusUnprocessableEntity}
		res.Error(c, err)
		return
	}
	err = reqUser.Login()
	if err != nil {
		c.SendStatus(fiber.StatusUnauthorized)
		return
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = reqUser
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}

	c.JSON(fiber.Map{"token": t})
}
