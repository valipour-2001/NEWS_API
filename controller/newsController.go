package controller

import (
	"NEWS_API/Utility"
	"NEWS_API/ViewModel/common/httpResponse"
	newsViewModel "NEWS_API/ViewModel/news"
	"NEWS_API/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

type NewsController interface {
	GetNewsList(c echo.Context) error
	CreateNews(c echo.Context) error
	EditNews(c echo.Context) error
}

type newsController struct {
}

func NewNewsController() NewsController {
	return newsController{}
}

func (nc newsController) GetNewsList(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	fmt.Println(apiContext.GetUserId())

	newsService := service.NewNewsService()
	newsList, err := newsService.GetNewsList()
	if err != nil {
		println(err)
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(newsList))
}
func (nc newsController) CreateNews(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)

	newNews := new(newsViewModel.CreateNewsViewModel)

	if err := apiContext.Bind(newNews); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse("Data not found"))
	}

	if err := c.Validate(newNews); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse(err))
	}

	file, err := apiContext.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse("image not found"))
	}

	newsService := service.NewNewsService()
	newNewsId, err := newsService.CreateNewUser(*newNews, file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		NewUserId string
	}{
		NewUserId: newNewsId,
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(userResData))
}
func (nc newsController) EditNews(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetNewsId := apiContext.Param("id")

	editNews := new(newsViewModel.EditNewsViewModel)

	if err := apiContext.Bind(editNews); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse("Data not found"))
	}

	if err := c.Validate(editNews); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse(err))
	}

	file, err := apiContext.FormFile("file")

	editNews.Id = targetNewsId

	newsService := service.NewNewsService()

	if !newsService.IsNewsExist(targetNewsId) {
		return c.JSON(http.StatusBadRequest, errors.New("User Not Found"))
	}

	err = newsService.EditNews(*editNews, file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(nil))
}
