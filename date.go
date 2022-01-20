package lib

import (
	"fmt"
	"strings"
	"time"
)

const (
	Year     = "06"
	LongYear = "2006"
	Month     = "Jan"
	ZeroMonth = "01"
	NumMonth  = "1"
	LongMonth = "January"
	Day         = "2"
	ZeroDay     = "02"
	UnderDay    = "_2"
	WeekDay     = "Mon"
	LongWeekDay = "Monday"
	Hour       = "15"
	ZeroHour12 = "03"
	Hour12     = "3"
	Minute     = "4"
	ZeroMinute = "04"
	Second     = "5"
	ZeroSecond = "05"
)

// 自定义
const (
	CHINESE_DATE_LONG_FORMAT = "2006-01-02 15:04:05"
	CHINESE_DATE_SHOT_FORMAT = "2006-01-01 15:04:05"
)

//根据传入的Y-m-d、 Y/m/d日期格式  生成Go语言time中对应的 显示格式字符串
func format(layout string) string {
	r := strings.NewReplacer("Y", LongYear, "y", Year, "m", ZeroMonth, "d", ZeroDay, "H", Hour, "h", ZeroHour12, "i", ZeroMinute, "s", ZeroSecond)
	return r.Replace(layout)
}

// 获取当前时间的日期格式字符串
func GetTime() string {
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(CHINESE_DATE_LONG_FORMAT)
	fmt.Println(t)
	return str
}

// 将日期时间格式的字符串转为 time.Time类型
func StrToTime(str string) (*time.Time, error) {
	t, err := time.ParseInLocation(CHINESE_DATE_LONG_FORMAT, str, time.Local)
	if err != nil {
		return &time.Time{}, err
	}
	return &t, nil
}

/**
 * 将time.Time类型按字符串Y-m-d H:i:s格式格式化
 * t 要格式化的时间
 * layout 要格式化的时间格式 Y-m-d H:i:s、 Y-m-d等等
 */
func Date(layout string, t time.Time) string {
	f := format(layout)
	o := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	return o.Format(f)
}
