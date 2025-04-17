package services

import (
	"sync"
	"time"
	"url-shortener/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type ShortenRequestBody struct {
	Urls []string `form:"urls" binding:"required,dive,url"`
}

type ShortenRequestResponse struct {
	Urls []*models.Url `json:"urls"`
}

type UrlShortenerService struct {
	db *gorm.DB
}

func NewUrlShortenerService(db *gorm.DB) *UrlShortenerService {
	return &UrlShortenerService{db: db}
}

func (this *UrlShortenerService) GenerateUrls(c *gin.Context, body *ShortenRequestBody) ([]*models.Url, error) {
	items := this.processUrlInternal(body.Urls)
	result := this.db.Create(items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

func (this *UrlShortenerService) processUrlInternal(urls []string) []*models.Url {
	out := make(chan *models.Url, len(urls))
	// wg := &sync.WaitGroup{}
	var wg *sync.WaitGroup
	wg.Add(len(urls))
	for _, permalink := range urls {
		go func(p string) { //imagine is this a heavy computed task
			defer wg.Done()
			id := this.generateUniqueID()
			shortURL := this.encodeBase62(id)
			newUrl := &models.Url{
				Permalink: permalink,
				Hash:      shortURL,
			}
			out <- newUrl
		}(permalink)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	result := make([]*models.Url, 0, len(urls))
	for val := range out {
		result = append(result, val)
	}
	return result
}

func (this *UrlShortenerService) encodeBase62(num int64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	var base int64 = int64(len(base62Chars))
	result := ""
	for num > 0 {
		result = string(base62Chars[num%base]) + result
		num /= base
	}
	return result
}

func (this *UrlShortenerService) generateUniqueID() int64 {
	return time.Now().UnixNano()
}
