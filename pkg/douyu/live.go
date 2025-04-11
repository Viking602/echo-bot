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
	return &searchAPI, nil
}
