package server

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)

type Server struct {
	r      *gin.Engine
	db     *gorm.DB
	client *mongo.Client
}

func NewServer(r *gin.Engine, db *gorm.DB, client *mongo.Client) *Server {
	return &Server{
		r:      r,
		db:     db,
		client: client,
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
