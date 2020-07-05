package routes

import (
	"time"

	"github.com/Hadermite/invenmind/database"
	"github.com/Hadermite/invenmind/middleware"
	"github.com/Hadermite/invenmind/model"
	"github.com/Hadermite/invenmind/util"
	"github.com/gofiber/fiber"
)

type createLocationBody struct {
	Name string
}

// Location - Location routes
func Location(router *fiber.Group) {
	router.Use(middleware.AuthRequired)

	router.Get("", func(c *fiber.Ctx) {
		var user = c.Locals("user").(model.User)
		var locations []model.Location
		database.Connection.
			Table("locations").
			Joins("INNER JOIN location_users ON location_users.location_id = locations.id").
			Where("location_users.user_id = ?", user.ID).
			Find(&locations)
		c.JSON(&locations)
	})

	router.Post("", func(c *fiber.Ctx) {
		var user = c.Locals("user").(model.User)
		var body createLocationBody
		var error = c.BodyParser(&body)
		if error != nil || util.IsAnyStringEmpty([]string{body.Name}) {
			c.SendStatus(fiber.StatusBadRequest)
			return
		}
		var addedDate = time.Now()

		var location model.Location
		location.Name = body.Name
		location.AddedDate = addedDate
		database.Connection.Create(&location)

		var locationUser model.LocationUser
		locationUser.LocationID = location.ID
		locationUser.UserID = user.ID
		locationUser.AddedDate = addedDate
		database.Connection.Create(&locationUser)

		c.Status(fiber.StatusCreated)
	})
}
