package command

// Message Interaciton Command: 信息交互类指令 Command 类型

/*
当组件绑定 SendEvent 类型的 Command 时,点击会触发组件上报 UserEvent 事件, 用户可以自定义参数

@param componentId: 绑定的组件id

@param arguments: 动态的自定义参数

@param args[0] dialogType: 标识用户基于该事件请求上报时，携带当前会话状态的类型，枚举，取值如下
 	CURRENT（默认），多用于上报多个事件时，都能够得到响应并被执行的场景；本次事件上报基于当前对话轮次，
 	设备端在进入下一轮对话后，会将收到的基于上一轮对话的云端返回指令丢弃掉

 	NEW ，多用于场景的切换、新页面的请求，标识本次事件上报操作将开启新的一轮对话，设备端在执行该事件后
 	会更新自身的轮次状态（基于之前轮次对话事件的响应指令都将被丢弃）

@param args[1] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[2] dWhen: 指令的条件执行表达式 (optional)
*/
func SendEvent(componentId string, arguments []interface{}, args ...interface{}) *Command {
	cmd := &Command{
		Type:        "SendEvent",
		ComponentId: componentId,
		Arguments:   arguments,
	}
	argc := len(args)
	if argc > 0 {
		cmd.DialogType = args[0].(string)
	}
	if argc > 1 {
		cmd.Delay = args[1].(int)
	}
	if argc > 2 {
		cmd.DWhen = args[2].(string)
	}
	return cmd
}
