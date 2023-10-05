package controllers

import "github.com/gofiber/fiber/v2"

func GrindPage(c *fiber.Ctx) error {
	return c.Render("home/grind", fiber.Map{})
}

