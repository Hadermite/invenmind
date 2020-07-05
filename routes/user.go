package routes

import (
	"strings"
	"time"

	"github.com/Hadermite/invenmind/database"
	"github.com/Hadermite/invenmind/middleware"
	"github.com/Hadermite/invenmind/model"
	"github.com/Hadermite/invenmind/util"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type loginBody struct {
	Email      string
	Password   string
	DeviceName string
}

type registerBody struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}

// User - User routes
func User(router *fiber.Group) {
	router.Get("", middleware.AuthRequired, func(c *fiber.Ctx) {
		c.JSON(c.Locals("user"))
	})

	router.Post("/login", func(c *fiber.Ctx) {
		var body loginBody
		var error = c.BodyParser(&body)
		if error != nil {
			c.JSON(struct{ Error string }{Error: error.Error()})
			return
		}
		if util.IsAnyStringEmpty([]string{
			body.Email,
			body.Password,
			body.DeviceName,
		}) {
			c.SendStatus(fiber.StatusBadRequest)
			return
		}

		var email = strings.ToLower(strings.TrimSpace(body.Email))
		var user model.User
		var result = database.Connection.Where(&model.User{Email: email}).First(&user)
		if result.RecordNotFound() || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
			c.Status(fiber.StatusBadRequest).JSON(struct{ Error string }{Error: "Invalid credentials."})
			return
		}

		var token model.AuthToken
		token.UserID = user.ID
		token.Token = uuid.New().String() // TODO: Implement a more cryptographically strong token
		token.CreateDate = time.Now()
		token.DeviceName = body.DeviceName
		database.Connection.Create(&token)

		c.JSON(struct {
			User  model.User
			Token string
		}{
			User:  user,
			Token: token.Token,
		})
	})

	router.Post("/register", func(c *fiber.Ctx) {
		var body registerBody
		var error = c.BodyParser(&body)
		if error != nil {
			c.JSON(struct{ Error string }{Error: error.Error()})
			return
		}
		if util.IsAnyStringEmpty([]string{
			body.Email,
			body.FirstName,
			body.LastName,
			body.Password,
		}) {
			c.SendStatus(fiber.StatusBadRequest)
			return
		}

		var email = strings.ToLower(strings.TrimSpace(body.Email))
		var existingUser model.User
		var result = database.Connection.Where(&model.User{Email: email}).First(&existingUser)
		if !result.RecordNotFound() {
			c.Status(fiber.StatusBadRequest).JSON(struct{ Error string }{Error: "That email address is already associated with an account."})
			return
		}

		passwordHash, error := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
		if error != nil {
			c.SendStatus(fiber.StatusInternalServerError)
			return
		}

		var user model.User
		user.Email = email
		user.FirstName = strings.TrimSpace(body.FirstName)
		user.LastName = strings.TrimSpace(body.LastName)
		user.Password = string(passwordHash)
		user.RegisterDate = time.Now()
		database.Connection.Create(&user)

		c.JSON(&user)
	})
}
