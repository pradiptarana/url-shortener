package repository

import "github.com/pradiptarana/url-shortener/model"

//go:generate mockgen -destination=../mocks/mock_task.go -package=mocks github.com/pradiptarana/url-shortener/repository UserRepository
type URLRepository interface {
	CreateURL(*model.URL) error
	GetOriginalURL(string) (string, error)
	GetMaxID() (int64, error)
	GetOriginalURLCache(shortURL string) string
	SetURLCache(url *model.URL)
}

//go:generate mockgen -destination=../mocks/mock_user.go -package=mocks github.com/pradiptarana/url-shortener/repository URLRepository
type UserRepository interface {
	SignUp(us *model.User) error
	GetUser(username string) (*model.User, error)
}
