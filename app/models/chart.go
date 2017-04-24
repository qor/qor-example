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
