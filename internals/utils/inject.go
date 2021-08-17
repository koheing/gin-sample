package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Injector struct {
	TargetType string
	Context    *gin.Context
}

func (injector *Injector) Inject(instance interface{}) interface{} {
	value, isExisted := injector.Context.Get(injector.TargetType)

	if !isExisted {
		fmt.Println("instances are no existed")
		injector.Context.Abort()
		return nil
	}

	instances, isCasted := value.(map[string]interface{})
	if !isCasted {
		fmt.Println("instances cannot cast to map[string]interface{}")
		injector.Context.Abort()
		return nil
	}

	return instances[GetTypeName(instance)]
}
