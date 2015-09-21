package admin

import (
	"time"

	"github.com/qor/qor-example/app/models"
)

func initFuncMap() {
	Admin.RegisterFuncMap("latest_orders", latestOrders)
	// Admin.RegisterFuncMap("last_week_orders_chart", lastWeekOrderChart)
	// Admin.RegisterFuncMap("last_week_users_chart", lastWeekUserChart)
	//Admin.RegisterFuncMap("last_week_channel_data", lastWeekChannelData)

}

func latestOrders() (orders []models.Order) {
	Admin.Config.DB.Order("id desc").Limit(5).Find(&orders)
	return
}

func lastWeekOrderChart() (res []models.Chart) {
	res = models.GetChartData("orders", time.Now().AddDate(0, 0, -6).Format("2006-01-02"), time.Now().Format("2006-01-02"))
	return
}

func lastWeekUserChart() (res []models.Chart) {
	res = models.GetChartData("users", time.Now().AddDate(0, 0, -6).Format("2006-01-02"), time.Now().Format("2006-01-02"))
	return
}
