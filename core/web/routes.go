package web

import (
	"github.com/fjah/seachad/core/web/handlers"
)

func (s *Server) setRoutes() {
	v1 := s.engine.Group("/api/v1")
	{
		v1.GET("/test", handlers.Test)
	}
}
