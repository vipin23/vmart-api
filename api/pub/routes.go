package pub

import (
	"github.com/gofiber/fiber"
	v1 "github.com/vipin23/vmart-api/api/pub/v1"
	v2 "github.com/vipin23/vmart-api/api/pub/v2"
)

func Routes(r fiber.Router) {
	v1Router := r.Group("/v1")
	v1.V1(v1Router)
	v2Router := r.Group("/v2")
	v2.V2(v2Router)
}
