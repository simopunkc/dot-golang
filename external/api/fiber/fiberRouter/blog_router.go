package fiberRouter

import (
	"dot-golang/external/api/fiber/fiberAbstraction"

	"github.com/gofiber/fiber/v2"
)

type BlogRouter struct {
	blogHandler fiberAbstraction.BlogHandler
}

func NewBlogRouter(blogHandler fiberAbstraction.BlogHandler) *BlogRouter {
	return &BlogRouter{blogHandler}
}

func (br BlogRouter) BlogRouter(app fiber.Router) {
	app.Get("/", br.blogHandler.GetHomepage)

	singleNew := "/news/:id"
	app.Get(singleNew, br.blogHandler.GetSingleNews)
	app.Put(singleNew, br.blogHandler.PutSingleNews)
	app.Delete(singleNew, br.blogHandler.DeleteSingleNews)

	app.Get("/news/page/:page", br.blogHandler.GetNews)
	app.Get("/news/drafted/page/:page", br.blogHandler.GetNewsDrafted)
	app.Get("/news/published/page/:page", br.blogHandler.GetNewsPublished)
	app.Get("/news/deleted/page/:page", br.blogHandler.GetNewsDeleted)
	app.Post("/news", br.blogHandler.PostNews)
	app.Patch("/news/:id/status_content", br.blogHandler.PatchSingleNewsStatusContent)
	app.Get("/topics/page/:page", br.blogHandler.GetTopics)
	app.Get("/topics/:idTopic", br.blogHandler.GetSingleTopics)
	app.Post("/topics", br.blogHandler.PostTopics)
	app.Patch("/topics/:idTopic/category_name", br.blogHandler.PatchSingleTopicsCategoryName)
	app.Get("/topics/:idTopic/news/page/:page", br.blogHandler.GetSingleTopicsNews)
	app.Post("/ref/news/topics", br.blogHandler.PostRefNewsTopics)
	app.Delete("/ref/news/topics", br.blogHandler.DeleteRefNewsTopics)
}
