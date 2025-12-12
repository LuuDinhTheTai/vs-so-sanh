package main

import (
	"log/slog"
	server2 "vs-so-sanh/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	server := server2.NewServer(r, nil)

	slog.Info("Server start at http://localhost:8080")
	server.Run()
}
