package middleware

import (
	"strings"

	"github.com/Hadermite/invenmind/database"
	"github.com/Hadermite/invenmind/model"

	"github.com/gofiber/fiber"
)

// ValidateAuth - Middleware to validate authentication and add the user to context locals
func ValidateAuth(c *fiber.Ctx) {
	var header = strings.TrimSpace(c.Get(fiber.HeaderAuthorization))
	if len(header) == 0 {
		c.Next()
		return
	}

	var parts = strings.Split(header, " ")
	if len(parts) != 2 {
		c.Next()
		return
	}
	if parts[0] != "Bearer" {
		c.Next()
		return
	}

	var token model.AuthToken
	var result = database.Connection.Where(&model.AuthToken{Token: parts[1]}).First(&token)
	if result.RecordNotFound() {
		c.Next()
		return
	}

	var user model.User
	result = database.Connection.First(&user, token.UserID)
	if result.RecordNotFound() {
		c.Next()
		return
	}

	c.Locals("user", user)
	c.Next()
}

// AuthRequired - Middleware to enforce an authenticated user
func AuthRequired(c *fiber.Ctx) {
	if c.Locals("user") == nil {
		c.Status(401).Send("")
	} else {
		c.Next()
	}
}
