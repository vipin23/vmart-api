package user

import (
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/vipin23/vmart-api/models"
	"github.com/vipin23/vmart-api/rest_models"
)

func CreateUser(c *fiber.Ctx) {
	user := models.User{}
	err := c.BodyParser(&user)
	if err != nil {
		res := rest_models.BaseResponse{Code: http.StatusUnprocessableEntity}
		res.Error(c, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		res := rest_models.BaseResponse{Code: http.StatusUnprocessableEntity}
		res.Error(c, err)
		return
	}
	var savedUser *models.User
	savedUser, err = user.SaveUser()
	if err != nil {
		res := rest_models.BaseResponse{Code: http.StatusNotAcceptable}
		res.Error(c, err)
		return
	}
	res := rest_models.BaseResponse{}
	res.Created(c, savedUser)
}
