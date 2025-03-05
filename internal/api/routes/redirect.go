package routes

import (
	"errors"
	"net/http"
	"url-shortener/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RedirectParams struct {
	Slug string `uri:"slug" binding:"required"`
}

func SetupRedirectRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/:slug", func(ctx *gin.Context) {
		var params RedirectParams
		if err := ctx.ShouldBindUri(&params); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		var url models.Url
		result := db.Where("hash = ?", params.Slug).First(&url)
		notFound := errors.Is(result.Error, gorm.ErrRecordNotFound)
		if notFound {
			ctx.JSON(http.StatusNotFound, gin.H{"msg": result.Error.Error()})
			return
		}
		ctx.Redirect(http.StatusPermanentRedirect, url.Permalink)
	})
}
