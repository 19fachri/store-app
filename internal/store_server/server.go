package store_server

import (
	"log/slog"

	"github.com/19fachri/store-app/internal/store_server/config"
	"github.com/19fachri/store-app/internal/store_server/datalayer"
	"github.com/19fachri/store-app/internal/store_server/datalayer/actions"
	"github.com/19fachri/store-app/internal/store_server/handlers/user_handler"
	"github.com/19fachri/store-app/internal/store_server/services/user_service"
	"github.com/gin-gonic/gin"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() {
	logger := slog.Default()

	config.Load("./configs/store_server")
	if config.Get().Platform.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	datalayer.InitDB(logger)
	db := datalayer.GetDB()

	action := actions.New(db)

	userService := user_service.New(logger, action)

	userHandler := user_handler.New(logger, userService)

	router := gin.Default()
	apiV1 := router.Group("/api/v1")

	user_handler.RegisterAuthRoute(apiV1.Group(""), userHandler)

	// secure := apiV1.Group("/secure")
	// {
	// 	user_handler.RegisterUserRoute(secure.Group(""), userHandler)
	// }

	router.Run(config.Get().Platform.Port)
}
