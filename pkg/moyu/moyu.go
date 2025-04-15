package moyu

func (c *Client) FishAPI() (*FishAPI, error) {
	resp, err := c.client.R().Get("https://api.vvhan.com/api/moyu?type=json")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var fishAPI FishAPI
		err = resp.UnmarshalJson(&fishAPI)
		if err != nil {
			return nil, err
		}
		return &fishAPI, nil
	}

	return nil, nil

}
