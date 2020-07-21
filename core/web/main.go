package web

import (
	_ "github.com/fjah/seachad/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Seachad API
// @version 1.0
// @description The API for Seachad, a free, open-source and customisable donation site.

// @contact.name GitHub
// @contact.url https://github.com/fjah/seachad/issues/new

// @license.name BSD-3-Clause

// @BasePath /api/v1

// Server is a Seachad web server.
type Server struct {
	Address     string
	Development bool
	// TODO: SSL/TLS

	engine *gin.Engine
}

// Init sets up the server.
func (s *Server) Init() {
	s.engine = gin.New()
	if !s.Development {
		gin.SetMode(gin.ReleaseMode)
	}

	url := ginSwagger.URL("/swagger/doc.json")
	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	s.setRoutes()
}

// Run runs the HTTP server.
func (s *Server) Run() error {
	return s.engine.Run(s.Address)
}
