package api

import (
	"github.com/sanLimbu/hotel-reservation/types"

	"github.com/gofiber/fiber/v2"
)

func HandleUser(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "san",
		LastName:  "lim",
	}
	return c.JSON(user)
}
