package main

import (
	"log"
	"os"

	"github.com/dhanrajchaurasia/CP-GRIND/controllers"
	"github.com/dhanrajchaurasia/CP-GRIND/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/dhanrajchaurasia/CP-GRIND/initializers"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectionToDB()
	initializers.SyncDB()
}

var Engine = html.New("./views", ".html")
var App = fiber.New(fiber.Config{Views: Engine})

func Routes() {

	// Get Requests
	App.Get("/", controllers.HomePage)
	App.Get("/404", controllers.NotFound)
	App.Get("/login", controllers.LoginPage)
	App.Get("/grind", controllers.GrindPage)

	// Post Requests
	App.Post("/signup", controllers.Signup)
	App.Post("/login", controllers.Login)
	App.Post("/logout", controllers.Logout)
	App.Post("/cfProfile", controllers.GetCFProfile)

	// 404 Page
	App.Use(func(c *fiber.Ctx) error {
		return c.SendString("Page Not Found")
	})
}

func main() {

	// Configure App
	App.Static("/", "./public")

	// Middleware
	App.Use(middleware.AuthMiddleware)
	// Routes
	Routes()

	// Server
	log.Fatal(App.Listen(":" + os.Getenv("PORT")))
}
