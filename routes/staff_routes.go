package routes

import (
	"github.com/ChrisTheAbysswalker/rootly-backend/controllers"
	"github.com/gin-gonic/gin"
	"fmt"
)

func SetupStaffRoutes(r *gin.Engine, ctrl *controllers.StaffController) {
    fmt.Println("👥 [Rutas de Staff] Configurando...")
    api := r.Group("/api")
    {
        fmt.Println("   -> GET    /api/staff          | Obtiene la lista de animales encargados")
        api.GET("/staff", ctrl.GetStaff)
    }
}