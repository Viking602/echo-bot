package utils

import (
	"fmt"
	"time"
)

func ParseTimeToTimestamp(timeStr string) int64 {
	// 如果是空字符串或无效值，返回0
	if timeStr == "0000-00-00 00:00:00" || timeStr == "" {
		return 0
	}

	// 解析时间字符串为time.Time对象
	t, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		return 0
	}

	// 返回Unix时间戳（秒）
	return t.Unix()
}

func FormatDuration(startTime, endTime int64) string {
	// 计算时间差（秒数）
	durationSeconds := endTime - startTime

	// 确保时间差为正数
	if durationSeconds < 0 {
		durationSeconds = -durationSeconds
	}

	// 将时间差转换为小时、分钟、秒
	hours := durationSeconds / 3600
	minutes := (durationSeconds % 3600) / 60
	seconds := durationSeconds % 60

	// 根据时间差的大小决定输出格式
	if hours > 0 {
		return fmt.Sprintf("%d小时%d分钟%d秒", hours, minutes, seconds)
	} else if minutes > 0 {
		return fmt.Sprintf("%d分钟%d秒", minutes, seconds)
	} else {
		return fmt.Sprintf("%d秒", seconds)
	}
}
