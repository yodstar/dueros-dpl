package command

// Component-specified Command: 作用在专门组件可使用的 Command 类型

/*
实现 ScrollView 和 List 的窗口滑动，将 List 当前展现的窗口视图向上或者向下滑动指定的距离

@param componentId: 绑定的组件 id

@param distance: 每次窗口滑动的距离，如 20dp 或者 10%。负值如 -20dp 或者 -10%，表示向下滑动

@param args[0] duration: 滚动动画持续时间，单位毫秒，默认 500 毫秒 (optional)

@param args[1] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[2] dWhen: 指令的条件执行表达式 (optional)
*/
func Scroll(componentId, distance string, args ...interface{}) *Command {
	cmd := &Command{
		Type:        "Scroll",
		ComponentId: componentId,
		Distance:    distance,
	}
	argc := len(args)
	if argc > 0 {
		cmd.Duration = args[0].(int)
	}
	if argc > 1 {
		cmd.Delay = args[1].(int)
	}
	if argc > 2 {
		cmd.DWhen = args[2].(string)
	}
	return cmd
}

/*
将当前的 List 视图滑动到指定 item 所在的列表项处, index 表示列表项的序号

@param componentId: 绑定的组件 id

@param index: 指定的 item 的 index, 如有 N 个 item 取值 0 到 N-1

@param align: 枚举, 取值如下:

 	first 指定的 index 处于屏幕顶部第一个或者最左边位置
 	center 指定的 index 处于屏幕中间位置
 	last 指定的 index 处于屏幕底部或者最右边位置

@param args[0] duration: 滚动动画持续时间，单位毫秒，默认 500 毫秒 (optional)

@param args[1] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[2] dWhen: 指令的条件执行表达式 (optional)
*/
func ScrollToIndex(componentId string, index int, align string, args ...interface{}) *Command {
	cmd := &Command{
		Type:        "ScrollToIndex",
		ComponentId: componentId,
		Index:       &index,
		Align:       align,
	}
	argc := len(args)
	if argc > 0 {
		cmd.Duration = args[0].(int)
	}
	if argc > 1 {
		cmd.Delay = args[1].(int)
	}
	if argc > 2 {
		cmd.DWhen = args[2].(string)
	}
	return cmd
}

/*
将当前的列表视图（ScrollView、List）滑动到指定 item 所在的列表项处, targetComponentId 表示列表的对应目标子项

@param componentId: 绑定的组件 id

@param targetComponentId: 指定的 item 的 componentId, 需明确指明

@param align: 枚举, 取值如下:

 	first 指定的 index 处于屏幕顶部第一个或者最左边位置
 	center 指定的 index 处于屏幕中间位置
 	last 指定的 index 处于屏幕底部或者最右边位置

@param args[0] duration: 滚动动画持续时间，单位毫秒，默认 500 毫秒 (optional)

@param args[1] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[2] dWhen: 指令的条件执行表达式 (optional)
*/
func ScrollToElement(componentId, targetComponentId, align string, args ...interface{}) *Command {
	cmd := &Command{
		Type:              "ScrollToElement",
		ComponentId:       componentId,
		TargetComponentId: targetComponentId,
		Align:             align,
	}
	argc := len(args)
	if argc > 0 {
		cmd.Duration = args[0].(int)
	}
	if argc > 1 {
		cmd.Delay = args[1].(int)
	}
	if argc > 2 {
		cmd.DWhen = args[2].(string)
	}
	return cmd
}

/*
实现 pager 组件的的自动翻页功能, 最终停留在最后的页面

@param componentId: 绑定的组件 id

@param durationInMillisecond: 页面翻页时切换的时间间隔，单位毫秒

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)
*/
func AutoPage(componentId string, durationInMillisecond int, args ...interface{}) *Command {
	cmd := &Command{
		Type:                  "AutoPage",
		ComponentId:           componentId,
		DurationInMillisecond: durationInMillisecond,
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
将 pager 组件的视图切换到指定目录页面，index 值表示 pager 的页面序号

@param componentId: 绑定的组件 id

@param position: 枚举, 取值如下：relative 相对位置，absolute 绝对位置

@param value: 翻页步长, 假如有 N 个页, 取值 0 到 N-1

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)
*/
func SetPage(componentId, position string, value int, args ...interface{}) *Command {
	cmd := &Command{
		Type:        "SetPage",
		ComponentId: componentId,
		Position:    position,
		Value:       &value,
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
