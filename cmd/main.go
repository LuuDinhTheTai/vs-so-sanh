package main

import (
	"log/slog"
	"vs-so-sanh/infrastructures/configuration"
	"vs-so-sanh/infrastructures/database"
	server2 "vs-so-sanh/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configuration.LoadEnv()

	client, ctx, cancel, errDB := database.Connect(cfg.Database.URI)
	if errDB != nil {
		slog.Error("Database connect err:", errDB)
		return
	}
	errPing := database.Ping(client, ctx)
	if errPing != nil {
		slog.Error("Database ping err:", errPing)
		return
	}
	defer database.Close(client, ctx, cancel)

	r := gin.Default()

	server := server2.NewServer(r, nil)

	slog.Info("Server start at http://" + cfg.Server.Host + ":" + cfg.Server.Port)
	err := server.Run()
	if err != nil {
		slog.Error("Server start err:", err)
		return
	}
}
