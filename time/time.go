package time

import (
	"time"
)

const (
	format     = "2006-01-02 15:04:05"
	formatYear = "2006-01-02"
)

// Sleep 暂停
// sleepSecond: 暂停的秒数
func Sleep(sleepSecond int) {
	if sleepSecond <= 0 {
		return
	}
	time.Sleep(time.Duration(sleepSecond) * time.Second)
}

// Second 设置秒
func Second(second int) time.Duration {
	return time.Duration(second) * time.Second
}

// Minute 设置分钟
func Minute(minute int) time.Duration {
	return time.Duration(minute) * time.Minute
}

// Hour 设置小时
func Hour(hour int) time.Duration {
	return time.Duration(hour) * time.Hour
}

const (
	// Y 2018
	Y int = iota
	// YM  2018-12
	YM
	// YM2 2018/12
	YM2
	// YM3 201812
	YM3
	// YMD 2018-12-31
	YMD
	// YMD2 2018/12/31
	YMD2
	// YMD3 20181231
	YMD3
	// YMDHMS 2018-12-31 12:33:55
	YMDHMS
	// YMDHMS2 2018/12/31 12:33:55
	YMDHMS2
	// YMDHMS3 20181231123355
	YMDHMS3
	// YMDHMSM 2018-12-31 12:33:55.332
	YMDHMSM
)

// Date 获取日期，字符串类型
// param: level 类型
// eg: Date(YMDHMS) => 2018-12-31 12:33:55
func Date(level int) string {
	format := "2006-01-02 15:04:05"
	switch level {
	case Y:
		format = "2006"
	case YM:
		format = "2006-01"
	case YM2:
		format = "2006/01"
	case YMD:
		format = "2006-01-02"
	case YMD2:
		format = "2006/01/02"
	case YMD3:
		format = "20060102"
	case YMDHMS:
		format = "2006-01-02 15:04:05"
	case YMDHMS2:
		format = "2006/01/02 15:04:05"
	case YMDHMS3:
		format = "20060102150405"
	case YMDHMSM:
		format = "2006-01-02 15:04:05.000"
	default:
	}
	return time.Now().Format(format)
}

// Now 获取当前时间戳，秒
// eg: Now() => 1543626923
func Now() int64 {
	return time.Now().Unix()
}

// Str2Now 将字符串转为时间戳
// param: dateString: YYYY-MM-DD hh:mm:ss
// eg: Str2Now("2018-10-01 00:00:00") => 1538352000
func Str2Now(dateString string) int64 {
	t, _ := time.Parse(format, dateString)
	return t.Unix()
}

// Str2Time 将字符串转为时间
// param: dateString YYYY-MM-DD hh:mm:ss
// eg: Str2Time("2018-10-01 00:00:00") => 2018-10-01 00:00:00 +0000 UTC
func Str2Time(dateString string) time.Time {
	t, _ := time.Parse(format, dateString)
	return t
}

// Zero 获取相差num天的零点时间戳
// param: num 相差的天数，可以为负数, 0 表示今天
// eg: Zero(-5) =>
func Zero(num int) int64 {
	s := time.Now().Format(formatYear)
	t, _ := time.ParseInLocation(formatYear, s, time.Local)
	return t.AddDate(0, 0, num).Unix()
}

// ZeroTime 获取相差num天的零点时间
// param: num 相差的天数，可以为负数, 0 表示今天
// eg: ZeroTime(-5) =>
func ZeroTime(num int) time.Time {
	s := time.Now().Format(formatYear)
	t, _ := time.ParseInLocation(formatYear, s, time.Local)
	return t.AddDate(0, 0, num)
}

// Pass 今天过去了多少秒
// eg: Pass() => 32458
func Pass() int64 {
	return Now() - Zero(0)
}

// Remain 今天还剩多少秒
// eg: Remain() => 14488
func Remain() int64 {
	return Zero(1) - Now()
}
