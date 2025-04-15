package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// Crontab 结构体表示解析后的 crontab 表达式
type Crontab struct {
	Minutes     []int // 分钟 (0-59)
	Hours       []int // 小时 (0-23)
	DaysOfMonth []int // 每月日期 (1-31)
	Months      []int // 月份 (1-12)
	DaysOfWeek  []int // 星期 (0-6, 0=Sunday)
}

// parseRange 解析类似 "1-5" 或 "*/2" 的范围表达式
func parseRange(field string, min, max int) ([]int, error) {
	var result []int

	if strings.HasPrefix(field, "*/") {
		step, err := strconv.Atoi(field[2:])
		if err != nil {
			return nil, fmt.Errorf("invalid step: %v", err)
		}
		for i := min; i <= max; i += step {
			result = append(result, i)
		}
		return result, nil
	}

	if strings.Contains(field, "-") {
		parts := strings.Split(field, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid range: %s", field)
		}
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid start: %v", err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid end: %v", err)
		}
		if start < min || end > max || start > end {
			return nil, fmt.Errorf("range out of bounds: %s", field)
		}
		for i := start; i <= end; i++ {
			result = append(result, i)
		}
		return result, nil
	}

	num, err := strconv.Atoi(field)
	if err != nil {
		return nil, fmt.Errorf("invalid number: %v", err)
	}
	if num < min || num > max {
		return nil, fmt.Errorf("number out of bounds: %d", num)
	}
	return []int{num}, nil
}

// parseField 解析 crontab 的一个字段
func parseField(field string, min, max int) ([]int, error) {
	var result []int
	seen := make(map[int]bool)

	if field == "*" {
		for i := min; i <= max; i++ {
			result = append(result, i)
		}
		return result, nil
	}

	parts := strings.Split(field, ",")
	for _, part := range parts {
		values, err := parseRange(part, min, max)
		if err != nil {
			return nil, err
		}
		for _, v := range values {
			if !seen[v] {
				result = append(result, v)
				seen[v] = true
			}
		}
	}

	return result, nil
}

// ParseCrontab 解析 crontab 表达式
func ParseCrontab(expr string) (*Crontab, error) {
	fields := strings.Fields(expr)
	if len(fields) != 5 {
		return nil, fmt.Errorf("invalid crontab format: expected 5 fields, got %d", len(fields))
	}

	cron := &Crontab{}

	var err error
	cron.Minutes, err = parseField(fields[0], 0, 59)
	if err != nil {
		return nil, fmt.Errorf("invalid minutes: %v", err)
	}
	cron.Hours, err = parseField(fields[1], 0, 23)
	if err != nil {
		return nil, fmt.Errorf("invalid hours: %v", err)
	}
	cron.DaysOfMonth, err = parseField(fields[2], 1, 31)
	if err != nil {
		return nil, fmt.Errorf("invalid days of month: %v", err)
	}
	cron.Months, err = parseField(fields[3], 1, 12)
	if err != nil {
		return nil, fmt.Errorf("invalid months: %v", err)
	}
	cron.DaysOfWeek, err = parseField(fields[4], 0, 6)
	if err != nil {
		return nil, fmt.Errorf("invalid days of week: %v", err)
	}

	return cron, nil
}

// ToNaturalString 将 Crontab 转换为简洁的自然语言描述
func (c *Crontab) ToNaturalString() string {
	var parts []string

	// 分钟和小时
	if len(c.Minutes) == 1 && len(c.Hours) == 1 {
		parts = append(parts, fmt.Sprintf("%d:%02d", c.Hours[0], c.Minutes[0]))
	} else if len(c.Hours) == 1 {
		parts = append(parts, fmt.Sprintf("%d 点", c.Hours[0]))
	} else {
		return "复杂时间表"
	}

	// 日期和星期
	if len(c.DaysOfMonth) == 31 && len(c.DaysOfWeek) == 7 {
		parts = append(parts, "每天")
	} else if len(c.DaysOfMonth) == 1 && len(c.DaysOfWeek) == 7 {
		parts = append(parts, fmt.Sprintf("每月 %d 号", c.DaysOfMonth[0]))
	} else if len(c.DaysOfWeek) == 1 && len(c.DaysOfMonth) == 31 {
		parts = append(parts, fmt.Sprintf("每周 %d", c.DaysOfWeek[0]))
	} else {
		return "复杂时间表"
	}

	// 月份
	if len(c.Months) == 12 {
		// 不额外添加月份信息
	} else if len(c.Months) == 1 {
		parts = append(parts, fmt.Sprintf("%d 月", c.Months[0]))
	} else {
		return "复杂时间表"
	}

	// 拼接并添加“运行”
	return strings.Join(parts, "") + "运行"
}
