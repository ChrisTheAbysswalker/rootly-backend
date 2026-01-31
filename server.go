package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "github.com/ChrisTheAbysswalker/rootly-backend/services"
    "github.com/ChrisTheAbysswalker/rootly-backend/controllers"
    "github.com/ChrisTheAbysswalker/rootly-backend/routes"
)

func RunServer(db *gorm.DB) {
    r := gin.Default()

    authService := &services.AuthService{DB: db}
    authController := &controllers.AuthController{AuthService: authService}

    routes.SetupAuthRoutes(r, authController)

    r.Run(":8000")
}