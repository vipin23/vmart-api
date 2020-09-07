package api_router

import (
	"os"
	"strings"

	"github.com/vipin23/vmart-api/api_server"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	"github.com/vipin23/vmart-api/api"
	"github.com/vipin23/vmart-api/middleware"
)

func SetupRouter() {
	api_server.DefaultServer.App = fiber.New()
	api_server.DefaultServer.App.Use(cors.New())
	api_server.DefaultServer.App.Use(middleware.Recover())

	// JWT Middleware
	api_server.DefaultServer.App.Use(jwtware.New(jwtware.Config{
		Filter: func(c *fiber.Ctx) bool {
			return !strings.Contains(c.Path(), "/pri/")
		},
		SigningKey: []byte(os.Getenv("API_SECRET")),
	}))

	api_server.DefaultServer.App.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})

	apiRoutes := api_server.DefaultServer.App.Group("/api")
	api.Routes(apiRoutes)
}

func Run(addr string) {
	api_server.DefaultServer.App.Listen(addr)
}
