package models

import (
	"time"

	"github.com/jinzhu/now"
	"github.com/qor/qor-example/db"
)

type Chart struct {
	Total string
	Date  time.Time
}

/*
date format 2015-01-23
*/
func GetChartData(table, start, end string) (res []Chart) {
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
