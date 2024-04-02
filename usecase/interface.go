package usecase

import "github.com/pradiptarana/url-shortener/model"

type URLUsecase interface {
	CreateShortURL(url *model.URL) (string, error)
	GetOriginalURL(string) (string, error)
}
