package routes

import (
	"Ga1ors/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Public Routes
	app.Post("/api/register", controllers.Register)
	app.Post("/api/verify", controllers.Verification)
	app.Post("/api/login", controllers.Login)

	// Middleware uses for making sure that a valid user is logged in
	// Otherwise, other functions are not available.
	//app.Use(middleware.IsAuthenticated)

	// Private Routes
	app.Put("/api/users/info", controllers.UpdateInfo)
	app.Put("/api/users/password", controllers.UpdatePassword)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/messages", controllers.AllMsgs)
	app.Post("/api/message", controllers.Message)
	app.Post("/api/messages", controllers.CreateMsg)
	app.Delete("/api/messages:id", controllers.DeleteMsgs)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	//app.Post("/api/users", controllers.sendEmail)   //Routes for later verification use
	//app.Get("/api/users/:id", controllers.sendEmail)
}
