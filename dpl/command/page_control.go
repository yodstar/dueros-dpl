package command

// Page Control Command: 模板控制类指令 Command 类型

/*
页面销毁指令

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)
*/
func PageDestroy(args ...interface{}) *Command {
	cmd := &Command{Type: "PageDestroy"}
	argc := len(args)
	if argc > 0 {
		cmd.Delay = args[0].(int)
	}
	if argc > 1 {
		cmd.DWhen = args[1].(string)
	}
	return cmd
}

/*
重新指定页面无交互退出时间

@param duration: 需要设定的新的页面无交互退出时间值，大于 0，单位毫秒

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)
*/
func ResetNonInteractionExitTime(duration int, args ...interface{}) *Command {
	cmd := &Command{
		Type:     "ResetNonInteractionExitTime",
		Duration: duration,
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

/*
监听该指令被执行时进行中的 outspeech 内容的播放状态，并基于语音播放是否完成（由bot基于自己的需求设定的
speechFinishedPosition 和下发文本的内容长度进行比较判断，监听中获取到的递增 ttsposition 语音播报进度
第一次达到 speechFinishedPosition 时认为完成监听；监听中存在 ttsposition 语音播报被打断、被抢占、监听
持续 500ms 内未监听 ttsposition 语音播报进行则认为失败），并执行相应事件的指令回调（onFinished，onInterrupted）

@param componentId: 固定为 DPL，作为模板指令被执行

@param speechFinishedPosition: 标识为 speech 播放完成需要达到的 ttsposition 值（必填，且大于0），speech 播报中 ttsposition 位置值在逐增过程中第一次达到 speechFinishedPosition 值时，执行 onFinished 并结束本次监听

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)

@param args[2] onFinished: 标识为播放完成后可执行的指令（目前限定为仅可使用 SendEvent 指令反馈监听结果） (optional)

@param args[3] onInterrupted: 标识为播放被中断后可执行的指令（目前限定为仅可使用 SendEvent 指令反馈监听结果） (optional)
*/
func SetSpeechMonitor(componentId string, speechFinishedPosition int, args ...interface{}) *Command {
	cmd := &Command{
		Type:                   "SetSpeechMonitor",
		ComponentId:            componentId,
		SpeechFinishedPosition: speechFinishedPosition,
	}
	argc := len(args)
	if argc > 0 {
		cmd.Delay = args[0].(int)
	}
	if argc > 1 {
		cmd.DWhen = args[1].(string)
	}
	if argc > 2 {
		cmd.OnFinished = args[2].(*Command)
	}
	if argc > 3 {
		cmd.OnInterrupted = args[3].(*Command)
	}
	return cmd
}
