package server

import (
	"log"

	authHandler "github.com/DiegoUrrego4/Wheeper/internal/auth/infrastructure/handler"
	"github.com/labstack/echo/v4"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	e := echo.New()

	// Rutas
	login := e.Group("/v1/login")
	login.GET("", authHandler.Ping)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Error running the server: %v", err)
	}
}
