package model

type OneBotMessage struct {
	Time        int64     `json:"time"`
	SelfId      int64     `json:"self_id"`
	PostType    string    `json:"post_type"`
	MessageType string    `json:"message_type"`
	SubType     string    `json:"sub_type"`
	UserId      int64     `json:"user_id"`
	MessageId   int64     `json:"message_id"`
	Message     []Message `json:"message"`
	RawMessage  string    `json:"raw_message"`
	Font        int64     `json:"font"`
	Sender      Sender    `json:"sender"`
}

type Message struct {
	Type string      `json:"type"`
	Data MessageData `json:"data"`
}

type MessageData struct {
	Text string `json:"text"`
}

type Sender struct {
	UserId   interface{} `json:"user_id"`
	Nickname string      `json:"nickname"`
	Sex      string      `json:"sex"`
	Age      int         `json:"age"`
}

type OneBotAction struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo,omitempty"`
}

type SendMessageParams struct {
	UserID     interface{} `json:"user_id,omitempty"`
	GroupID    interface{} `json:"group_id,omitempty"`
	Message    interface{} `json:"message"`
	AutoEscape bool        `json:"auto_escape"`
}
