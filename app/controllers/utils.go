package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/qor/qor-example/app/models"
)

func CurrentOrder(ctx *gin.Context) *models.Order {
	var (
		user  = CurrentUser(ctx)
		order models.Order
	)
	if user == nil {
		return nil
	}

	DB(ctx).Preload("OrderItems").Where(map[string]interface{}{"user_id": user.ID, "state": "draft"}).FirstOrInit(&order)
	DB(ctx).Save(&order)

	return &order
}
