package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ChrisTheAbysswalker/rootly-backend/services"
	"github.com/ChrisTheAbysswalker/rootly-backend/models"
)

type EspecieController struct {
	EspecieService *services.EspecieService
}

func (ctrl *EspecieController) GetInventario(c *gin.Context) {
	especies, err := ctrl.EspecieService.GetInventario()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener el inventario"})
		return
	}
	c.JSON(http.StatusOK, especies)
}

func (ctrl *EspecieController) GetEspecie(c *gin.Context) {
	id := c.Param("id")
	especie, err := ctrl.EspecieService.GetEspecie(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener la especie"})
		return
	}
	c.JSON(http.StatusOK, especie)
}

func (ctrl *EspecieController) GetFamilias(c *gin.Context) {
	familias, err := ctrl.EspecieService.GetFamilias()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener las familias"})
		return
	}
	c.JSON(http.StatusOK, familias)
}

func (ctrl *EspecieController) GetEstados(c *gin.Context) {
	estados, err := ctrl.EspecieService.GetEstados()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener los estados"})
		return
	}
	c.JSON(http.StatusOK, estados)
}

func (ctrl *EspecieController) GetStats(c *gin.Context) {
	stats, err := ctrl.EspecieService.GetEcosistemaStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al calcular estadísticas"})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func (ctrl *EspecieController) CreateEspecie(c *gin.Context) {
    var nuevaEspecie models.Especie
    if err := c.ShouldBindJSON(&nuevaEspecie); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }
    if err := ctrl.EspecieService.Create(&nuevaEspecie); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, nuevaEspecie)
}

func (ctrl *EspecieController) UpdateEspecie(c *gin.Context) {
    id := c.Param("id")
    var data models.Especie
    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }
    if err := ctrl.EspecieService.Update(id, &data); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Especie no encontrada"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Especie actualizada"})
}

func (ctrl *EspecieController) DeleteEspecie(c *gin.Context) {
    id := c.Param("id")
    if err := ctrl.EspecieService.Delete(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Especie eliminada"})
}

func (ctrl *EspecieController) GetRegistroSaludByID(c *gin.Context) {
    id := c.Param("id")
    salud, err := ctrl.EspecieService.GetRegistroSaludByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el registro"})
        return
    }
    c.JSON(http.StatusOK, salud)
}

func (ctrl *EspecieController) CreateRegistroSalud(c *gin.Context) {
    var data models.RegistroSalud
    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }
    if err := ctrl.EspecieService.CreateRegistroSalud(&data); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, data)
}

func (ctrl *EspecieController) UpdateRegistroSalud(c *gin.Context) {
    id := c.Param("id")
    var data models.RegistroSalud
    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }
    if err := ctrl.EspecieService.UpdateRegistroSalud(id, &data); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Registro salud no encontrada"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Registro salud actualizada"})
}