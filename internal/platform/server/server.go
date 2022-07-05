package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	domain "simci-go/internal"
)

type Server struct {
	address string
	engine  *gin.Engine

	userRepository domain.UserRepository
}

func New(host string, port uint, userRepository domain.UserRepository) Server {
	srv := Server{
		address:        fmt.Sprintf("%s:%d", host, port),
		engine:         gin.New(),
		userRepository: userRepository,
	}

	return srv
}

func (s *Server) Run() error {
	return s.engine.Run(s.address)
}

func (s *Server) registerRoutes() {

}
