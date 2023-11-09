package fiberResponse

import (
	"dot-golang/internal/abstraction"
	"dot-golang/internal/domain"

	"github.com/gofiber/fiber/v2"
)

const (
	HEADER_CONTENT_TYPE = "Content-Type"
	HEADER_JSON         = "application/json"
	STATUS_RESPONSE_OK  = 200
	RESPONSE_OK         = "OK"
)

type BlogResponse struct {
	util abstraction.BlogUtil
}

func NewBlogResponse(util abstraction.BlogUtil) *BlogResponse {
	return &BlogResponse{util}
}

func (br BlogResponse) ResponseJsonOk(c *fiber.Ctx, raw interface{}, statusCode int) error {
	if raw == nil {
		raw = RESPONSE_OK
	}
	c.Status(statusCode)
	c.Set(HEADER_CONTENT_TYPE, HEADER_JSON)
	success := domain.ResponseHttpSuccess{
		StatusCode: STATUS_RESPONSE_OK,
		Data:       raw,
	}
	resp := br.util.ToJson(success)
	return c.Send(resp)
}

func (br BlogResponse) ResponseJsonError(c *fiber.Ctx, err []error, statusCode int) error {
	c.Status(statusCode)
	c.Set(HEADER_CONTENT_TYPE, HEADER_JSON)
	arr := []string{}
	for _, v := range err {
		arr = append(arr, v.Error())
	}
	failed := domain.ResponseHttpError{
		StatusCode: statusCode,
		Error:      arr,
	}
	resp := br.util.ToJson(failed)
	return c.Send(resp)
}
