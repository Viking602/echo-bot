package model

// Scope 定义命令适用范围
type Scope int

const (
	ScopePrivate Scope = 1 << iota // 私聊
	ScopeGroup                     // 群聊
	ScopeAll     = ScopePrivate | ScopeGroup
)

// Command 定义命令结构，支持子命令和适用范围
type Command struct {
	Name     string
	Handler  func(msg *OneBotMessage, args []string) string
	Children map[string]*Command
	Scope    Scope // 新增：命令适用范围
}

func NewCommand(name string, handler func(msg *OneBotMessage, args []string) string, scope Scope) *Command {
	return &Command{
		Name:     name,
		Handler:  handler,
		Children: make(map[string]*Command),
		Scope:    scope,
	}
}

// AddChild 添加子命令
func (c *Command) AddChild(child *Command) {
	c.Children[child.Name] = child
}

type CommandRegistrar interface {
	Register(registry Registrar)
}

type Registrar interface {
	Register(cmd *Command)
}
