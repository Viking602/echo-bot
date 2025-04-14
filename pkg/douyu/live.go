package douyu

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
	resp, err := c.client.R().
		SetHeaders(map[string]string{
			"Referer": "https://www.douyu.com/" + roomId,
		}).
		Get("https://open.douyucdn.cn/api/RoomApi/room/" + roomId)
	if err != nil {
		return nil, err
	}
	var roomAPI RoomAPI
	err = resp.UnmarshalJson(&roomAPI)
	if err != nil {
		return nil, err
	}
	return &roomAPI, nil
}
