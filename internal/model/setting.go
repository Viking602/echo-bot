package model

type Setting struct {
	BiliLive int64 `json:"bili_live"`
	AiChat   int64 `json:"ai_chat"`
}

type AiChatSetting struct {
	Enable bool   `json:"enable"`
	ApiKey string `json:"api_key"`
}
