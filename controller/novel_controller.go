package controller

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gustafabayu/go-crudRedis/domain"
	"github.com/gustafabayu/go-crudRedis/model"
)

type NovelController struct {
	novelUseCase domain.NovelUseCase
}

func NewNovelController(novelUseCase domain.NovelUseCase) *NovelController {
	return &NovelController{novelUseCase: novelUseCase}
}

func (n *NovelController) CreateNovel(ctx *fiber.Ctx) error {
	var novelRequest model.Novel
	var response model.Response

	if err := ctx.BodyParser(&novelRequest); err != nil {
		response = model.Response{StatusCode: http.StatusBadRequest, Message: err.Error()}
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if novelRequest.Author == "" || novelRequest.Name == "" || novelRequest.Description == "" {
		response = model.Response{StatusCode: http.StatusBadRequest, Message: "Request cannot be empty"}
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := n.novelUseCase.CreateNovel(novelRequest); err != nil {
		response = model.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response = model.Response{StatusCode: http.StatusOK, Message: "Insert data success"}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (n *NovelController) GetNovelById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid id (cannot be null/0)"})
	}

	novel, err := n.novelUseCase.GetNovelById(idInt)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var res model.Response
	if novel.Name != "" {
		res = model.Response{StatusCode: http.StatusOK, Message: "Get Data By Id Success", Data: novel}
	} else {
		res = model.Response{StatusCode: http.StatusOK, Message: "Get Data By Id Success (Null Data)"}
	}

	return ctx.Status(http.StatusOK).JSON(res)
}
