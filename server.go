package pokemon_be

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"net/http"
	"pokemon-be/internal/api"
	v1 "pokemon-be/internal/api/v1"
	"pokemon-be/internal/config"
	"pokemon-be/internal/middleware"
	"pokemon-be/internal/repository"
	"pokemon-be/internal/service"
)

type Server struct {
	gin *gin.Engine
	srv *http.Server
}

func NewServer(cfg config.Config, db *bun.DB) *Server {

	router := gin.New()

	router.Use(middleware.RequestLogger())

	router.Use(gin.Recovery())

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(cfg, userRepo)

	userController := v1.NewUserController(cfg, userService)

	registerHandlers(router, &api.HealthCheck{}, userController)

	return &Server{
		gin: router,
	}
}

func (s *Server) Start(addr string) error {
	srv := &http.Server{
		Addr:    addr,
		Handler: s.gin,
	}

	s.srv = srv

	return srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {

	return s.srv.Shutdown(ctx)
}

func registerHandlers(g *gin.Engine, handlers ...api.Handler) {
	for _, handler := range handlers {
		handler.AddRoutes(g)
	}
}
