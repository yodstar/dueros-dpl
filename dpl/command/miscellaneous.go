package command

// Miscellaneous Command

/*
通用方法调用指令

@param componentId: 绑定的组件id

@param methodName: 动态的自定义参数

@param params: 指令参数: arguments 指令调用方法的接收参数，数组类型

 	arguments数组中参数类型支持任意类型（string、object、array、number、boolean）；

 	arguments数组中参数类型支持获取事件触发的绑定参数（在事件触发指令执行的场景下），如："arguments":["$EVENT_ARGS.0", 1],
 	代表该指令会将事件触发该指令时，事件触发携带的参数的第一个，作为arguments的第一个参数（其结构支持多层级取值）

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)
*/
func InvokeMethod(componentId, methodName string, params []interface{}, args ...interface{}) *Command {
	cmd := &Command{
		Type:        "InvokeMethod",
		ComponentId: componentId,
		MethodName:  methodName,
		Params:      params,
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
触发组件动画, 按照设置的属性展示动画效果，目前可作用在 Image, Text, Container 三类组件上

@param componentId: 绑定的组件 id, 当前只在几种有限的 component 上生效：Text、Image, Container

@param attribute: 动画所作用的属性。视图组件中所有属性值可以线性变化的，都可以作为动画作用的属性，枚举如下：
 	scaleX 通过设置 X 轴的值来定义缩放转换
 	scaleY 通过设置 Y 轴的值来定义缩放转换
 	rotation 旋转

@param from: 动画作用属性的起始值。可以是 number 类型，也可以是 length，也可以是 percentage，具体由 attribute 对应的值的类型决定

@param to: 动画作用属性的结束值

@param easing: 描述动画执行的速度曲线，用于使动画变化更为平滑。默认值是 linear，表示动画从开始到结束都拥有同样的速度，枚举, 取值如下
 	linear 动画从头到尾的速度是相同的
 	ease 动画速度逐渐变慢
 	ease-in 动画速度由慢到快
 	ease-out 动画速度由快到慢
 	shake 抖动动画 取值：shake(power,speed,loop,time), 如: shake(100, 20, 0, 300)
 		当取值为 shake 时, 其他属性失效
 		power 振动的幅度
 		speed 振动的速度
 		loop 取值 0 或 1，0 表示不循环执行一次，1 表示循环执行
 		time 振动的次数, 默认是 20ms 执行一次
 	cubic-bezier(x1, y1, x2, y2) 在三次贝塞尔函数中定义变化过程，函数的参数值必须处于 0 到 1 之间 更多关于三次贝塞尔的信息请参阅 cubic-bezier 和 Bézier curve

@param duration: 动画执行的时间

@param args[0] repeatCount: 可选，动画重复的次数。填'infinite'为无限循环 填字符串形式的数字为重复次数 如若填'3'，则重复 3 次 一共 4 次

@param args[1] repeatMode: 动画重复方式。restart 为从一直头播放；reverse 为从头到尾再倒回来。默认 restart

@param args[2] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[3] dWhen: 指令的条件执行表达式 (optional)
*/
func Animation(componentId, attribute, from, to, easing string, duration int, args ...interface{}) *Command {
	cmd := &Command{
		Type:        "Animation",
		ComponentId: componentId,
		Attribute:   attribute,
		From:        from,
		To:          to,
		Easing:      easing,
		Duration:    duration,
	}
	argc := len(args)
	if argc > 0 {
		cmd.RepeatCount = args[0].(string)
	}
	if argc > 1 {
		cmd.RepeatMode = args[1].(string)
	}
	if argc > 2 {
		cmd.Delay = args[2].(int)
	}
	if argc > 3 {
		cmd.DWhen = args[3].(string)
	}
	return cmd
}

/*
触发组件动画, 按照设置的属性展示动画效果，目前可作用在 Image, Text, Container 三类组件上

@param componentId: 绑定的组件 id

@param repeatCount: 重复执行次数

@param delayInMilliseconds: 重复执行间隔时间, 单位毫秒

@param commands: 指令数组，可以是 AutoPage、Scroll... 等等指令

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)
*/
func Parallel(componentId string, repeatCount, delayInMilliseconds int, commands []*Command, args ...interface{}) *Command {
	cmd := &Command{
		Type:                "Parallel",
		ComponentId:         componentId,
		RepeatCount:         repeatCount,
		DelayInMilliseconds: delayInMilliseconds,
		Commands:            commands,
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
