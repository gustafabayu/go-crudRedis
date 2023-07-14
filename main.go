package main

import (
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gustafabayu/go-crudRedis/config"
	"github.com/gustafabayu/go-crudRedis/controller"
	"github.com/gustafabayu/go-crudRedis/database"
	"github.com/gustafabayu/go-crudRedis/model"
	"github.com/gustafabayu/go-crudRedis/repo"
	"github.com/gustafabayu/go-crudRedis/router"
	"github.com/gustafabayu/go-crudRedis/usecase"
	"gorm.io/gorm"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load envi variable", err)
	}

	db := database.ConnectionMySqlDb(&loadConfig)
	db.AutoMigrate(&model.Novel{})

	rdb := database.ConnectionRedisDb(&loadConfig)
	startServer(db, rdb)
}

func startServer(db *gorm.DB, rdb *redis.Client) {
	app := fiber.New()
	novelRepo := repo.NewNovelRepo(db, rdb)
	novelUseCase := usecase.NewNovelUseCase(novelRepo)
	novelController := controller.NewNovelController(novelUseCase)
	routes := router.NewRouter(app, novelController)
	err := routes.Listen(":3400")
	if err != nil {
		panic(err)
	}
}
