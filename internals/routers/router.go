package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nor-ko-hi-jp/gorilla-tutorial/internals/middlewares"
	"github.com/nor-ko-hi-jp/gorilla-tutorial/internals/repositories"
	"github.com/nor-ko-hi-jp/gorilla-tutorial/internals/usecases"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.NewRepositoryProvider([]interface{}{repositories.UserRepository{}}).Provide)
	router.Use(middlewares.NewUseCaseProvider([]interface{}{&usecases.FetchUser{}}).Provide)

	v1 := router.Group("/v1")
	{
		v1.Use(middlewares.ApiKeyRequired)
		users := v1.Group("/users")
		users.GET("/:id", NewUserController().Fetch)
	}

	return router
}
