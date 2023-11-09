package fiberAbstraction

import "github.com/gofiber/fiber/v2"

type BlogHandler interface {
	GetHomepage(c *fiber.Ctx) error
	GetNews(c *fiber.Ctx) error
	GetSingleNews(c *fiber.Ctx) error
	PutSingleNews(c *fiber.Ctx) error
	DeleteSingleNews(c *fiber.Ctx) error
	GetNewsDrafted(c *fiber.Ctx) error
	GetNewsPublished(c *fiber.Ctx) error
	GetNewsDeleted(c *fiber.Ctx) error
	PostNews(c *fiber.Ctx) error
	PatchSingleNewsStatusContent(c *fiber.Ctx) error
	GetTopics(c *fiber.Ctx) error
	GetSingleTopics(c *fiber.Ctx) error
	PostTopics(c *fiber.Ctx) error
	PatchSingleTopicsCategoryName(c *fiber.Ctx) error
	GetSingleTopicsNews(c *fiber.Ctx) error
	PostRefNewsTopics(c *fiber.Ctx) error
	DeleteRefNewsTopics(c *fiber.Ctx) error
}
