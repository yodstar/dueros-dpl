package request

import (
	"encoding/json"
	"fmt"

	"github.com/dueros/bot-sdk-go/bot/model"
	"github.com/yodstar/dueros-dpl/dpl"
)

func UserEvent(bot *dpl.Bot, request *model.EventRequest) {
	LOG.Debug("UserEvent")

	data, _ := json.Marshal(bot.UserEvent)
	fmt.Printf("UserEvent: %s\n", data)

	data, _ = json.Marshal(request)
	fmt.Printf("UserRequest: %s\n", data)

	bot.Response.Tell("Good!")
}
