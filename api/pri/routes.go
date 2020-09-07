package pri

import (
	"github.com/gofiber/fiber"
	v1 "github.com/vipin23/vmart-api/api/pri/v1"
)

func Routes(r fiber.Router) {
	v1Router := r.Group("/v1")
	v1.V1(v1Router)
}
