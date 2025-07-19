package routers

import "time"

// formatDate 格式化时间
// 参数：
//   - t: 时间对象
//
// 返回值：
//   - string: 格式化后的时间字符串
func formatDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// timeStamp 获取当前时间戳
//
// 返回值：
//   - int64: 当前时间戳
func timeStamp() int64 {
	now := time.Now()
	secondTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location())
	return secondTime.Unix()
}
