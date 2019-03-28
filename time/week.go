package time

import (
	"time"
)

// YearWeek 获取年份，第几周
// eg: YearWeek() => 2018 52
func YearWeek() (int, int) {
	return time.Now().ISOWeek()
}

// YearDay 今天是今年的第几天
// eg: YearDay() => 362
func YearDay() int {
	return time.Now().YearDay()
}

// FirstTimeOfWeek 指定年，周 的第一天日期
// 代码来自于: https://xferion.com/golang-reverse-isoweek-get-the-date-of-the-first-day-of-iso-week/
// eg: YearDay(2019, 1) => 2018-12-31 00:00:00 +0800 CST
func FirstTimeOfWeek(year, week int) time.Time {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, time.Local)
	isoYear, isoWeek := date.ISOWeek()
	for date.Weekday() != time.Monday { // iterate back to Monday
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoYear < year { // iterate forward to the first day of the first week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoWeek < week { // iterate forward to the first day of the given week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	return date
}

// YearWeekZero 指定年，周 第一天的零点时间戳
// eg: YearWeekZero(2019, 1) => 1546185600
func YearWeekZero(year, week int) int64 {
	return FirstTimeOfWeek(year, week).Unix()
}

// WeekZero 本周第一天零点时间戳
// eg: WeekZero() => 1545580800
func WeekZero() int64 {
	year, week := YearWeek()
	return FirstTimeOfWeek(year, week).Unix()
}

// WeekZeroTime 本周第一天零点时间
// eg: WeekZeroTime() => 2018-12-24 00:00:00 +0800 CST
func WeekZeroTime() time.Time {
	year, week := YearWeek()
	return FirstTimeOfWeek(year, week)
}

// NextWeekZero 下一周第一天零点时间戳
// eg: NextWeekZero() => 1546185600
func NextWeekZero() int64 {
	year, week := YearWeek()
	if week == 52 {
		year++
		week = 1
	}
	return FirstTimeOfWeek(year, week).Unix()
}

// NextWeekZeroTime 下一周第一天零点时间
// eg: NextWeekZeroTime() => 2018-12-31 00:00:00 +0800 CST
func NextWeekZeroTime() time.Time {
	year, week := YearWeek()
	if week == 52 {
		year++
		week = 1
	}
	return FirstTimeOfWeek(year, week)
}

// WeekPass 本周已过去多少秒
func WeekPass() int64 {
	return Now() - WeekZero()
}

// WeekRemain 本周还剩多少秒
func WeekRemain() int64 {
	return NextWeekZero() - Now()
}

// WeekIndex 今天是本周第几天，周一为第一天
func WeekIndex() int {
	t := time.Now()
	return WeekIndexByTime(&t)
}

// WeekIndexByTime 指定日期为一周的第几天，周一为第一天
// 返回 1|2|3|4|5|6|7
func WeekIndexByTime(t *time.Time) int {
	if t == nil {
		return 0
	}
	return int((t.Weekday()+6)%7) + 1
}
