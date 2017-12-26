package admin

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/now"
	"github.com/qor/admin"
	"github.com/qor/qor-example/config/db"
)

type Chart struct {
	Total string
	Date  time.Time
}

/*
date format 2015-01-23
*/
func GetChartData(table, start, end string) (res []Chart) {
	startdate, err := now.Parse(start)
	if err != nil {
		return
	}

	enddate, err := now.Parse(end)
	if err != nil || enddate.UnixNano() < startdate.UnixNano() {
		enddate = now.EndOfDay()
	} else {
		enddate = enddate.AddDate(0, 0, 1)
	}

	db.DB.Table(table).
		Where("created_at > ? AND created_at < ?", startdate, enddate).
		Select("date(created_at) as date, count(*) as total").
		Group("date(created_at)").
		Order("date(created_at)").
		Scan(&res)
	return
}

type Charts struct {
	Orders []Chart
	Users  []Chart
}

func ReportsDataHandler(context *admin.Context) {
	charts := &Charts{}
	startDate := context.Request.URL.Query().Get("startDate")
	endDate := context.Request.URL.Query().Get("endDate")

	charts.Orders = GetChartData("orders", startDate, endDate)
	charts.Users = GetChartData("users", startDate, endDate)

	b, _ := json.Marshal(charts)
	context.Writer.Write(b)
	return
}

// SetupDashboard setup dashboard
func SetupDashboard(Admin *admin.Admin) {
	// Add Dashboard
	Admin.AddMenu(&admin.Menu{Name: "Dashboard", Link: "/admin", Priority: 1})

	Admin.GetRouter().Get("/reports", ReportsDataHandler)
	initFuncMap(Admin)
}
