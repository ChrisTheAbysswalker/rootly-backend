package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ChrisTheAbysswalker/rootly-backend/controllers"
    "fmt"
)

func SetupEspecieRoutes(r *gin.Engine, ctrl *controllers.EspecieController) {
    fmt.Println("🌿 [Rutas de Especies] Configurando...")
    api := r.Group("/api")
    {
        fmt.Println("   -> GET    /api/especies       | Lista el inventario botánico completo")
        fmt.Println("   -> POST   /api/especies       | Registra una nueva planta en el sistema")
        fmt.Println("   -> PUT    /api/especies/:id   | Actualiza datos de una planta existente")
        fmt.Println("   -> DELETE /api/especies/:id   | Elimina una especie del inventario")
        fmt.Println("   -> GET    /api/ecosistema/stats | Obtiene métricas (humedad, salud, alertas)")
        
        api.GET("/especies", ctrl.GetInventario)     
        api.GET("/especies/:id", ctrl.GetEspecie)
        api.POST("/especies", ctrl.CreateEspecie)     
        api.PUT("/especies/:id", ctrl.UpdateEspecie)   
        api.DELETE("/especies/:id", ctrl.DeleteEspecie)

        api.GET("/ecosistema/stats", ctrl.GetStats)

        api.GET("/familias", ctrl.GetFamilias)
        api.GET("/estados", ctrl.GetEstados)

        api.GET("/registro_salud/:id", ctrl.GetRegistroSaludByID)
        api.POST("/registro_salud", ctrl.CreateRegistroSalud)
        api.PUT("/registro_salud/:id", ctrl.UpdateRegistroSalud)
    }
}