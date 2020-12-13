package request

import (
	dueros "github.com/dueros/bot-sdk-go/bot"
	"github.com/yodstar/goutil/logger"

	"sample/config"
)

var (
	CONF = config.CONF
	LOG  = logger.LOG
)

func Default(bot *dueros.Bot, request interface{}) {
	LOG.Debug("Default")
	
	bot.Response.HoldOn()
	bot.Response.CloseMicrophone()
}
