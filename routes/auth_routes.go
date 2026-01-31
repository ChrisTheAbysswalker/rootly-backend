package routes

import (
	"github.com/ChrisTheAbysswalker/rootly-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine, ctrl *controllers.AuthController) {
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/register", ctrl.Register)
		authGroup.POST("/login", ctrl.Login)
	}
}