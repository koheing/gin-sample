package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nor-ko-hi-jp/gorilla-tutorial/internals/interfaces"
	"github.com/nor-ko-hi-jp/gorilla-tutorial/internals/usecases"
	"github.com/nor-ko-hi-jp/gorilla-tutorial/internals/utils"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (user *UserController) Fetch(context *gin.Context) {
	injector := utils.Injector{Context: context, TargetType: "usecases"}
	fetchUser, isCasted := injector.Inject(&usecases.FetchUser{}).(interfaces.BaseUseCase)
	if !isCasted {
		context.JSON(200, gin.H{"messsage": "NONE"})
		context.Abort()
		return
	}

	fetchUser.Initialize(context)

	result := fetchUser.Execute(map[string]interface{}{
		"id": "user_0",
	})

	context.JSON(200, result.Mapper.ToDomain(result.Value))
}
