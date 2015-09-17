package admin

import (
	"strconv"

	"github.com/jinzhu/now"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
)

func initFuncMap() {
	Admin.RegisterFuncMap("last_week_label", lastWeekLabelData)
	Admin.RegisterFuncMap("last_week_order_data", lastWeekOrderData)
	Admin.RegisterFuncMap("last_week_user_data", lastWeekUserData)

}

func lastWeekLabelData() (res []string) {
	var dateLabel string
	for i := 0; i < 7; i++ {
		dateLabel = now.BeginningOfDay().AddDate(0, 0, -i).Format("Jan 2")
		res = append([]string{dateLabel}, res...)
	}
	return
}

func lastWeekOrderData() (res []string) {
	var count int
	for i := 0; i < 7; i++ {
		db.DB.Model(&models.Order{}).Where("created_at > ? AND created_at < ?", now.BeginningOfDay().AddDate(0, 0, -i), now.BeginningOfDay().AddDate(0, 0, -(i-1))).Count(&count)
		res = append([]string{strconv.Itoa(count)}, res...)
	}
	return
}

func lastWeekUserData() (res []string) {
	var count int
	for i := 0; i < 7; i++ {
		db.DB.Model(&models.User{}).Where("created_at > ? AND created_at < ?", now.BeginningOfDay().AddDate(0, 0, -i), now.BeginningOfDay().AddDate(0, 0, -(i-1))).Count(&count)
		res = append([]string{strconv.Itoa(count)}, res...)
	}
	return
}
