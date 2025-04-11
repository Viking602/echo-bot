package douyu

type SearchAPI struct {
	Data struct {
		RecList []struct {
			Algorithm struct {
				Rt         string `json:"rt"`
				RecallType string `json:"_recall_type"`
			} `json:"algorithm"`
			BkUrl    string `json:"bkUrl"`
			Kw       string `json:"kw"`
			RoomInfo struct {
				Avatar         string `json:"avatar"`
				BkUrl          string `json:"bkUrl"`
				CateName       string `json:"cateName"`
				Cid            int64  `json:"cid"`
				DesType        int64  `json:"desType"`
				DesVersion     string `json:"desVersion"`
				Description    string `json:"description"`
				DescriptionV2  string `json:"descriptionV2"`
				FansNum        int64  `json:"fansNum"`
				FansNumStr     string `json:"fansNumStr"`
				FeedNum        int64  `json:"feedNum"`
				FeedNumStr     string `json:"feedNumStr"`
				FollowerCount  string `json:"followerCount"`
				IsLive         int64  `json:"isLive"`
				IsLoop         int64  `json:"isLoop"`
				IsVertical     int64  `json:"isVertical"`
				LastShowTime   int64  `json:"lastShowTime"`
				NickName       string `json:"nickName"`
				Rid            int64  `json:"rid"`
				RoomSrc        string `json:"roomSrc"`
				RoomType       int64  `json:"roomType"`
				Tag            string `json:"tag"`
				Tid            int64  `json:"tid"`
				Url            string `json:"url"`
				VideoSchemeUrl string `json:"videoSchemeUrl"`
				VipId          int64  `json:"vipId"`
			} `json:"roomInfo"`
			SchemeUrl string `json:"schemeUrl"`
			Type      int64  `json:"type"`
		} `json:"recList"`
	} `json:"data"`
	Error int    `json:"error"`
	Msg   string `json:"msg"`
}
