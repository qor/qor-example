package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
)

// curl -i -X POST -H "Content-Type: application/json" -d "{ \"username\": \"pos0001\", \"password\": \"123456\", \"metod\": \"password\" }" http://localhost:7000/api/v1/auth/0
func LoginApi(ctx *gin.Context) {
	var auth models.Auth
	var currentUser models.User
	if ctx.BindJSON(&auth) == nil {
		if auth.Metod == "password" {
			if !db.DB.Where("password = ?", auth.Password).First(&currentUser).RecordNotFound() {
				// t := time.NowLoginAt: time.Now,
				t := time.Now()
				login := models.LogLogin{
					ClietIp:   ctx.ClientIP(),
					UserID:    currentUser.ID,
					User:      currentUser,
					InOut:     "in",
					LoginedAt: &t,
					Device:    "terminal",
				}
				// fmt.Println(login)
				if err := db.DB.Create(&login).Error; err != nil {
					fmt.Println(err)
				}
				user := models.UserApi{
					ID:     currentUser.ID,
					Name:   currentUser.Name,
					Email:  currentUser.Email,
					Gender: currentUser.Gender,
					Role:   currentUser.Role,
					Token:  "",
				}
				ctx.JSON(http.StatusOK, &user)
			} else {
				ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
			}
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request metod"})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})
	}
}

func LogoutApi(ctx *gin.Context) {
	var auth models.Auth
	var currentUser models.User
	if ctx.BindJSON(&auth) == nil {

		if !db.DB.Where("password = ?", auth.Password).First(&currentUser).RecordNotFound() {
			// t := time.NowLoginAt: time.Now,
			t := time.Now()
			login := models.LogLogin{
				ClietIp:   ctx.ClientIP(),
				UserID:    currentUser.ID,
				User:      currentUser,
				InOut:     "out",
				LoginedAt: &t,
			}
			// fmt.Println(login)
			if err := db.DB.Create(&login).Error; err != nil {
				fmt.Println(err)
			}
			user := models.UserApi{
				ID:     currentUser.ID,
				Name:   currentUser.Name,
				Email:  currentUser.Email,
				Gender: currentUser.Gender,
				Role:   currentUser.Role,
			}
			ctx.JSON(http.StatusOK, &user)
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})
	}
}
