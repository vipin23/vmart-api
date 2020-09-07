package api

import (
	"github.com/gofiber/fiber"
	"github.com/vipin23/vmart-api/api/pri"
	"github.com/vipin23/vmart-api/api/pub"
)

func Routes(r fiber.Router) {
	pubRoutes := r.Group("/pub")
	pub.Routes(pubRoutes)
	priRoutes := r.Group("/pri")
	pri.Routes(priRoutes)
}
