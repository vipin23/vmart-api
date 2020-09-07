package v2

import (
	"github.com/gofiber/fiber"
	"github.com/vipin23/vmart-api/api/pub/v2/ping"
)

func V2(c fiber.Router) {
	c.Get("/ping", ping.PingV2)
}
