package model

// Command 定义指令结构
type Command struct {
	Name    string
	Handler func(msg *OneBotMessage) string
}

// CommandRegistrar 定义指令注册器接口
type CommandRegistrar interface {
	Register(registry Registrar)
}

// Registrar 定义注册表接口（供 command 使用）
type Registrar interface {
	Register(cmd Command)
}
