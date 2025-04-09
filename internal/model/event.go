package model

type OneBotMessage struct {
	Time          int64               `json:"time"`
	SelfId        int64               `json:"self_id"`
	PostType      string              `json:"post_type"`
	MetaEventType string              `json:"meta_event_type"`
	GroupId       int64               `json:"group_id"`
	MessageType   string              `json:"message_type"`
	SubType       string              `json:"sub_type"`
	UserId        int64               `json:"user_id"`
	MessageId     int64               `json:"message_id"`
	Message       []Message           `json:"message"`
	RawMessage    string              `json:"raw_message"`
	Interval      int64               `json:"interval"`
	Status        OneBotMessageStatus `json:"status"`
	Font          int64               `json:"font"`
	Sender        Sender              `json:"sender"`
}

type OneBotMessageStatus struct {
	Online bool `json:"online"`
	Good   bool `json:"good"`
}

type Message struct {
	Type string      `json:"type"`
	Data MessageData `json:"data"`
}

type MessageData struct {
	Text string `json:"text"`
}

type Sender struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int64  `json:"age"`
	Area     string `json:"area"`
	Card     string `json:"card"`
	Role     string `json:"role"`
	Level    string `json:"level"`
	Title    string `json:"title"`
}

type OneBotAction struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
}

type SendMessageParams struct {
	UserID     int64       `json:"user_id,omitempty"`
	GroupID    int64       `json:"group_id,omitempty"`
	Message    interface{} `json:"message"`
	AutoEscape bool        `json:"auto_escape"`
}
