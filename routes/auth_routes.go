package routes

import (
	"github.com/ChrisTheAbysswalker/rootly-backend/controllers"
	"github.com/gin-gonic/gin"
	"fmt"
)

func SetupAuthRoutes(r *gin.Engine, ctrl *controllers.AuthController) {
    fmt.Println("🔐 [Rutas de Seguridad] Configurando...")
    authGroup := r.Group("/api/auth")
    {
        fmt.Println("   -> POST   /api/auth/register  | Crea una cuenta nueva de usuario")
        fmt.Println("   -> POST   /api/auth/login     | Valida credenciales y entrega Token JWT")
        
        authGroup.POST("/register", ctrl.Register)
        authGroup.POST("/login", ctrl.Login)
    }
}