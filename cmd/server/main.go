package main

import (
	"url-shortener/configs"
	"url-shortener/internal/api/routes"
	"url-shortener/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := configs.NewSqliteConnection("../../sqlite.db")
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(models.Models...)

	r := gin.Default()
	r.GET("/", routes.Ping)
	routes.SetupShortenRoutes(r, db)
	routes.SetupRedirectRoutes(r, db)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
