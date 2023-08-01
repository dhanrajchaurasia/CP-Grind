package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UserProfile(c *fiber.Ctx) error {
	return c.Render("profile/index", fiber.Map{})
}

func GetCFProfile(c *fiber.Ctx) error {
	url := fmt.Sprintf("https://codeforces.com/api/user.info?handles=%v", c.FormValue("cf-handle"))
	response, err := http.Get(url)
	htmlBody := "Could not get the data!"
	c.Cookie(&fiber.Cookie{Name: "message", Value: "", Expires: time.Now()})
	if err != nil {
		c.Cookie(&fiber.Cookie{
			Name:  "cfProfile",
			Value: htmlBody,
		})
		c.Cookie(&fiber.Cookie{
			Name:  "message",
			Value: err.Error(),
		})
		return c.Redirect("/")
	}
	defer response.Body.Close()
	jsonData, _ := io.ReadAll(response.Body)
	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		c.Cookie(&fiber.Cookie{
			Name:  "cfProfile",
			Value: htmlBody,
		})
		c.Cookie(&fiber.Cookie{
			Name:  "message",
			Value: err.Error(),
		})
		return c.Redirect("/")
	}
	result := data["result"].([]interface{})
	if len(result) > 0 {
		user := result[0].(map[string]interface{})
		rank := user["rank"].(string)
		rating := user["rating"].(float64)
		htmlBody = fmt.Sprintf("CF Rank: %v<br>CF Rating: %v", rank, int(rating))
	}
	c.Cookie(&fiber.Cookie{
		Name:  "cfProfile",
		Value: htmlBody,
	})
	return c.Redirect("/")
}
