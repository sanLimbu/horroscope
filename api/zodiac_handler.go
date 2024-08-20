package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sanLimbu/horroscope/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ZodiacHandler struct {
	store *db.Store
}

func NewZodiacHandler(store *db.Store) *ZodiacHandler {
	return &ZodiacHandler{
		store: store,
	}
}

func (h *ZodiacHandler) HandleGetHoroscopes(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}

	filter := bson.M{"zodiacID": oid}
	zodiacs, err := h.store.Horoscope.GetHoroscopes(c.Context(), filter)
	if err != nil {
		return ErrNotResourceNotFound("zodiacs")
	}

	return c.JSON(zodiacs)
}

func (h *ZodiacHandler) HandleGetZodiac(c *fiber.Ctx) error {

	id := c.Params("id")
	zodiac, err := h.store.Zodiac.GetZodiacByID(c.Context(), id)
	if err != nil {
		return ErrNotResourceNotFound("zodiac")
	}

	return c.JSON(zodiac)

}

type ResourceResp struct {
	Results int `json:"results"`
	Data    any `json:"data"`
	Page    int `json:"page"`
}

type ZodiacQueryParams struct {
	db.Pagination
}
