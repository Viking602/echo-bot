package bilibili

import (
	"fmt"
	"time"
)

func (c *Client) GetLiveState(roomId string) (*LiveAPI, error) {
	const maxRetries = 3
	const retryDelay = 1 * time.Second

	var liveAPI LiveAPI
	for attempt := 1; attempt <= maxRetries; attempt++ {
		// 发送 HTTP GET 请求
		resp, err := c.client.R().
			SetQueryParam("room_id", roomId).
			Get("https://api.live.bilibili.com/room/v1/Room/get_info")

		if err != nil {
			if attempt == maxRetries {
				return nil, fmt.Errorf("HTTP 请求失败，尝试 %d 次后仍出错: %w", maxRetries, err)
			}
			time.Sleep(retryDelay)
			continue
		}

		if resp.StatusCode != 200 {
			if attempt == maxRetries {
				return nil, fmt.Errorf("HTTP 状态码 %d，尝试 %d 次后仍失败", resp.StatusCode, maxRetries)
			}
			time.Sleep(retryDelay)
			continue
		}

		err = resp.UnmarshalJson(&liveAPI)
		if err != nil {
			if attempt == maxRetries {
				return nil, fmt.Errorf("JSON 解析失败，尝试 %d 次后仍出错: %w", maxRetries, err)
			}
			time.Sleep(retryDelay)
			continue
		}

		return &liveAPI, nil
	}

	return nil, fmt.Errorf("获取直播间状态失败，尝试 %d 次后仍无结果", maxRetries)
}

func (c *Client) GetMasterInfo(uid string) (*MasterInfoAPI, error) {
	const maxRetries = 3
	const retryDelay = 1 * time.Second

	var masterInfo MasterInfoAPI
	for attempt := 1; attempt <= maxRetries; attempt++ {
		resp, err := c.client.R().
			SetQueryParam("uid", uid).
			Get("https://api.live.bilibili.com/live_user/v1/Master/info")

		if err != nil {
			if attempt == maxRetries {
				return nil, fmt.Errorf("HTTP 状态码 %d，尝试 %d 次后仍失败", maxRetries, err)
			}
			time.Sleep(retryDelay)
			continue
		}

		if resp.StatusCode != 200 {
			if attempt == maxRetries {
				return nil, fmt.Errorf("HTTP 状态码 %d，尝试 %d 次后仍失败", resp.StatusCode, maxRetries)
			}
			time.Sleep(retryDelay)
			continue
		}

		err = resp.UnmarshalJson(&masterInfo)
		if err != nil {
			if attempt == maxRetries {
				return nil, fmt.Errorf("JSON 解析失败，尝试 %d 次后仍出错: %w", maxRetries, err)
			}
			time.Sleep(retryDelay)
			continue
		}

		return &masterInfo, nil
	}

	return nil, fmt.Errorf("获取主播信息失败，尝试 %d 次后仍无结果", maxRetries)
}
