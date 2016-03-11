package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
)

func UserIndex(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})
}

func UserShow(ctx *gin.Context) {
	var currentUser models.User
	if !db.DB.Where("id = ?", ctx.Param("id")).First(&currentUser).RecordNotFound() {
		ctx.JSON(http.StatusOK, &currentUser)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})
	}
}
