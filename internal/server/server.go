package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	r  *gin.Engine
	db *gorm.DB
}

func NewServer(r *gin.Engine, db *gorm.DB) *Server {
	return &Server{
		r:  r,
		db: db,
	}
}

func (s *Server) Run() error {
	s.r = gin.Default()

	if err := s.MapHandlers(); err != nil {
		return err
	}

	if err := s.r.Run(":8080"); err != nil {
		return err
	}

	return nil
}
