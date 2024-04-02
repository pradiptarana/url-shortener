package main

import (
	"time"

	"github.com/gin-gonic/gin"
	ginlogrus "github.com/toorop/gin-logrus"

	"github.com/patrickmn/go-cache"
	dbd "github.com/pradiptarana/url-shortener/internal/db"
	env "github.com/pradiptarana/url-shortener/internal/env"
	limiter "github.com/pradiptarana/url-shortener/internal/limiter"
	logger "github.com/pradiptarana/url-shortener/internal/logger"
	urlRepo "github.com/pradiptarana/url-shortener/repository/url"
	urlTr "github.com/pradiptarana/url-shortener/transport/api/url"
	urlUC "github.com/pradiptarana/url-shortener/usecase/url"
)

func main() {
	log := logger.Init()

	c := cache.New(5*time.Minute, 10*time.Minute)
	err := env.LoadEnv()
	if err != nil {
		log.Fatalf("error when load env file")
	}
	db := dbd.NewDBConnection()
	urRepo := urlRepo.NewURLRepository(db, c)
	urlUC := urlUC.NewURLUC(urRepo, log)
	urlTr := urlTr.NewURLTransport(urlUC)

	router := gin.Default()
	router.Use(ginlogrus.Logger(log), gin.Recovery())
	r := router.Group("/api/v1")
	r.POST("/url", limiter.Middleware(), urlTr.CreateShortURL)
	router.GET("/:short_url", limiter.Middleware(), urlTr.GetURL)

	router.Run("localhost:8080")
}
