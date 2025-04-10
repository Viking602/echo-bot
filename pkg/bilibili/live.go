package bilibili

func (c *Client) GetLiveState(roomId string) (*LiveAPI, error) {
	resp, err := c.client.R().
		SetQueryParam("room_id", roomId).
		Get("https://api.live.bilibili.com/room/v1/Room/get_info")
	if err != nil {
		return nil, err
	}
	var liveAPI LiveAPI
	err = resp.UnmarshalJson(&liveAPI)
	if err != nil {
		return nil, err
	}
	return &liveAPI, nil
}

func (c *Client) GetMasterInfo(uid string) (*MasterInfoAPI, error) {
	resp, err := c.client.R().
		SetQueryParam("uid", uid).
		Get("https://api.live.bilibili.com/live_user/v1/Master/info")
	if err != nil {
		return nil, err
	}
	var masterInfo MasterInfoAPI
	err = resp.UnmarshalJson(&masterInfo)
	if err != nil {
		return nil, err
	}
	return &masterInfo, nil
}
