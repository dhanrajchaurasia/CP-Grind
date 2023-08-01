package controllers

import (
	"html/template"
	"time"

	"github.com/dhanrajchaurasia/CP-GRIND/initializers"
	"github.com/dhanrajchaurasia/CP-GRIND/models"
	"github.com/gofiber/fiber/v2"
)

func HomePage(c *fiber.Ctx) error {
	msg := c.Cookies("message")
	cfProfile := template.HTML(c.Cookies("cfProfile"))
	c.ClearCookie("message")
	c.ClearCookie("cfProfile")
	data := fiber.Map{
		"User":       c.Cookies("username"),
		"CF_Profile": cfProfile,
		"Message":    msg,
	}
	return c.Render("home/index", data)
}
func LoginPage(c *fiber.Ctx) error {
	auth := c.Cookies("authorization_token")
	if initializers.IsValidToken(auth) {
		return c.Redirect("/")
	}
	return c.Render("home/login", fiber.Map{
		"Message": "Welcome to CP Grind!",
	})
}

func Signup(c *fiber.Ctx) error {
	fname := c.FormValue("fname")
	lname := c.FormValue("lname")
	username := c.FormValue("username")
	email := c.FormValue("email")
	pass := c.FormValue("password")
	cpass := c.FormValue("cpassword")
	if pass != cpass {
		return c.Render("home/login", fiber.Map{
			"Message": "Password didn't match!",
		})
	}
	user := models.User{
		FirstName:  fname,
		SecondName: lname,
		Username:   username,
		Email:      email,
		Password:   pass,
	}
	err := initializers.CreateNewUser(user)
	if err != nil {
		return c.Render("home/login", fiber.Map{
			"Message": err.Error(),
		})
	}
	return c.Render("home/login", fiber.Map{
		"Message": "Your account has been created successfully, Login to Continue!",
	})
}

func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	token, err := initializers.IsUserPresent(username, password)
	if err != nil {
		return c.Render("home/login", fiber.Map{
			"Message": err.Error(),
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:  "authorization_token",
		Value: token,
	})
	c.Cookie(&fiber.Cookie{
		Name:  "username",
		Value: username,
	})
	return c.Redirect("/")
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "authorization_token",
		Value:   "",
		Expires: time.Now(),
	})
	return c.Redirect("/login")
}

func NotFound(c *fiber.Ctx) error {
	return c.Render("home/404", fiber.Map{})
}
