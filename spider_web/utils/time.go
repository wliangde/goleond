/**********************************************
** @Des:
** @Author: wangliangde
** @Date:   2018/4/8 15:54
***********************************************/

package utils

import "time"

func Unix2TimeStr(ts int64) string {
	t := time.Unix(ts, 0)
	return t.Format("2006-01-02 15:04:05")
}

func Str2UnixTime(strTime string) uint32 {
	local, _ := time.LoadLocation("Local")
	toTime, _ := time.ParseInLocation("2006-01-02 15:04:05", strTime, local)
	return uint32(toTime.Unix())
}
