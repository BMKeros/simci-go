package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	address string
	engine  *gin.Engine
}

func New(host string, port uint) Server {
	srv := Server{
		address: fmt.Sprintf("%s:#%d", host, port),
		engine:  gin.New(),
	}

	return srv
}

func (s *Server) Run() error {
	return s.engine.Run(s.address)
}

func (s *Server) registerRoutes() {

}
