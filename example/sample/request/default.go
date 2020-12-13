package request

import (
	"github.com/yodstar/dueros-dpl/dpl"
	"github.com/yodstar/goutil/logger"

	"sample/config"
)

var (
	CONF = config.CONF
	LOG  = logger.LOG
)

func Default(bot *dpl.Bot, request interface{}) {
	LOG.Debug("Default")
	
	bot.Response.HoldOn()
	bot.Response.CloseMicrophone()
}
