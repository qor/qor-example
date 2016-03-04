package admin

import "github.com/qor/qor-example/app/models"

func initFuncMap() {
	Admin.RegisterFuncMap("latest_orders", latestOrders)
}

func latestOrders() (orders []models.Order) {
	Admin.Config.DB.Order("id desc").Limit(5).Find(&orders)
	return
}
