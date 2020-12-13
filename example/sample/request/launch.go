package request

import (
	dueros "github.com/dueros/bot-sdk-go/bot"
	"github.com/dueros/bot-sdk-go/bot/model"
	"github.com/yodstar/dueros-dpl/dpl"
)

func Launch(bot *dueros.Bot, request *model.LaunchRequest) {
	LOG.Debug("Launch")
	bot.Response.Tell("欢迎使用，请说开始测试")
	path := []string{
		CONF.RootDir+"/example/template/dpl-component01.json",
	}
	doc := dpl.NewDocument().
		AddParameter("payload").
		AddComponentFromFile(path...)
	if err := doc.GetLastError(); err != nil {
		bot.Response.Tell(err.Error())
		LOG.Error(err.Error())
		return
	}
	bot.Response.Command(doc.GetRenderDocument())
}
