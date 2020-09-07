package ping

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/vipin23/vmart-api/rest_models"
)

func SecurePingV1(c *fiber.Ctx) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	res := rest_models.BaseResponse{}
	res.SendMsg(c, "pong for user : "+name)
}
