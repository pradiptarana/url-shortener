package url_test

import (
	"errors"
	"testing"

	"github.com/pradiptarana/url-shortener/mocks"
	"github.com/pradiptarana/url-shortener/model"

	"github.com/golang/mock/gomock"

	logger "github.com/pradiptarana/url-shortener/internal/logger"
	urlUC "github.com/pradiptarana/url-shortener/usecase/url"
)

var log = logger.Init()

func TestCreateFailedValidation(t *testing.T) {
	req := &model.URL{
		OriginalURL: "",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockURLRepo := mocks.NewMockURLRepository(mockCtrl)
	urlUC := urlUC.NewURLUC(mockURLRepo, log)

	_, err := urlUC.CreateShortURL(req)
	if err == nil {
		t.Fail()
	}

	req.OriginalURL = "www"
	_, err = urlUC.CreateShortURL(req)
	if err == nil {
		t.Fail()
	}
}

func TestCreateFailedGetMaxId(t *testing.T) {
	req := &model.URL{
		OriginalURL: "https://www.google.com/abc",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockURLRepo := mocks.NewMockURLRepository(mockCtrl)
	urlUC := urlUC.NewURLUC(mockURLRepo, log)

	mockURLRepo.EXPECT().GetMaxID().Return(int64(0), errors.New("error get Max ID")).Times(1)

	_, err := urlUC.CreateShortURL(req)
	if err == nil {
		t.Fail()
	}
}

func TestCreateFailedCreateToDB(t *testing.T) {
	req := &model.URL{
		OriginalURL: "https://www.google.com/abc",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockURLRepo := mocks.NewMockURLRepository(mockCtrl)
	urlUC := urlUC.NewURLUC(mockURLRepo, log)

	mockURLRepo.EXPECT().GetMaxID().Return(int64(0), nil).Times(1)
	repoReq := req
	repoReq.ShortURL = "27qMi57J"
	mockURLRepo.EXPECT().CreateURL(repoReq).Return(errors.New("error db")).Times(1)

	_, err := urlUC.CreateShortURL(req)
	if err == nil {
		t.Fail()
	}
}

func TestCreateSuccess(t *testing.T) {
	req := &model.URL{
		OriginalURL: "https://www.google.com/abc",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockURLRepo := mocks.NewMockURLRepository(mockCtrl)
	urlUC := urlUC.NewURLUC(mockURLRepo, log)

	mockURLRepo.EXPECT().GetMaxID().Return(int64(0), nil).Times(1)
	repoReq := req
	repoReq.ShortURL = "2XNGAK"
	mockURLRepo.EXPECT().CreateURL(repoReq).Return(nil).Times(1)

	shortURL, err := urlUC.CreateShortURL(req)
	if err != nil {
		t.Fail()
	}

	if shortURL != repoReq.ShortURL {
		t.Fail()
	}
}

func TestGetOriginalUrlFailed(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockURLRepo := mocks.NewMockURLRepository(mockCtrl)
	urlUC := urlUC.NewURLUC(mockURLRepo, log)

	mockURLRepo.EXPECT().GetOriginalURLCache("2XNGAK").Return("").Times(1)
	mockURLRepo.EXPECT().GetOriginalURL("2XNGAK").Return("", errors.New("error db")).Times(1)

	_, err := urlUC.GetOriginalURL("2XNGAK")
	if err == nil {
		t.Fail()
	}
}

func TestGetOriginalUrlSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockURLRepo := mocks.NewMockURLRepository(mockCtrl)
	urlUC := urlUC.NewURLUC(mockURLRepo, log)

	mockURLRepo.EXPECT().GetOriginalURLCache("2XNGAK").Return("").Times(1)
	mockURLRepo.EXPECT().GetOriginalURL("2XNGAK").Return("https://www.google.com", nil).Times(1)
	mockURLRepo.EXPECT().SetURLCache(&model.URL{
		OriginalURL: "https://www.google.com",
		ShortURL:    "2XNGAK",
	})

	oriURL, err := urlUC.GetOriginalURL("2XNGAK")
	if err != nil {
		t.Fail()
	}

	if oriURL != "https://www.google.com" {
		t.Fail()
	}
}
