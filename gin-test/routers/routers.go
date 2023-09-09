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
		userRouter.DELETE("/delete", handler.DeleteUser)
		userRouter.GET("/find", handler.FindUnDelete)
		userRouter.PUT("/update", handler.UpdateUser)
	}
	return r
}
