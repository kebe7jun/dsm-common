package utils

import (
	"strconv"
	"time"
)

func TimeFormat(timestamp string) string {
	//时间戳字符串转换为日期格式获取指定格式日期
	date, _ := strconv.ParseInt(timestamp, 10, 64)
	return time.Unix(0, 0).Add(time.Duration(date) * time.Microsecond).Format("2006-01-02")
}

// NowTimeString will return now time as second, the type is string.
func NowTimeString() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

/**
 * startime: 1590464690690417
 * endtime: 1591464690691418
 * return: [2020-05-26,2020-05-27,2020-05-28,2020-05-29,2020-05-30,2020-05-31,2020-06-01,2020-06-02,2020-06-03,2020-06-04,2020-06-05,2020-06-06,2020-06-07]
 */
func GetDurationDateList(starTime string, endTime string) []string {
	var dateList []string

	starDate, _ := strconv.ParseInt(starTime, 10, 64)
	endDate, _ := strconv.ParseInt(endTime, 10, 64)

	startFormat := time.Unix(0, 0).Add(time.Duration(starDate) * time.Microsecond)
	endFormat := time.Unix(0, 0).Add(time.Duration(endDate)*time.Microsecond).AddDate(0, 0, 1)

	for startFormat.Before(endFormat) {
		dateList = append(dateList, startFormat.Format("2006-01-02"))
		startFormat = startFormat.AddDate(0, 0, 1)
	}

	return dateList
}
