package douyu

import (
	"fmt"
	"time"
)

func (c *Client) Search(roomId string) (*SearchAPI, error) {
	resp, err := c.client.R().
		SetQueryParams(map[string]string{
			"kw":         roomId,
			"tagTest":    "a",
			"client_sys": "android#",
		}).
		Get("https://apiv2.douyucdn.cn/japi/search/api/getSearchRec")
	if err != nil {
		return nil, err
	}
	var searchAPI SearchAPI
	err = resp.UnmarshalJson(&searchAPI)
	if err != nil {
		return nil, err
	}

	if searchAPI.Error != 0 {
		c.log.Error().Str("url", resp.Request.URL.String())
	}

	return &searchAPI, nil
}

func (c *Client) GetRoomInfo(roomId string) (*RoomAPI, error) {
	const maxRetries = 3
	const retryDelay = 1 * time.Second

	var roomAPI RoomAPI
	for attempt := 1; attempt <= maxRetries; attempt++ {
		// 发送 HTTP GET 请求
		resp, err := c.client.R().
			SetHeaders(map[string]string{
				"Referer": "https://www.douyu.com/" + roomId,
			}).
			Get("https://open.douyucdn.cn/api/RoomApi/room/" + roomId)

		// 检查 HTTP 请求错误
		if err != nil {
			if attempt == maxRetries {
				return nil, fmt.Errorf("HTTP 请求失败，尝试 %d 次后仍出错: %w", maxRetries, err)
			}
			time.Sleep(retryDelay)
			continue
		}

		// 检查 HTTP 状态码
		if resp.StatusCode != 200 {
			if attempt == maxRetries {
				return nil, fmt.Errorf("HTTP 状态码 %d，尝试 %d 次后仍失败", resp.StatusCode, maxRetries)
			}
			time.Sleep(retryDelay)
			continue
		}

		// 尝试解析 JSON
		err = resp.UnmarshalJson(&roomAPI)
		if err != nil {
			if attempt == maxRetries {
				return nil, fmt.Errorf("JSON 解析失败，尝试 %d 次后仍出错: %w", maxRetries, err)
			}
			time.Sleep(retryDelay)
			continue
		}

		// 成功：返回结果
		return &roomAPI, nil
	}

	// 理论上不会到达此行，但为完整性保留
	return nil, fmt.Errorf("获取直播间信息失败，尝试 %d 次后仍无结果", maxRetries)
}
