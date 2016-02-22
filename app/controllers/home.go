package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeIndex(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"home_index.tmpl",
		gin.H{},
	)
}
