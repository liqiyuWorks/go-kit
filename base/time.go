/*
 * @Author: lisheng
 * @Date: 2022-10-29 11:18:30
 * @LastEditTime: 2022-12-17 16:59:09
 * @LastEditors: lisheng
 * @Description: 时间模块
 * @FilePath: /jf-go-kit/base/time.go
 */
package base

import (
	"fmt"
	"strconv"
	"time"
)

func GetDate() string {
	now := time.Now()
	year, month, day := now.Date()
	return fmt.Sprintf("%d%02d%02d", year, month, day)
}

func GetDateByLocation(location *time.Location) string {
	now := time.Now().In(location)
	year, month, day := now.Date()
	return fmt.Sprintf("%d%02d%02d", year, month, day)
}

func GetTimeByTz(tz string) (*time.Time, error) {
	localtion, err := time.LoadLocation(tz)
	if err != nil {
		return nil, err
	}
	now := time.Now().In(localtion)
	return &now, nil
}

func GetDateNum(location *time.Location) int {
	var dateName string
	if location == nil {
		dateName = GetDate()
	} else {
		dateName = GetDateByLocation(location)
	}
	dateNum, _ := strconv.ParseInt(dateName, 10, 32)
	return int(dateNum)
}

func GetDateNum2(now *time.Time) int {
	year, month, day := now.Date()
	dateName := fmt.Sprintf("%d%02d%02d", year, month, day)
	dateNum, _ := strconv.ParseInt(dateName, 10, 32)
	return int(dateNum)
}

func GetDateHour() string {
	now := time.Now()
	year, month, day := now.Date()
	return fmt.Sprintf("%d%02d%02d%02d", year, month, day, now.Hour())
}

func GetTimeStampMs() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}

func GetTimeStampSec() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TimeName(t time.Time) string {
	return t.Format("2006-01-02 15:04:05 -0700")
}

func PrintPerformTimeConsuming(funcName string, startTime time.Time) {
	timeConumeSeconds := time.Now().Sub(startTime).Seconds()
	Glog.Debugf("function[%s] using [%f] seconds", funcName, timeConumeSeconds)
}

func GetDayEndTimeLocal() time.Time {
	year, month, day := time.Now().Date()
	endTime := time.Date(year, month, day, 23, 59, 59, 0, time.Local)
	return endTime
}

func GetDayEndTimeUtc() time.Time {
	year, month, day := time.Now().UTC().Date()
	endTime := time.Date(year, month, day, 23, 59, 59, 0, time.UTC)
	return endTime
}

func GetNextDayStartTimeLocal() time.Time {
	year, month, day := time.Now().Date()
	startTime := time.Date(year, month, day+1, 0, 0, 0, 0, time.Local)
	return startTime
}

func GetNextDayStartTimeUtc() time.Time {
	year, month, day := time.Now().UTC().Date()
	startTime := time.Date(year, month, day+1, 0, 0, 0, 0, time.UTC)
	return startTime
}

func GetNextDayStartTimeByLocation(location *time.Location) time.Time {
	year, month, day := time.Now().In(location).Date()
	startTime := time.Date(year, month, day+1, 0, 0, 0, 0, location)
	return startTime
}

func GetWeekName(location *time.Location) int {
	if location == nil {
		location = time.Local
	}
	year, week := time.Now().In(location).ISOWeek()
	weekName, _ := strconv.ParseInt(fmt.Sprintf("%d%02d", year, week), 10, 32)
	return int(weekName)
}

func GetWeekDay(location *time.Location) int32 {
	if location == nil {
		location = time.Local
	}
	return int32(time.Now().In(location).Weekday())
}

func GetMonthName(location *time.Location) int {
	if location == nil {
		location = time.Local
	}
	year, month, _ := time.Now().In(location).Date()
	monthName, _ := strconv.ParseInt(fmt.Sprintf("%d%02d", year, month), 10, 32)
	return int(monthName)
}

// 计算某年某月的天数
func GetMonthlyDayCount(year int, month int) int {
	var days int
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30
		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	return days
}

/**
 * @description: 计算最接近某小时的时间戳
 * @param {int} ts
 * @return {*}
 * @author: liqiyuWorks
 */
func GetLastHourToTs(ts int64) (lastHour int64) {
	if ts%(60*60*1000) == 0 {
		lastHour = ts
	} else {
		lastHour = (ts / (60 * 60 * 1000)) * 60 * 60 * 1000
	}
	return lastHour
}

/**
 * @description: 计算最接近n个小时的时间戳, 前后匹配，找到最接近的n小时的时刻
 * @param {int64} ts
 * @param {int64} hour
 * @return {*}
 * @author: liqiyuWorks
 */
func GetSpecifiedNumToTs(ts int64, hour int64) (SpecifiedHour int64) {
	diffN := ts % (hour * 60 * 60 * 1000)
	if ts%(hour*60*60*1000) == 0 {
		SpecifiedHour = ts
	} else {
		if diffN > (hour * 3600 * 1000 / 2) {
			SpecifiedHour = (ts/(hour*3600*1000))*hour*3600*1000 + hour*3600*1000
			// fmt.Println("=>贴近后者, ts: ", SpecifiedHour)
		} else {
			SpecifiedHour = (ts / (hour * 3600 * 1000)) * hour * 3600 * 1000
			// fmt.Println("=>贴近前者, ts: ", SpecifiedHour)
		}
	}
	return SpecifiedHour
}
