package url

import (
	"errors"
	"math/big"
	"time"

	"net/url"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/pradiptarana/url-shortener/model"
	"github.com/pradiptarana/url-shortener/repository"
	"github.com/sirupsen/logrus"
)

type URLUC struct {
	repository.URLRepository
	*logrus.Logger
}

func NewURLUC(repo repository.URLRepository, log *logrus.Logger) *URLUC {
	return &URLUC{repo, log}
}

func (uc *URLUC) CreateShortURL(req *model.URL) (string, error) {
	if req.OriginalURL == "" {
		err := errors.New("not a valid url")
		uc.Logger.Error(err)
		return "", err
	}
	_, err := url.ParseRequestURI(req.OriginalURL)
	if err != nil {
		uc.Logger.Error(err)
		return "", err
	}

	maxID, err := uc.URLRepository.GetMaxID()
	if err != nil {
		uc.Logger.Error(err)
		return "", err
	}
	shortURL, id := uc.generateShortURL(maxID)
	req.CreatedAt = time.Now().Unix()
	req.ShortURL = shortURL
	req.Id = id
	err = uc.URLRepository.CreateURL(req)
	if err != nil {
		uc.Logger.Error(err)
		return "", err
	}
	return shortURL, nil
}

func (uc *URLUC) GetOriginalURL(shortURL string) (string, error) {
	originalURLCache := uc.URLRepository.GetOriginalURLCache(shortURL)
	if originalURLCache != "" {
		return originalURLCache, nil
	}
	originalURL, err := uc.URLRepository.GetOriginalURL(shortURL)
	if err != nil {
		uc.Logger.Error(err)
		return "", err
	}
	uc.URLRepository.SetURLCache(&model.URL{
		ShortURL:    shortURL,
		OriginalURL: originalURL,
	})
	return originalURL, nil
}

func (uc *URLUC) generateShortURL(maxID int64) (string, int) {
	if maxID == 0 {
		// to generate short url with length 6
		maxID = 1000000000
	} else {
		maxID = maxID + 1
	}
	return base58.Encode(big.NewInt(maxID).Bytes()), int(maxID)
}
