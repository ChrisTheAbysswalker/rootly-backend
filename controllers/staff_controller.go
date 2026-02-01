package controllers

import (
	"net/http"
	"github.com/ChrisTheAbysswalker/rootly-backend/services"
	"github.com/gin-gonic/gin"
)

type StaffController struct {
	StaffService *services.StaffService
}

func (ctrl *StaffController) GetStaff(c *gin.Context) {
	staff, err := ctrl.StaffService.GetStaff()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener la lista del staff"})
		return
	}
	c.JSON(http.StatusOK, staff)
}