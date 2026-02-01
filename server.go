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
	especieService := &services.EspecieService{DB: db}
	staffService := &services.StaffService{DB: db}

    authController := &controllers.AuthController{AuthService: authService}
	especieController := &controllers.EspecieController{EspecieService: especieService}
	staffController := &controllers.StaffController{StaffService: staffService}

    routes.SetupAuthRoutes(r, authController)
	routes.SetupEspecieRoutes(r, especieController)
	routes.SetupStaffRoutes(r, staffController)

    r.Run(":8000")
}