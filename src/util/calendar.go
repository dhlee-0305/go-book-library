package util

import (
	"strings"
	"time"
)

func NowDay() string {
	now := time.Now()
	nowDay := now.Format("2006-01-02")

	return nowDay
}

func NowDayTime() string {
	now := time.Now()
	nowDay := now.Format("2006-01-02 15:04:05")

	return nowDay
}

func ConvertDayFormat(dayStr string) string {
	dayFormatStr := ""
	if len(dayStr) != 0 && len(dayStr) < 10 {
		dayFormatStr = "20" + strings.Replace(dayStr, ".", "-", -1)
	} else {
		dayFormatStr = strings.Replace(dayStr, ".", "-", -1)

	}

	return strings.Replace(dayFormatStr, "00", "01", -1)
}
