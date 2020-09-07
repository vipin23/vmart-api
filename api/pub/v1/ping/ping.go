package ping

import (
	"github.com/gofiber/fiber"
	"github.com/vipin23/vmart-api/rest_models"
)

func PingV1(c *fiber.Ctx) {
	res := rest_models.BaseResponse{}
	res.SendMsg(c, "pong")
}
