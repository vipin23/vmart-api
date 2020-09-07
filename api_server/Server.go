package api_server

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Server struct {
	DB  *gorm.DB
	App *fiber.App
}

var DefaultServer = &Server{}
