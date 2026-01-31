package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

)

func RunServer(db *gorm.DB) {
	r := gin.Default()

	r.Run(":8000")
}