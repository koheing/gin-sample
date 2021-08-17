package usecases

import (
	"github.com/gin-gonic/gin"
	"github.com/nor-ko-hi-jp/gorilla-tutorial/internals/interfaces"
	"github.com/nor-ko-hi-jp/gorilla-tutorial/internals/repositories"
	"github.com/nor-ko-hi-jp/gorilla-tutorial/internals/utils"
)

type FetchUser struct {
	UserRepository interfaces.ReadRepository
}

func (fetchUser *FetchUser) Initialize(context *gin.Context) {
	injector := utils.Injector{Context: context, TargetType: "repositories"}
	userRepository, isCasted := injector.Inject(repositories.UserRepository{}).(interfaces.ReadRepository)
	if !isCasted {
		context.JSON(200, gin.H{"messsage": "NONE"})
		context.Abort()
		return
	}

	fetchUser.UserRepository = userRepository
}

func (fetchUser *FetchUser) Execute(parameter interfaces.UseCaseParameter) *interfaces.Result {
	id := parameter["id"].(string)
	result := fetchUser.UserRepository.Find(id)
	return result
}
