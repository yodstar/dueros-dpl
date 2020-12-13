package startup

import (
	dueros "github.com/dueros/bot-sdk-go/bot"
	"github.com/yodstar/dueros-dpl/dpl"
	"github.com/yodstar/goutil/logger"

	"sample/config"
	"sample/request"
)

var (
	CONF = config.CONF
	LOG  = logger.LOG
)

func Run() {
	// dueros
	bot := dpl.NewBot()
	bot.AddDefaultEventListener(request.Default)
	bot.AddEventListener("UserEvent", request.UserEvent)
	bot.OnLaunchRequest(request.Launch)

	LOG.Info("Listen: %s", CONF.Listen)

	app := dueros.Application{AppId: CONF.DuerOS.AppID, Handler: bot.Handler}
	app.Start(CONF.Listen)
}
