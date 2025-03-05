package routes

import (
	"net/http"
	"url-shortener/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupShortenRoutes(r *gin.Engine, db *gorm.DB) {
	shorten := r.Group("/api").Group("/shorten")
	{
		urlShortenerService := services.NewUrlShortenerService(db)
		shorten.POST("/", func(ctx *gin.Context) {
			var body services.ShortenRequestBody
			if err := ctx.ShouldBind(&body); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			items, err := urlShortenerService.GenerateUrls(ctx, &body)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			ctx.JSON(http.StatusOK, &services.ShortenRequestResponse{
				Urls: items,
			})
		})
	}
}
