# dueros-dpl
DPL(DuerOS Presentation Language) on bot-sdk-go

## Example:

- Example 1

```
import (
	...

	"github.com/yodstar/dueros-dpl/dpl"
)


func launchRequest(bot *dueros.Bot, request *model.LaunchRequest) {
	bot.Response.Tell("欢迎使用，请说开始测试")

	doc, err := dpl.NewDocument("./tempalte/dpl-tempalte1.json")
	if err != nil {
		bot.Response.Tell("查询失败")
		log.Println(err.Error())
		return
	}
	bot.Response.Command(doc.GetRenderDocument())
}

func main() {
	...
}

```
