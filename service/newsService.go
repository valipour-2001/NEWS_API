package service

import (
	newsViewModel "NEWS_API/ViewModel/news"
	"NEWS_API/model/news"
	"NEWS_API/repository"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type NewsService interface {
	GetNewsList() ([]news.News, error)
	CreateNewUser(userInput newsViewModel.CreateNewsViewModel, imageFile *multipart.FileHeader) (string, error)
	IsNewsExist(id string) bool
	EditNews(userInput newsViewModel.EditNewsViewModel, imageFile *multipart.FileHeader) error
}

type newsService struct {
}

func NewNewsService() NewsService {
	return newsService{}
}

func (newsService) GetNewsList() ([]news.News, error) {

	newsRepository := repository.NewNewsRepository()
	newsList, err := newsRepository.GetNewsList()
	return newsList, err
}
func (s newsService) CreateNewUser(userInput newsViewModel.CreateNewsViewModel, imageFile *multipart.FileHeader) (string, error) {

	newsEntity := news.News{
		Title:            userInput.Title,
		ImageName:        userInput.ImageName,
		ShortDescription: userInput.ShortDescription,
		Description:      userInput.Description,
		CreateDate:       time.Now(),
		CreatorUserId:    userInput.CreatorUserId,
	}
	if imageFile != nil {
		src, err := imageFile.Open()
		if err != nil {
			return "", err
		}

		fileName := uuid.New().String() + filepath.Ext(imageFile.Filename)

		wd, err := os.Getwd()
		imageServerPath := filepath.Join(wd, "wwwroot", "images", "news", fileName)

		des, err := os.Create(imageServerPath)
		if err != nil {
			return "", err
		}
		defer des.Close()

		_, err = io.Copy(des, src)
		if err != nil {
			return "", err
		}
		newsEntity.ImageName = fileName
	}

	newsRepository := repository.NewNewsRepository()
	newsId, err := newsRepository.InsertNews(newsEntity)

	return newsId, err
}

func (s newsService) IsNewsExist(id string) bool {
	newsRepository := repository.NewNewsRepository()
	_, err := newsRepository.GetNewsById(id)

	if err != nil {
		return false
	}

	return true
}

func (s newsService) EditNews(userInput newsViewModel.EditNewsViewModel, imageFile *multipart.FileHeader) error {

	newsRepository := repository.NewNewsRepository()
	newsEntity := news.News{
		Id:               userInput.Id,
		Title:            userInput.Title,
		ImageName:        userInput.ImageName,
		ShortDescription: userInput.ShortDescription,
		Description:      userInput.Description,
		CreateDate:       time.Now(),
		CreatorUserId:    userInput.CreatorUserId,
	}
	if imageFile != nil {
		src, err := imageFile.Open()
		if err != nil {
			return err
		}
		oldNews, err := newsRepository.GetNewsById(userInput.Id)
		if err != nil {
			return err
		}

		wd, err := os.Getwd()

		if oldNews.ImageName != "" {
			oldImageServerPath := filepath.Join(wd, "wwwroot", "images", "news", oldNews.ImageName)
			os.Remove(oldImageServerPath)
		}

		fileName := uuid.New().String() + filepath.Ext(imageFile.Filename)

		imageServerPath := filepath.Join(wd, "wwwroot", "images", "news", fileName)

		des, err := os.Create(imageServerPath)
		if err != nil {
			return err
		}
		defer des.Close()

		_, err = io.Copy(des, src)
		if err != nil {
			return err
		}
		newsEntity.ImageName = fileName
	}

	err := newsRepository.UpdateNewsById(newsEntity)

	return err
}
