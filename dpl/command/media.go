package command

// Media Command

/*
该指令适用在媒体内容上执行动作，以达到完成功能切换的目的

@param componentId: 固定为 DPL，作为模板指令被执行

@param command: 要执行的具体操作，可选值如下表
 	componetType		command		description
 	Video         		next		下一个
 	Video         		previous	上一个
 	Video         		play		播放中
 	Video         		pause		暂停
 	Video         		end 		播放结束
 	Audio         		next		下一个
 	Audio         		previous	上一个
 	Audio         		play		播放中
 	Audio         		pause		暂停
 	Audio         		end 		播放结束
 	Timer         		reset		重置定时器
 	Timer         		pause		暂停定时器
 	Timer         		play		执行定时器
 	Counter	      		reset		重置计时器
 	Counter       		pause		暂停计时器
 	Counter       		play		执行计时器
 	FrameAnimation		reset		重置逐帧动画
 	FrameAnimation		pause		暂停逐帧动画
 	FrameAnimation		play		执行逐帧动画

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)
*/
func ControlMedia(componentId, command string, args ...interface{}) *Command {
	cmd := &Command{
		Type:        "ControlMedia",
		ComponentId: componentId,
		Command:     command,
	}
	argc := len(args)
	if argc > 0 {
		cmd.Delay = args[0].(int)
	}
	if argc > 1 {
		cmd.DWhen = args[1].(string)
	}
	return cmd
}
