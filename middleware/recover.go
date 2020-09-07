package middleware

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/vipin23/vmart-api/rest_models"
)

// Recover will recover from panics and calls the ErrorHandler
func Recover() fiber.Handler {
	return func(ctx *fiber.Ctx) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				res := rest_models.BaseResponse{}
				res.Error(ctx, err)
				return
			}
		}()
		ctx.Next()
	}
}
