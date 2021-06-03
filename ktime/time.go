package ktime

import (
	"errors"
	"time"
)

//一些常量 解析时间时使用
const (
	defData     = "2006-01-02"
	defTime     = "15:04:05"
	defDateTime = "Mon, 02 Jan 2006-01-02 15:04:05"
)

var cstZone *time.Location

//初始化
func init() {
	//名称：CST中国时间 偏移量 八个小时
	cstZone = time.FixedZone("CST", 8*3600)
}

//解析时间为日期类型的字符串 格式：类 2006-01-02
func TimeParseData(t time.Time) string {
	return parseTimeToStr(t.In(cstZone),defData)
}

//解析时间为时间类型的字符串 格式：类 15:04:05
func TimeParseTime(t time.Time) string {
	return parseTimeToStr(t.In(cstZone),defTime)
}

//解析时间为时间和日期类型的字符串 格式：类 Mon, 02 Jan 2006-01-02 15:04:05
func TimeParseDataAndTime(t time.Time) string {
	return parseTimeToStr(t.In(cstZone),defDateTime)
}

//解析字符串为日期类型
func StrParseData(str string)(time.Time,error)  {
	return parseStrToTime(str,defData)
}

//解析字符串为时间类型
func StrParseTime(str string)(time.Time,error)  {
	return parseStrToTime(str,defTime)
}

//解析字符串为日期和时间类型
func StrParseDataAndTime(str string)(time.Time,error)  {
	return parseStrToTime(str,defDateTime)
}

//解析时间到字符串 不对外暴露
func parseTimeToStr(t time.Time, layout string) string {
	return t.Format(layout)
}

//解析字符串到时间 不对外暴露
func parseStrToTime(str, layout string) (time.Time, error) {
	if str == "" {
		return time.Now(), errors.New("the source string cannot be empty")
	}
	return time.ParseInLocation(layout, str, cstZone)
}

//获取当前时间的字符串
func GetNowTimeStr() string {
	return parseTimeToStr(time.Now().In(cstZone),defTime)
}

//获取当前日期的字符串
func GetNowDataStr() string {
	return parseTimeToStr(time.Now().In(cstZone),defData)
}

//获取当前时间和日期的字符串
func GetNowTimeAndDataStr() string {
	return parseTimeToStr(time.Now().In(cstZone),defDateTime)
}


//比较两个时间是否相等
func EqualTimes(t1, t2 time.Time) bool {
	return t1.Equal(t2)
}

