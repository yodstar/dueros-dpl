package dpl

import (
	"encoding/json"
	"github.com/dueros/bot-sdk-go/bot/model"

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
	_ = json.Unmarshal([]byte(request), req)
	bot.UserEvent = req.Request
	return bot.parent.Handler(request)
}

// 添加对intent的处理函数
func (bot *Bot) AddIntentHandler(intentName string, fn func(*Bot, *model.IntentRequest)) {
	bot.parent.AddIntentHandler(intentName, func(_ *dueros.Bot, request *model.IntentRequest) {
		fn(bot, request)
	})
}

// 添加对事件的处理函数
func (bot *Bot) AddEventListener(eventName string, fn func(*Bot, *model.EventRequest)) {
	bot.parent.AddEventListener(eventName, func(_ *dueros.Bot, request interface{}) {
		eventRequest := request.(model.EventRequest)
		fn(bot, &eventRequest)
	})
}

// 添加事件默认处理函数
// 比如，在播放视频时，技能会收到各种事件的上报，如果不想一一处理可以使用这个来添加处理
func (bot *Bot) AddDefaultEventListener(fn func(*Bot, interface{})) {
	bot.parent.AddDefaultEventListener(func(_ *dueros.Bot, request interface{}) {
		fn(bot, request)
	})
}

// 打开技能时的处理
func (bot *Bot) OnLaunchRequest(fn func(*Bot, *model.LaunchRequest)) {
	bot.parent.OnLaunchRequest(func(_ *dueros.Bot, request *model.LaunchRequest) {
		fn(bot, request)
	})
}

// 技能关闭的处理，比如可以做一些清理的工作
// TIP: 根据协议，技能关闭返回的结果，DuerOS不会返回给用户。
func (bot *Bot) OnSessionEndedRequest(fn func(*Bot, *model.SessionEndedRequest)) {
	bot.parent.OnSessionEndedRequest(func(_ *dueros.Bot, request *model.SessionEndedRequest) {
		fn(bot, request)
	})
}
