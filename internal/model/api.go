package model

type Resp struct {
	Status  string   `json:"status"`
	Retcode int64    `json:"retcode"`
	Data    RespData `json:"data"`
}

type RespData struct {
	MessageId int64  `json:"message_id"`
	Message   string `json:"message"`
}
