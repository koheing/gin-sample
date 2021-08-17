package middlewares

import (
	"github.com/gin-gonic/gin"

	"github.com/nor-ko-hi-jp/gorilla-tutorial/internals/utils"
)

type RepositoryProvider struct {
	Instances []interface{}
}

func NewRepositoryProvider(Instances []interface{}) *RepositoryProvider {
	return &RepositoryProvider{Instances}
}

func (provider *RepositoryProvider) Provide(context *gin.Context) {
	instanceMap := map[string]interface{}{}
	for _, instance := range provider.Instances {
		instanceMap[utils.GetTypeName(instance)] = instance
	}

	if _, isExisted := context.Get("repositories"); !isExisted {
		context.Set("repositories", instanceMap)
	}

	context.Next()
}

type UseCaseProvider struct {
	Instances []interface{}
}

func NewUseCaseProvider(Instances []interface{}) *UseCaseProvider {
	return &UseCaseProvider{Instances}
}

func (provider *UseCaseProvider) Provide(context *gin.Context) {
	instanceMap := map[string]interface{}{}
	for _, instance := range provider.Instances {
		instanceMap[utils.GetTypeName(instance)] = instance
	}

	if _, isExisted := context.Get("usecases"); !isExisted {
		context.Set("usecases", instanceMap)
	}

	context.Next()
}
