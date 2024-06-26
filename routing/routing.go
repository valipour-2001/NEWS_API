package routing

import (
	customMiddlewares "NEWS_API/Utility/middleware"
	"NEWS_API/config"
	"NEWS_API/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRouting(e *echo.Echo) error {

	RouteUserController(e)
	RouteAccountController(e)
	RouteNewsController(e)

	return nil
}

func RouteUserController(e *echo.Echo) {
	userController := controller.NewUserController()

	e.POST("/UploadAvatar", userController.UploadAvatar)

	g := e.Group("users")

	g.GET("/getList", userController.GetUserList)

	g.POST("/CreateNewUser", userController.CreateNewUser, customMiddlewares.PermissionChecker("CreateUser"), middleware.JWTWithConfig(config.AppConfig.DefJwtConfig))

	g.PUT("/EditUser/:id", userController.EditUser, customMiddlewares.PermissionChecker("EditUser"), middleware.JWTWithConfig(config.AppConfig.DefJwtConfig))
	g.DELETE("/DeleteUser/:id", userController.DeleteUser, customMiddlewares.PermissionChecker("DeleteUser"), middleware.JWTWithConfig(config.AppConfig.DefJwtConfig))

	g.PUT("/EditUserRole/:id", userController.EditUserRole, middleware.JWTWithConfig(config.AppConfig.DefJwtConfig))
	g.PUT("/EditUserPassword/:id", userController.EditUserPassword, middleware.JWTWithConfig(config.AppConfig.DefJwtConfig))

}

func RouteAccountController(e *echo.Echo) {
	accountController := controller.NewAccountController()
	e.POST("/login", accountController.LoginUser)
}

func RouteNewsController(e *echo.Echo) {
	newsController := controller.NewNewsController()

	newsGroup := e.Group("news")

	newsGroup.GET("/getList", newsController.GetNewsList)
	newsGroup.POST("/Create", newsController.CreateNews)
	newsGroup.POST("/Edit/:id", newsController.EditNews)
}
