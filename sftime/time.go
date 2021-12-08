package sftime

import (
	"github.com/omygoden/gotools/sfconst"
	"time"
)

//时间格式转时间戳
func DateToTime(date string) int64 {
	t, _ := time.ParseInLocation(sfconst.GO_TIME_FULL, date, time.Local)
	return t.Unix()
}

//时间戳转时间格式
func TimeToFullDate(t int64) string {
	return time.Unix(t, 0).Format(sfconst.GO_TIME_FULL)
}

//时间戳转自定义时间格式
func TimeToDate(t int64, format string) string {
	return time.Unix(t, 0).Format(format)
}

//格林威治日期转时间戳
func GMTDateToTime(date string) int64 {
	t, _ := time.ParseInLocation(sfconst.GO_TIME_GMT_FULL, date, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	return t.Unix()
}

//格林威治日期转正常时间
func GMTDateToNormalDate(date string) string {
	t, _ := time.ParseInLocation(sfconst.GO_TIME_GMT_FULL, date, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	return t.Format(sfconst.GO_TIME_FULL)
}
