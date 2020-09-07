package rest_models

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber"
)

type BaseResponse struct {
	Msg   string        `json:"msg"`
	Code  int           `json:"code"`
	Data  []interface{} `json:"data"`
	Total int           `json:"total"`
	Size  int           `json:"size"`
	Page  int           `json:"page"`
}

func (b BaseResponse) Error(c *fiber.Ctx, err error) {
	if b.Code == 0 {
		b.Code = http.StatusInternalServerError
	}
	if b.Msg == "" {
		b.Msg = err.Error()
	}
	c.Status(b.Code)
	c.JSON(b)
}

func (b BaseResponse) Created(c *fiber.Ctx, data interface{}) {
	if data == nil {
		b.Error(c, errors.New("No Data"))
		return
	}
	b.Data = []interface{}{data}
	b.Code = http.StatusCreated
	b.Msg = "New data created"
	b.Total = 1
	b.Size = 1
	b.Page = 1
	c.Status(b.Code)
	c.JSON(b)
}

func (b BaseResponse) SendMsg(c *fiber.Ctx, msg string) {
	b.Msg = msg
	b.Code = http.StatusOK
	b.Data = []interface{}{}
	b.Total = 0
	b.Size = 0
	b.Page = 1
	c.Status(b.Code)
	c.JSON(b)
}
