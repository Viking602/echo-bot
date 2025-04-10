package bilibili

type LiveAPI struct {
	Code    int64       `json:"code"`
	Msg     string      `json:"msg"`
	Message string      `json:"message"`
	Data    LiveAPIData `json:"data"`
}

type LiveAPIData struct {
	Uid                  int64       `json:"uid"`
	RoomId               int64       `json:"room_id"`
	ShortId              int64       `json:"short_id"`
	Attention            int64       `json:"attention"`
	Online               int64       `json:"online"`
	IsPortrait           bool        `json:"is_portrait"`
	Description          string      `json:"description"`
	LiveStatus           int64       `json:"live_status"`
	AreaId               int64       `json:"area_id"`
	ParentAreaId         int64       `json:"parent_area_id"`
	ParentAreaName       string      `json:"parent_area_name"`
	OldAreaId            int64       `json:"old_area_id"`
	Background           string      `json:"background"`
	Title                string      `json:"title"`
	UserCover            string      `json:"user_cover"`
	Keyframe             string      `json:"keyframe"`
	IsStrictRoom         bool        `json:"is_strict_room"`
	LiveTime             string      `json:"live_time"`
	Tags                 string      `json:"tags"`
	IsAnchor             int64       `json:"is_anchor"`
	RoomSilentType       string      `json:"room_silent_type"`
	RoomSilentLevel      int64       `json:"room_silent_level"`
	RoomSilentSecond     int64       `json:"room_silent_second"`
	AreaName             string      `json:"area_name"`
	Pendants             string      `json:"pendants"`
	AreaPendants         string      `json:"area_pendants"`
	HotWords             []string    `json:"hot_words"`
	HotWordsStatus       int64       `json:"hot_words_status"`
	Verify               string      `json:"verify"`
	NewPendants          NewPendants `json:"new_pendants"`
	UpSession            string      `json:"up_session"`
	PkStatus             int64       `json:"pk_status"`
	PkId                 int64       `json:"pk_id"`
	BattleId             int64       `json:"battle_id"`
	AllowChangeAreaTime  int64       `json:"allow_change_area_time"`
	AllowUploadCoverTime int64       `json:"allow_upload_cover_time"`
	StudioInfo           StudioInfo  `json:"studio_info"`
}

type NewPendants struct {
	Frame       NewPendantsFrame       `json:"frame"`
	Badge       interface{}            `json:"badge"`
	MobileFrame NewPendantsMobileFrame `json:"mobile_frame"`
	MobileBadge interface{}            `json:"mobile_badge"`
}

type NewPendantsFrame struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	Position   int64  `json:"position"`
	Desc       string `json:"desc"`
	Area       int64  `json:"area"`
	AreaOld    int64  `json:"area_old"`
	BgColor    string `json:"bg_color"`
	BgPic      string `json:"bg_pic"`
	UseOldArea bool   `json:"use_old_area"`
}

type NewPendantsMobileFrame struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	Position   int64  `json:"position"`
	Desc       string `json:"desc"`
	Area       int64  `json:"area"`
	AreaOld    int64  `json:"area_old"`
	BgColor    string `json:"bg_color"`
	BgPic      string `json:"bg_pic"`
	UseOldArea bool   `json:"use_old_area"`
}

type StudioInfo struct {
	Status     int64         `json:"status"`
	MasterList []interface{} `json:"master_list"`
}

type MasterInfoAPI struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		Info struct {
			Uid            int    `json:"uid"`
			Uname          string `json:"uname"`
			Face           string `json:"face"`
			OfficialVerify struct {
				Type int    `json:"type"`
				Desc string `json:"desc"`
			} `json:"official_verify"`
			Gender int `json:"gender"`
		} `json:"info"`
		Exp struct {
			MasterLevel struct {
				Level   int   `json:"level"`
				Color   int   `json:"color"`
				Current []int `json:"current"`
				Next    []int `json:"next"`
			} `json:"master_level"`
		} `json:"exp"`
		FollowerNum  int    `json:"follower_num"`
		RoomId       int    `json:"room_id"`
		MedalName    string `json:"medal_name"`
		GloryCount   int    `json:"glory_count"`
		Pendant      string `json:"pendant"`
		LinkGroupNum int    `json:"link_group_num"`
		RoomNews     struct {
			Content   string `json:"content"`
			Ctime     string `json:"ctime"`
			CtimeText string `json:"ctime_text"`
		} `json:"room_news"`
	} `json:"data"`
}
