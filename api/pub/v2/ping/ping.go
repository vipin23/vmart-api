package ping

import (
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/vipin23/vmart-api/rest_models"
)

func PingV2(c *fiber.Ctx) {
	err := c.SendFile("file-does-not-exist")

	if err != nil {
		res := rest_models.BaseResponse{Code: http.StatusInsufficientStorage}
		res.Error(c, err)
		// panic("This panic is catched by the ErrorHandler")
	}
}
