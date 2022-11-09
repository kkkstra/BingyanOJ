package router

import (
	"byitter/src/controller"
	"byitter/src/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func initUserRouters(r *gin.Engine) {
	r.POST("/signup", controller.SignUp)
	r.POST("/signin", controller.SignIn)
	user := r.Group("/user")
	{
		user.GET("/:username", controller.GetUserProfile)
		userAction := user.Group("/:username")
		{
			userAction.Use(jwt.JwtAuthMiddleware())
			userActionEdit := userAction.Group("/edit")
			{
				userActionEdit.POST("/profile", controller.EditUserProfile)
				userActionEdit.POST("/password", controller.EditUserProfile)
				userActionEdit.POST("/email", controller.EditUserProfile)
			}
			userAction.POST("/del", controller.DeleteUser)
		}
	}
}
