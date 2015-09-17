package admin

import (
	"time"

	"github.com/jinzhu/now"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
)

func initFuncMap() {
	Admin.RegisterFuncMap("latest_orders", latestOrders)
	Admin.RegisterFuncMap("last_week_orders", lastWeekOrderData)
	Admin.RegisterFuncMap("last_week_users", lastWeekUserData)
}

func latestOrders() (orders []models.Order) {
	Admin.Config.DB.Order("id desc").Limit(5).Find(&orders)
	return
}

type ChartData struct {
	Total string
	Date  time.Time
}

/*
date format 2015-01-23
*/
func GetChartData(table, start, end string) (res []ChartData) {
	startdate, err := time.Parse("2006-01-02", start)
	if err != nil {
		return
	}
	enddate, err := time.Parse("2006-01-02", end)
	if err != nil || enddate.UnixNano() < startdate.UnixNano() {
		enddate = now.EndOfDay()
	} else {
		enddate.AddDate(0, 0, 1)
	}
	db.DB.Table(table).Where("created_at > ? AND created_at < ?", startdate, enddate).Select("date(created_at) as date, count(1) as total").Group("date(created_at)").Scan(&res)
	return
}

func lastWeekOrderData() (res []ChartData) {
	res = GetChartData("orders", time.Now().AddDate(0, 0, -6).Format("2006-01-02"), time.Now().Format("2006-01-02"))
	return
}

func lastWeekUserData() (res []ChartData) {
	res = GetChartData("users", time.Now().AddDate(0, 0, -6).Format("2006-01-02"), time.Now().Format("2006-01-02"))
	return
}
