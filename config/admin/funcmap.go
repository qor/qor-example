package admin

import (
	"time"

	"github.com/jinzhu/now"
	"github.com/qor/qor-example/db"
)

func initFuncMap() {
	Admin.RegisterFuncMap("last_week_label", lastWeekLabelData)
	Admin.RegisterFuncMap("last_week_orders", lastWeekOrderData)
	Admin.RegisterFuncMap("last_week_users", lastWeekUserData)
}

type ChartData struct {
	Total string
	Date  time.Time
}

func lastWeekLabelData() (res []string) {
	var dateLabel string
	for i := 0; i < 7; i++ {
		dateLabel = now.BeginningOfDay().AddDate(0, 0, -i).Format("Jan 2")
		res = append([]string{dateLabel}, res...)
	}
	return
}

func lastWeekOrderData() (res []ChartData) {
	db.DB.Table("orders").Select("date(created_at) as date, count(1) as total").Group("date(created_at)").Scan(&res)
	return
}

func lastWeekUserData() (res []ChartData) {
	db.DB.Table("users").Select("date(created_at) as date, count(1) as total").Group("date(created_at)").Scan(&res)
	return
}
