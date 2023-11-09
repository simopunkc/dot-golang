package fiberServer

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"github.com/joho/godotenv"

	"dot-golang/external/api/fiber/fiberHandler"
	"dot-golang/external/api/fiber/fiberMiddleware"
	"dot-golang/external/api/fiber/fiberResponse"
	"dot-golang/external/api/fiber/fiberRouter"
	"dot-golang/external/api/fiber/fiberValidator"
	"dot-golang/external/database/gorm/mysql"
	"dot-golang/external/database/gorm/repository"
	"dot-golang/external/database/redis"
	"dot-golang/internal/constant"
	"dot-golang/internal/domain"
	"dot-golang/internal/pkg/util"
	"dot-golang/internal/service"
)

func ApiServer() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	backendHost := ":" + os.Getenv("BACKEND_PORT")

	config := fiber.Config{
		ServerHeader:  "Dot",
		StrictRouting: true,
		CaseSensitive: true,
	}

	app := fiber.New(config)

	eventBus := redis.NewEventBus()
	eventChan := make(chan domain.Event)
	eventBus.Subscribe(constant.EVENT_DELETE_SINGLE_NEWS, eventChan)

	rdb := redis.NewConnectionRedis()
	blogCache := redis.NewBlogCache(rdb)
	db := mysql.NewConnectionMysql()
	blogRepository := repository.NewBlogRepository(db)
	blogEvent := redis.NewEventService(eventBus)
	blogUtil := util.NewBlogUtil()
	go redis.EventConsumer(blogCache, blogUtil, eventChan)
	blogService := service.NewBlogService(blogCache, blogRepository, blogEvent, blogUtil)
	blogResponse := fiberResponse.NewBlogResponse(blogUtil)
	blogValidator := fiberValidator.NewBlogValidator(blogUtil)
	blogHandler := fiberHandler.NewBlogHandler(blogService, blogValidator, blogResponse)
	blogRouter := fiberRouter.NewBlogRouter(blogHandler)

	blogGroup := app.Group("/blog", fiberMiddleware.SetSecurityHeader)
	blogRouter.BlogRouter(blogGroup)

	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join([]string{
			os.Getenv("FRONTEND_PROTOCOL") + os.Getenv("FRONTEND_HOST"),
		}, ", "),
	}))

	app.Use(limiter.New(limiter.Config{
		Max:               30,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	fmt.Println("server running")

	addr := flag.String("addr", backendHost, "http service address")
	flag.Parse()
	log.Fatal(app.Listen(*addr))
}
