package dpl

import (
	"github.com/dueros/bot-sdk-go/bot/directive"
	"github.com/yodstar/dueros-dpl/dpl/command"
)

type ExecuteCommands struct {
	directive.BaseDirective                    // type: 指令类型，此处应为：DPL.ExecuteCommands
	Token                   string             `json:"token"`    // token: 指令 token 值需要 match 匹配基于用户请求时的页面状态 token
	Commands                []*command.Command `json:"commands"` // 指定执行的 commands 数组
}

func (e *ExecuteCommands) SetToken(token string) {
	e.Token = token
}

func (e *ExecuteCommands) SetCommands(cmds ...*command.Command) {
	e.Commands = cmds
}

func NewExecuteCommands(cmds ...*command.Command) *ExecuteCommands {
	e := &ExecuteCommands{}
	e.Type = "DPL.ExecuteCommands"
	e.Token = e.GenToken()
	e.Commands = cmds
	return e
}
