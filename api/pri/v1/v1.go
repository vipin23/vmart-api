package v1

import (
	"github.com/gofiber/fiber"
	"github.com/vipin23/vmart-api/api/pri/v1/ping"
	"github.com/vipin23/vmart-api/api/pri/v1/user"
)

func V1(r fiber.Router) {
	r.Get("/ping", ping.SecurePingV1)
	r.Post("/user", user.CreateUser)
}
