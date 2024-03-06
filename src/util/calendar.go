package util

import "time"

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
