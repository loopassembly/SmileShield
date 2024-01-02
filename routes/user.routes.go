// user.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"cypher-server/controllers"
	"cypher-server/middleware"
)

func SetupUserRoutes(router fiber.Router) {
	router.Get("/me", middleware.DeserializeUser, controllers.GetMe)
}