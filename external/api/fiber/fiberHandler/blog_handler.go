package fiberHandler

import (
	"dot-golang/external/api/fiber/fiberAbstraction"
	"dot-golang/internal/abstraction"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type BlogHandler struct {
	blogService   abstraction.BlogService
	blogValidator fiberAbstraction.BlogValidator
	blogResponse  fiberAbstraction.BlogResponse
}

func NewBlogHandler(blogService abstraction.BlogService, blogValidator fiberAbstraction.BlogValidator, blogResponse fiberAbstraction.BlogResponse) *BlogHandler {
	return &BlogHandler{blogService, blogValidator, blogResponse}
}

func (bh BlogHandler) GetHomepage(c *fiber.Ctx) error {
	resp := "hello world"
	return bh.blogResponse.ResponseJsonOk(c, resp, 200)
}

func (bh BlogHandler) GetNews(c *fiber.Ctx) error {
	page, err := bh.blogValidator.ValidateParameterDecimal(c, "page")
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 404)
	}

	service, err := bh.blogService.GetNews(page)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}

	result, err := json.Marshal(service)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}
	return bh.blogResponse.ResponseJsonOk(c, result, 200)
}

func (bh BlogHandler) GetNewsDrafted(c *fiber.Ctx) error {
	page, err := bh.blogValidator.ValidateParameterDecimal(c, "page")
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 404)
	}

	service, err := bh.blogService.GetNewsDrafted(page)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}

	result, err := json.Marshal(service)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}
	return bh.blogResponse.ResponseJsonOk(c, result, 200)
}

func (bh BlogHandler) GetNewsPublished(c *fiber.Ctx) error {
	page, err := bh.blogValidator.ValidateParameterDecimal(c, "page")
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 404)
	}

	service, err := bh.blogService.GetNewsPublished(page)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}

	result, err := json.Marshal(service)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}
	return bh.blogResponse.ResponseJsonOk(c, result, 200)
}

func (bh BlogHandler) GetNewsDeleted(c *fiber.Ctx) error {
	page, err := bh.blogValidator.ValidateParameterDecimal(c, "page")
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 404)
	}

	service, err := bh.blogService.GetNewsDeleted(page)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}

	result, err := json.Marshal(service)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}
	return bh.blogResponse.ResponseJsonOk(c, result, 200)
}

func (bh BlogHandler) PostNews(c *fiber.Ctx) error {
	parseBody, err := bh.blogValidator.ValidateBodyNews(c)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}

	err = bh.blogService.PostNews(parseBody)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}
	return bh.blogResponse.ResponseJsonOk(c, nil, 200)
}

func (bh BlogHandler) GetSingleNews(c *fiber.Ctx) error {
	id, err := bh.blogValidator.ValidateParameterIdDb(c, "id")
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 404)
	}

	service, err := bh.blogService.GetSingleNews(id)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}

	result, err := json.Marshal(service)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}
	return bh.blogResponse.ResponseJsonOk(c, result, 200)
}

func (bh BlogHandler) PutSingleNews(c *fiber.Ctx) error {
	parseBody, err := bh.blogValidator.ValidateBodyNews(c)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}

	err = bh.blogService.PutSingleNews(parseBody)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}
	return bh.blogResponse.ResponseJsonOk(c, nil, 200)
}

func (bh BlogHandler) DeleteSingleNews(c *fiber.Ctx) error {
	id, err := bh.blogValidator.ValidateParameterIdDb(c, "id")
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 404)
	}

	err = bh.blogService.DeleteSingleNews(id)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}
	return bh.blogResponse.ResponseJsonOk(c, nil, 200)
}

func (bh BlogHandler) PatchSingleNewsStatusContent(c *fiber.Ctx) error {
	parseId, err := bh.blogValidator.ValidateParameterIdDb(c, "id")
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 404)
	}

	parseBody, err := bh.blogValidator.ValidateBodyNews(c)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}

	err = bh.blogService.PatchSingleNewsStatusContent(parseId, parseBody.StatusContent)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}
	return bh.blogResponse.ResponseJsonOk(c, nil, 200)
}

func (bh BlogHandler) GetTopics(c *fiber.Ctx) error {
	page, err := bh.blogValidator.ValidateParameterDecimal(c, "page")
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 404)
	}

	service, err := bh.blogService.GetTopics(page)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}

	result, err := json.Marshal(service)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}
	return bh.blogResponse.ResponseJsonOk(c, result, 200)
}

func (bh BlogHandler) GetSingleTopics(c *fiber.Ctx) error {
	id, err := bh.blogValidator.ValidateParameterIdDb(c, "idTopic")
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 404)
	}

	service, err := bh.blogService.GetSingleTopics(id)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}

	result, err := json.Marshal(service)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}
	return bh.blogResponse.ResponseJsonOk(c, result, 200)
}

func (bh BlogHandler) PostTopics(c *fiber.Ctx) error {
	parseBody, err := bh.blogValidator.ValidateBodyTopics(c)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}

	err = bh.blogService.PostTopics(parseBody)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}
	return bh.blogResponse.ResponseJsonOk(c, nil, 200)
}

func (bh BlogHandler) PatchSingleTopicsCategoryName(c *fiber.Ctx) error {
	id, err := bh.blogValidator.ValidateParameterIdDb(c, "idTopic")
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 404)
	}

	parseBody, err := bh.blogValidator.ValidateBodyTopics(c)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}

	err = bh.blogService.PatchSingleTopicsCategoryName(id, parseBody.CategoryName)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}
	return bh.blogResponse.ResponseJsonOk(c, nil, 200)
}

func (bh BlogHandler) GetSingleTopicsNews(c *fiber.Ctx) error {
	id, err := bh.blogValidator.ValidateParameterIdDb(c, "idTopic")
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 404)
	}

	page, err := bh.blogValidator.ValidateParameterDecimal(c, "page")
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 404)
	}

	service, err := bh.blogService.GetSingleTopicsNews(id, page)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}

	result, err := json.Marshal(service)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}
	return bh.blogResponse.ResponseJsonOk(c, result, 200)
}

func (bh BlogHandler) PostRefNewsTopics(c *fiber.Ctx) error {
	parseBody, err := bh.blogValidator.ValidateBodyRefNewsTopics(c)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}

	err = bh.blogService.PostRefNewsTopics(parseBody)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}
	return bh.blogResponse.ResponseJsonOk(c, nil, 200)
}

func (bh BlogHandler) DeleteRefNewsTopics(c *fiber.Ctx) error {
	parseBody, err := bh.blogValidator.ValidateBodyRefNewsTopics(c)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}

	err = bh.blogService.DeleteRefNewsTopics(parseBody)
	if err != nil {
		return bh.blogResponse.ResponseJsonError(c, []error{err}, 400)
	}
	return bh.blogResponse.ResponseJsonOk(c, nil, 200)
}
