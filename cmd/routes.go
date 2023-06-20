package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yagoernandes/album_league_api/cmd/domain/models"
	"github.com/yagoernandes/album_league_api/cmd/environment"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetRoutes(s *environment.Server) {
	s.App.Post("/prospects", func(c *fiber.Ctx) error {
		p := new(models.Prospect)

		if err := c.BodyParser(p); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		p.CreatedAt = time.Now()
		p.IPAddress = c.IP()

		insertionResult, err := s.Collection.InsertOne(c.Context(), p)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		c.Set("X-Insert-Id", insertionResult.InsertedID.(primitive.ObjectID).Hex())
		return c.JSON(p)
	})
}
