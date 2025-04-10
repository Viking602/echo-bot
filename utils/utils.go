package utils

import (
	"time"
)

func ParseTimeToTimestamp(timeStr string) int64 {
	// 如果是空字符串或无效值，返回0
	if timeStr == "0000-00-00 00:00:00" || timeStr == "" {
		return 0
	}

	// 解析时间字符串为time.Time对象
	t, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return 0
	}

	// 返回Unix时间戳（秒）
	return t.Unix()
}
