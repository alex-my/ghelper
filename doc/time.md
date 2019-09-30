# 说明

时间，日期相关函数

# API

- **func Sleep(sleepSecond int)**

  - 功能: 暂停

- **func Second(second int) time.Duration**
  - 功能: 设置秒

* **func Minute(minute int) time.Duration**
  - 功能: 设置分钟
* **func Hour(hour int) time.Duration**
  - 功能: 设置小时

- **func Date(level int) string**

  - 功能: 获取日期，字符串类型
  - 参数:

    - **level**: 类型，见

      ```go
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
      ```

  - 示例:

    ```go
    Date(YMDHMS) // => 2018-12-31 12:33:55
    ```

- **func Now() int64**

  - 功能: 获取当前时间戳，秒
  - 示例:

    ```go
    Now() // => 1543626923
    ```

- **func Str2Now(dateString string) int64**

  - 功能: 将字符串转为时间戳
  - 示例:

    ```go
    Str2Now("2018-10-01 00:00:00") // => 1538352000
    ```

- **func Str2Time(dateString string) time.Time**

  - 功能: 将字符串转为时间
  - 示例:

    ```go
    Str2Time("2018-10-01 00:00:00") // => 2018-10-01 00:00:00 +0000 UTC
    ```

- **func Zero(num int) int64**

  - 功能: 获取相差 num 天的零点时间戳
  - 参数:

    - **num**: 相差的天数，可以为负数, 0 表示今天

- **func ZeroTime(num int) time.Time**

  - 功能: 获取相差 num 天的零点时间
  - 参数:
    - **num**: 相差的天数，可以为负数, 0 表示今天

- **func Pass() int64**

  - 功能: 今天过去了多少秒
  - 示例:

    ```go
    Pass() // => 32458
    ```

- **func Remain() int64**

  - 功能: 今天还剩多少秒
  - 示例:

    ```go
    Remain() // => 14488
    ```

- **func YearWeek() (int, int)**

  - 功能: 获取年份，一年中的第几周
  - 示例:

    ```go
    YearWeek() // => 2018 52
    ```

- **func YearDay() int**

  - 功能: 今天是今年的第几天
  - 示例:

    ```go
    YearDay() // => 362
    ```

* **func FirstTimeOfWeek(year, week int) time.Time**

  - 功能: 指定年，周 的第一天日期
  - 示例:

    ```go
    // 需要注意 2019 年第一周从 2018 年就开始了
    YearDay(2019, 1) // => 2018-12-31 00:00:00 +0800 CST
    ```

* **func YearWeekZero(year, week int) int64**

  - 功能:
  - 示例:

    ```go
    YearWeekZero(2019, 1) // => 1546185600
    ```

* **func WeekZero() int64**

  - 功能: 本周第一天零点时间戳
  - 示例:

    ```go
    WeekZero() // => 1545580800
    ```

* **func WeekZeroTime() time.Time**

  - 功能: 本周第一天零点时间
  - 示例:

    ```go
    WeekZeroTime() // => 2018-12-24 00:00:00 +0800 CST
    ```

* **func NextWeekZero() int64**

  - 功能: 下一周第一天零点时间戳
  - 示例:

    ```go
    NextWeekZero() // => 1546185600
    ```

* **func NextWeekZeroTime() time.Time**

  - 功能: 下一周第一天零点时间
  - 示例:

    ```go
    NextWeekZeroTime() // => 2018-12-31 00:00:00 +0800 CST
    ```

* **func WeekPass() int64**

  - 功能: 本周已过去多少秒

* **func WeekRemain() int64**

  - 功能: 本周还剩多少秒

* **func WeekIndex() int**

  - 功能: 今天是本周第几天，周一为第一天

* **func WeekIndexByTime(t \*time.Time) int**
  - 功能: 指定日期为一周的第几天，周一为第一天
  - 返回值: 返回 1|2|3|4|5|6|7
