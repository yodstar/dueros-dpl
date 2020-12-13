package command

// Date Updating Command

/*
更新引用数据源的部分数据内容，会使 dpl 中通过动态引用方式使用该数据的对应内容同步更新

@param data: 与 dataSource 保持相同结构的需要更新的数据内容，该指令被执行时，data 内容会被 merge 更新到原 dataSource 中

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)
*/
func UpdateDataSource(data map[string]interface{}, args ...interface{}) *Command {
	cmd := &Command{Type: "UpdateDataSource", Data: data}
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
修改/新增组件的可使用 style 样式属性

@param componentId: 绑定的组件id

@param styles: 样式属性对象: {key: 样式属性名称, value: 样式属性值}

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)
*/
func SetStyles(componentId string, styles map[string]interface{}, args ...interface{}) *Command {
	cmd := &Command{
		Type:        "SetStyles",
		ComponentId: componentId,
		Styles:      styles,
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
修改组件的可使用 prop 功能属性

@param componentId: 绑定的组件id

@param props: 功能属性对象: {key: 功能属性名称, value: 功能属性值}

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)
*/
func SetProps(componentId string, props map[string]interface{}, args ...interface{}) *Command {
	cmd := &Command{
		Type:        "SetProps",
		ComponentId: componentId,
		Props:       props,
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
改变组件的部分可供修改的属性值，注：目前该属性仅支持通过事件绑定使用 (不再建议使用，推荐使用如上的 SetStyles、SetProps 指令来替代它)

@param componentId: 绑定的组件id

@param state: 属性名称

@param value: 对应的属性值, 可为空

@param args[0] stateType: 属性类型，枚举值: PROP 组件属性（默认）, STYLE 样式属性 (optional)

@param args[1] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[2] dWhen: 指令的条件执行表达式 (optional)
*/
func SetState(componentId, state, value string, args ...interface{}) *Command {
	cmd := &Command{
		Type:        "SetState",
		ComponentId: componentId,
		State:       state,
		Value:       value,
	}
	argc := len(args)
	if argc > 0 {
		cmd.StateType = args[0].(string)
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
（该指令只能通过云端下发指令的方式使用）异步刷新指令, 使用指定的 document 描述的组件替换指定的 componentId 对应的组件,
document 的语法结构和 DPL_DOCUMENT 一致

@param componentId: 需要被替换内容的组件 id

@param document: 更新的组件内容 结构参考：DPL_DOCUMENT

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)
*/
func UpdateComponent(componentId string, document interface{}, args ...interface{}) *Command {
	cmd := &Command{
		Type:        "UpdateComponent",
		ComponentId: componentId,
		Document:    document,
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
（该指令只能通过云端下发指令的方式使用）向模板/组件中追加子组件

@param componentId: 追加的父组件 id

@param document: 追加的组件内容 结构参考：DPL_DOCUMENT

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)
*/
func AppendComponent(componentId string, document interface{}, args ...interface{}) *Command {
	cmd := &Command{
		Type:        "AppendComponent",
		ComponentId: componentId,
		Document:    document,
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
