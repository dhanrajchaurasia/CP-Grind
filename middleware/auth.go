package middleware

import (
	"github.com/dhanrajchaurasia/CP-GRIND/initializers"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Cookies("authorization_token")
	// Allow access to the login and signup routes without authentication
	if c.Path() == "/login" || c.Path() == "/signup" {
		return c.Next()
	}

	// Check if the token is missing or invalid
	if tokenString == "" || !initializers.IsValidToken(tokenString) {
		// Redirect the user to the login page with a status of 401 Unauthorized
		return c.Redirect("/login")
	}
	if c.Path() == "/favicon.ico" || c.Path() == "/404" {
		return c.Redirect("/404")
	}
	// Continue to the next middleware or route handler
	return c.Next()
}
