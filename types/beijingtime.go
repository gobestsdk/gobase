package types

import (
	"time"
)

//北京时间

func BeijingToday() (beijingday time.Time) {
	timeStr := time.Now().Add(-time.Hour * 8).Format("2006-01-02")
	//转成字符串时间，会丢失
	beijingday, _ = time.Parse("2006-01-02", timeStr)
	return
}
