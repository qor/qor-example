package admin

import (
	"encoding/json"

	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor/admin"
)

type Charts struct {
	Orders   []models.Chart
	Users    []models.Chart
	Channels []models.Chart
}

func ReportsDataHandler(context *admin.Context) {
	charts := &Charts{}
	startDate := context.Request.URL.Query().Get("startDate")
	endDate := context.Request.URL.Query().Get("endDate")

	charts.Orders = models.GetChartData("order_count", startDate, endDate)
	charts.Users = models.GetChartData("user_count", startDate, endDate)
	charts.Channels = models.GetChartData("order_channels", startDate, endDate)

	b, _ := json.Marshal(charts)
	context.Writer.Write(b)
	return
}

func initRouter() {
	Admin.GetRouter().Get("/reports", ReportsDataHandler)
}
