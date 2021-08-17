package interfaces

import "github.com/gin-gonic/gin"

type UseCaseParameter map[string]interface{}

type BaseUseCase interface {
	Initialize(context *gin.Context)
	Execute(param UseCaseParameter) *Result
}
