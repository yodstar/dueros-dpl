package request

import (
	"encoding/json"
	"fmt"
	"unsafe"

	dueros "github.com/dueros/bot-sdk-go/bot"
	"github.com/dueros/bot-sdk-go/bot/model"
	"github.com/yodstar/dueros-dpl/dpl"
)

func UserEvent(bot *dueros.Bot, request interface{}) {
	LOG.Debug("UserEvent")

	newBot := (*dpl.Bot)(unsafe.Pointer(bot))
	data, _ := json.Marshal(newBot.UserEvent)
	fmt.Printf("UserEvent: %s\n", data)

	data, _ = json.Marshal(request.(*model.UserEventRequest))
	fmt.Printf("UserEventRequest: %s\n", data)

	bot.Response.Tell("Good!")
}
