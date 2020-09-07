package v1

import (
	"github.com/gofiber/fiber"
	"github.com/vipin23/vmart-api/api/pub/v1/auth"
	"github.com/vipin23/vmart-api/api/pub/v1/ping"
)

func V1(c fiber.Router) {
	c.Get("/ping", ping.PingV1)
	c.Post("/login", auth.Login)
}
