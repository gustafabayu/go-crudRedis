package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gustafabayu/go-crudRedis/controller"
)

func NewRouter(router *fiber.App, novelController *controller.NovelController) *fiber.App {
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world")
	})

	router.Post("/novel", novelController.CreateNovel)
	router.Get("/novel/:id", novelController.GetNovelById)
	return router
}
