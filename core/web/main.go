package web

import "github.com/gin-gonic/gin"

// Server is a Seachad web server.
type Server struct {
	Address string
	// TODO: SSL/TLS

	engine *gin.Engine
}

// Init sets up the server.
func (s *Server) Init() {
	s.engine = gin.New()
}

// Run runs the HTTP server.
func (s *Server) Run() error {
	return s.engine.Run(s.Address)
}
