package routers

import (
	"gin-test/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRouter() *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/user/v1")
	{
		userRouter.POST("/register", handler.UserRegister)
	}
	return r
}
