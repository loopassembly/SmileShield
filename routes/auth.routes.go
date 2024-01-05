// auth.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"cypher-server/controllers"
	"cypher-server/middleware"
)
// all auth routes including oauth
func SetupAuthRoutes(router fiber.Router) {
	router.Post("/register", controllers.SignUpUser)
	router.Post("/login", controllers.SignInUser)
	router.Get("/logout", middleware.DeserializeUser, controllers.LogoutUser)
	router.Get("/verifyemail/:verificationCode", controllers.VerifyEmail)
	router.Post("/forgotpassword", controllers.ForgotPassword)
	router.Patch("/resetpassword/:resetToken",controllers.ResetPassword)
	router.Get("/sessions/oauth/google", controllers.GoogleOAuth)
	router.Get("/sessions/oauth/github", controllers.GitHubOAuth)
}
