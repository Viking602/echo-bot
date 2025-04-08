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
