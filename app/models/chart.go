package models

import (
	"time"

	"github.com/jinzhu/now"
	"github.com/qor/qor-example/db"
)

type Chart struct {
	Total string
	// Channel string
	Date time.Time
	Name string
}

/*
date format 2015-01-23
*/
func GetChartData(name, start, end string) (res []Chart) {
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
	switch name {
	case "order_count":
		db.DB.Table("orders").Where("created_at > ? AND created_at < ?", startdate, enddate).Select("date(created_at) as date, count(*) as total").Group("date(created_at)").Order("date(created_at)").Scan(&res)
	case "user_count":
		db.DB.Table("users").Where("created_at > ? AND created_at < ?", startdate, enddate).Select("date(created_at) as date, count(*) as total").Group("date(created_at)").Order("date(created_at)").Scan(&res)
	case "order_channels":
		db.DB.Table("orders").Where("created_at > ? AND created_at < ?", startdate, enddate).Select("channel as name, count(*) as total").Group("channel").Order("channel").Scan(&res)
	default:
		return
	}
	return
}
