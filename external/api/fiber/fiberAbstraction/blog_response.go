package fiberAbstraction

import (
	"github.com/gofiber/fiber/v2"
)

type BlogResponse interface {
	ResponseJsonOk(c *fiber.Ctx, raw interface{}, statusCode int) error
	ResponseJsonError(c *fiber.Ctx, err []error, statusCode int) error
}
