package dpl

import (
	"encoding/json"

	dueros "github.com/dueros/bot-sdk-go/bot"
	"github.com/yodstar/dueros-dpl/dpl/event"
	"github.com/yodstar/dueros-dpl/dpl/util"
)

// 技能基础类
type Bot struct {
	dueros.Bot
	UserEvent *event.UserEvent // 设备端通过 SendEvent 指令上报
	parent    *dueros.Bot
}

// 创建常驻bot类，可维持在内存状态中, addhandler 和 addEventer事件可以缩减为一次
func NewBot() *Bot {
	bot := (*Bot)(util.UnsafeExtend(&Bot{}, dueros.NewBot()))
	bot.parent = (*dueros.Bot)(util.UnsafeConvert(bot))
	return bot
}

// 根据每个请求分别处理
func (bot *Bot) Handler(request string) string {
	req := &event.UserEventRequest{}
	json.Unmarshal([]byte(request), req)
	bot.UserEvent = req.Request
	return bot.parent.Handler(request)
}
