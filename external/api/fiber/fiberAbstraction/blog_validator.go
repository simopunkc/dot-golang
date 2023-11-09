package fiberAbstraction

import (
	"dot-golang/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type BlogValidator interface {
	ValidateBodyNews(c *fiber.Ctx) (domain.News, error)
	ValidateBodyTopics(c *fiber.Ctx) (domain.Topics, error)
	ValidateBodyRefNewsTopics(c *fiber.Ctx) (domain.RefNewsTopics, error)
	ValidateParameterDecimal(c *fiber.Ctx, name string) (int, error)
	ValidateParameterIdDb(c *fiber.Ctx, name string) (int64, error)
}
