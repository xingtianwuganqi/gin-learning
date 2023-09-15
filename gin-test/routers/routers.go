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

	cityRouter := r.Group("/city/v1")
	{
		cityRouter.GET("/level", handler.FindCityHandler)
		cityRouter.GET("/info", handler.GetSubInfo)
		cityRouter.GET("/level/info", handler.GetSubCitys)
		cityRouter.GET("/city", handler.GetCityCode)
	}

	return r
}
