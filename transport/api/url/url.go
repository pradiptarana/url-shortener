package url

import (
	"net/http"

	"github.com/pradiptarana/url-shortener/model"
	"github.com/pradiptarana/url-shortener/usecase"

	"github.com/gin-gonic/gin"
)

type URLTransport struct {
	usecase.URLUsecase
}

func NewURLTransport(uc usecase.URLUsecase) *URLTransport {
	return &URLTransport{uc}
}

func (ut *URLTransport) CreateShortURL(c *gin.Context) {
	var req model.ShortingURLRequest
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	shortURL, err := ut.URLUsecase.CreateShortURL(&model.URL{
		OriginalURL: req.OriginalURL,
	})
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "create short url success", "data": shortURL})
	return
}

func (ut *URLTransport) GetURL(c *gin.Context) {
	shortURL := c.Param("short_url")
	if shortURL == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "short url is empty"})
		return
	}
	originalURL, err := ut.URLUsecase.GetOriginalURL(shortURL)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	c.Redirect(http.StatusMovedPermanently, originalURL)
	return
}
