package fiberValidator

import (
	"dot-golang/internal/abstraction"
	"dot-golang/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type BlogValidator struct {
	util abstraction.BlogUtil
}

func NewBlogValidator(util abstraction.BlogUtil) *BlogValidator {
	return &BlogValidator{util}
}

func (bv BlogValidator) ValidateBodyNews(c *fiber.Ctx) (domain.News, error) {
	var parseBody domain.News
	err := c.BodyParser(&parseBody)
	return parseBody, err
}

func (bv BlogValidator) ValidateBodyTopics(c *fiber.Ctx) (domain.Topics, error) {
	var parseBody domain.Topics
	err := c.BodyParser(&parseBody)
	return parseBody, err
}

func (bv BlogValidator) ValidateBodyRefNewsTopics(c *fiber.Ctx) (domain.RefNewsTopics, error) {
	var parseBody domain.RefNewsTopics
	err := c.BodyParser(&parseBody)
	return parseBody, err
}

func (bv BlogValidator) ValidateParameterDecimal(c *fiber.Ctx, name string) (int, error) {
	param := c.Params(name)
	if param == "" {
		param = "1"
	}
	result, err := bv.util.StringToInt(param)
	return result, err
}

func (bv BlogValidator) ValidateParameterIdDb(c *fiber.Ctx, name string) (int64, error) {
	paramId := c.Params(name)
	id, err := bv.util.StringToInt64(paramId)
	return id, err
}
