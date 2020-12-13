

# command
`import "github.com/yodstar/dueros-dpl/dpl/command"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type Command](#Command)
  * [func Animation(componentId, attribute, from, to, easing string, duration int, args ...interface{}) *Command](#Animation)
  * [func AppendComponent(componentId string, document interface{}, args ...interface{}) *Command](#AppendComponent)
  * [func AutoPage(componentId string, durationInMillisecond int, args ...interface{}) *Command](#AutoPage)
  * [func ControlMedia(componentId, command string, args ...interface{}) *Command](#ControlMedia)
  * [func InvokeMethod(componentId, methodName string, params []interface{}, args ...interface{}) *Command](#InvokeMethod)
  * [func PageDestroy(args ...interface{}) *Command](#PageDestroy)
  * [func Parallel(componentId string, repeatCount, delayInMilliseconds int, commands []*Command, args ...interface{}) *Command](#Parallel)
  * [func ResetNonInteractionExitTime(duration int, args ...interface{}) *Command](#ResetNonInteractionExitTime)
  * [func Scroll(componentId, distance string, args ...interface{}) *Command](#Scroll)
  * [func ScrollToElement(componentId, targetComponentId, align string, args ...interface{}) *Command](#ScrollToElement)
  * [func ScrollToIndex(componentId string, index int, align string, args ...interface{}) *Command](#ScrollToIndex)
  * [func SendEvent(componentId string, arguments []interface{}, args ...interface{}) *Command](#SendEvent)
  * [func SetPage(componentId, position string, value int, args ...interface{}) *Command](#SetPage)
  * [func SetProps(componentId string, props map[string]interface{}, args ...interface{}) *Command](#SetProps)
  * [func SetSpeechMonitor(componentId string, speechFinishedPosition int, args ...interface{}) *Command](#SetSpeechMonitor)
  * [func SetState(componentId, state, value string, args ...interface{}) *Command](#SetState)
  * [func SetStyles(componentId string, styles map[string]interface{}, args ...interface{}) *Command](#SetStyles)
  * [func UpdateComponent(componentId string, document interface{}, args ...interface{}) *Command](#UpdateComponent)
  * [func UpdateDataSource(data map[string]interface{}, args ...interface{}) *Command](#UpdateDataSource)
  * [func (cmd *Command) GetCommandFromFile(path ...string) *Command](#Command.GetCommandFromFile)
  * [func (cmd *Command) Marshal() ([]byte, error)](#Command.Marshal)
  * [func (cmd *Command) SetOnFinished(onFinished *Command) *Command](#Command.SetOnFinished)
  * [func (cmd *Command) SetOnInterrupted(onInterrupted *Command) *Command](#Command.SetOnInterrupted)
  * [func (cmd *Command) Unmarshal(data []byte) error](#Command.Unmarshal)


#### <a name="pkg-files">Package files</a>
[command.go](/src/github.com/yodstar/dueros-dpl/dpl/command/command.go) [component_specified.go](/src/github.com/yodstar/dueros-dpl/dpl/command/component_specified.go) [data_updating.go](/src/github.com/yodstar/dueros-dpl/dpl/command/data_updating.go) [media.go](/src/github.com/yodstar/dueros-dpl/dpl/command/media.go) [message_interaciton.go](/src/github.com/yodstar/dueros-dpl/dpl/command/message_interaciton.go) [miscellaneous.go](/src/github.com/yodstar/dueros-dpl/dpl/command/miscellaneous.go) [page_control.go](/src/github.com/yodstar/dueros-dpl/dpl/command/page_control.go) 






## <a name="Command">type</a> [Command](/src/target/command.go?s=110:2452#L10)
``` go
type Command struct {
    Error       error  `json:"-"`
    Type        string `json:"type"`                  // type: 指令类型
    ComponentId string `json:"componentId,omitempty"` // componentId: 组件 id
    Delay       int    `json:"delay,omitempty"`       // delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)
    DWhen       string `json:"dWhen,omitempty"`       // dWhen: 指令的条件执行表达式 (optional)

    Align     string        `json:"align,omitempty"`
    Arguments []interface{} `json:"arguments,omitempty"`
    Attribute string        `json:"attribute,omitempty"`

    Command  string     `json:"command,omitempty"`
    Commands []*Command `json:"commands,omitempty"`

    Data                  map[string]interface{} `json:"data,omitempty"`
    DelayInMilliseconds   int                    `json:"delayInMilliseconds,omitempty"`
    DialogType            string                 `json:"dialogType,omitempty"`
    Distance              string                 `json:"distance,omitempty"`
    Document              interface{}            `json:"document,omitempty"`
    Duration              int                    `json:"duration,omitempty"`
    DurationInMillisecond int                    `json:"durationInMillisecond,omitempty"`

    Easing     string `json:"easing,omitempty"`
    From       string `json:"from,omitempty"`
    Index      *int   `json:"index,omitempty"`
    MethodName string `json:"methodName,omitempty"`

    OnFinished    *Command `json:"onFinished,omitempty"`
    OnInterrupted *Command `json:"onInterrupted,omitempty"`

    Params   []interface{}          `json:"params,omitempty"`
    Position string                 `json:"position,omitempty"`
    Props    map[string]interface{} `json:"props,omitempty"`

    RepeatCount interface{} `json:"repeatCount,omitempty"`
    RepeatMode  string      `json:"repeatMode,omitempty"`

    SpeechFinishedPosition int                    `json:"speechFinishedPosition,omitempty"`
    Styles                 map[string]interface{} `json:"styles,omitempty"`
    StateType              string                 `json:"stateType,omitempty"`
    State                  string                 `json:"state,omitempty"`

    TargetComponentId string `json:"targetComponentId,omitempty"`
    To                string `json:"to,omitempty"`

    Value interface{} `json:"value,omitempty"`
}

```






### <a name="Animation">func</a> [Animation](/src/target/miscellaneous.go?s=3386:3493#L77)
``` go
func Animation(componentId, attribute, from, to, easing string, duration int, args ...interface{}) *Command
```
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


### <a name="AppendComponent">func</a> [AppendComponent](/src/target/data_updating.go?s=4434:4526#L154)
``` go
func AppendComponent(componentId string, document interface{}, args ...interface{}) *Command
```
（该指令只能通过云端下发指令的方式使用）向模板/组件中追加子组件

@param componentId: 追加的父组件 id

@param document: 追加的组件内容 结构参考：DPL_DOCUMENT

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)


### <a name="AutoPage">func</a> [AutoPage](/src/target/component_specified.go?s=3790:3880#L126)
``` go
func AutoPage(componentId string, durationInMillisecond int, args ...interface{}) *Command
```
实现 pager 组件的的自动翻页功能, 最终停留在最后的页面

@param componentId: 绑定的组件 id

@param durationInMillisecond: 页面翻页时切换的时间间隔，单位毫秒

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)


### <a name="ControlMedia">func</a> [ControlMedia](/src/target/media.go?s=1214:1290#L36)
``` go
func ControlMedia(componentId, command string, args ...interface{}) *Command
```
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


### <a name="InvokeMethod">func</a> [InvokeMethod](/src/target/miscellaneous.go?s=859:960#L23)
``` go
func InvokeMethod(componentId, methodName string, params []interface{}, args ...interface{}) *Command
```
通用方法调用指令

@param componentId: 绑定的组件id

@param methodName: 动态的自定义参数

@param params: 指令参数: arguments 指令调用方法的接收参数，数组类型


	arguments数组中参数类型支持任意类型（string、object、array、number、boolean）；
	
	arguments数组中参数类型支持获取事件触发的绑定参数（在事件触发指令执行的场景下），如："arguments":["$EVENT_ARGS.0", 1],
	代表该指令会将事件触发该指令时，事件触发携带的参数的第一个，作为arguments的第一个参数（其结构支持多层级取值）

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)


### <a name="PageDestroy">func</a> [PageDestroy](/src/target/page_control.go?s=281:327#L12)
``` go
func PageDestroy(args ...interface{}) *Command
```
页面销毁指令

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)


### <a name="Parallel">func</a> [Parallel](/src/target/miscellaneous.go?s=4482:4604#L118)
``` go
func Parallel(componentId string, repeatCount, delayInMilliseconds int, commands []*Command, args ...interface{}) *Command
```
触发组件动画, 按照设置的属性展示动画效果，目前可作用在 Image, Text, Container 三类组件上

@param componentId: 绑定的组件 id

@param repeatCount: 重复执行次数

@param delayInMilliseconds: 重复执行间隔时间, 单位毫秒

@param commands: 指令数组，可以是 AutoPage、Scroll... 等等指令

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)


### <a name="ResetNonInteractionExitTime">func</a> [ResetNonInteractionExitTime](/src/target/page_control.go?s=825:901#L33)
``` go
func ResetNonInteractionExitTime(duration int, args ...interface{}) *Command
```
重新指定页面无交互退出时间

@param duration: 需要设定的新的页面无交互退出时间值，大于 0，单位毫秒

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)


### <a name="Scroll">func</a> [Scroll](/src/target/component_specified.go?s=664:735#L18)
``` go
func Scroll(componentId, distance string, args ...interface{}) *Command
```
实现 ScrollView 和 List 的窗口滑动，将 List 当前展现的窗口视图向上或者向下滑动指定的距离

@param componentId: 绑定的组件 id

@param distance: 每次窗口滑动的距离，如 20dp 或者 10%。负值如 -20dp 或者 -10%，表示向下滑动

@param args[0] duration: 滚动动画持续时间，单位毫秒，默认 500 毫秒 (optional)

@param args[1] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[2] dWhen: 指令的条件执行表达式 (optional)


### <a name="ScrollToElement">func</a> [ScrollToElement](/src/target/component_specified.go?s=2949:3045#L95)
``` go
func ScrollToElement(componentId, targetComponentId, align string, args ...interface{}) *Command
```
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


### <a name="ScrollToIndex">func</a> [ScrollToIndex](/src/target/component_specified.go?s=1760:1853#L56)
``` go
func ScrollToIndex(componentId string, index int, align string, args ...interface{}) *Command
```
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


### <a name="SendEvent">func</a> [SendEvent](/src/target/message_interaciton.go?s=1120:1209#L23)
``` go
func SendEvent(componentId string, arguments []interface{}, args ...interface{}) *Command
```
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


### <a name="SetPage">func</a> [SetPage](/src/target/component_specified.go?s=4642:4725#L155)
``` go
func SetPage(componentId, position string, value int, args ...interface{}) *Command
```
将 pager 组件的视图切换到指定目录页面，index 值表示 pager 的页面序号

@param componentId: 绑定的组件 id

@param position: 枚举, 取值如下：relative 相对位置，absolute 绝对位置

@param value: 翻页步长, 假如有 N 个页, 取值 0 到 N-1

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)


### <a name="SetProps">func</a> [SetProps](/src/target/data_updating.go?s=1838:1931#L64)
``` go
func SetProps(componentId string, props map[string]interface{}, args ...interface{}) *Command
```
修改组件的可使用 prop 功能属性

@param componentId: 绑定的组件id

@param props: 功能属性对象: {key: 功能属性名称, value: 功能属性值}

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)


### <a name="SetSpeechMonitor">func</a> [SetSpeechMonitor](/src/target/page_control.go?s=2518:2617#L66)
``` go
func SetSpeechMonitor(componentId string, speechFinishedPosition int, args ...interface{}) *Command
```
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


### <a name="SetState">func</a> [SetState](/src/target/data_updating.go?s=2777:2854#L95)
``` go
func SetState(componentId, state, value string, args ...interface{}) *Command
```
改变组件的部分可供修改的属性值，注：目前该属性仅支持通过事件绑定使用 (不再建议使用，推荐使用如上的 SetStyles、SetProps 指令来替代它)

@param componentId: 绑定的组件id

@param state: 属性名称

@param value: 对应的属性值, 可为空

@param args[0] stateType: 属性类型，枚举值: PROP 组件属性（默认）, STYLE 样式属性 (optional)

@param args[1] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[2] dWhen: 指令的条件执行表达式 (optional)


### <a name="SetStyles">func</a> [SetStyles](/src/target/data_updating.go?s=1146:1241#L37)
``` go
func SetStyles(componentId string, styles map[string]interface{}, args ...interface{}) *Command
```
修改/新增组件的可使用 style 样式属性

@param componentId: 绑定的组件id

@param styles: 样式属性对象: {key: 样式属性名称, value: 样式属性值}

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)


### <a name="UpdateComponent">func</a> [UpdateComponent](/src/target/data_updating.go?s=3698:3790#L127)
``` go
func UpdateComponent(componentId string, document interface{}, args ...interface{}) *Command
```
（该指令只能通过云端下发指令的方式使用）异步刷新指令, 使用指定的 document 描述的组件替换指定的 componentId 对应的组件,
document 的语法结构和 DPL_DOCUMENT 一致

@param componentId: 需要被替换内容的组件 id

@param document: 更新的组件内容 结构参考：DPL_DOCUMENT

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)


### <a name="UpdateDataSource">func</a> [UpdateDataSource](/src/target/data_updating.go?s=509:589#L14)
``` go
func UpdateDataSource(data map[string]interface{}, args ...interface{}) *Command
```
更新引用数据源的部分数据内容，会使 dpl 中通过动态引用方式使用该数据的对应内容同步更新

@param data: 与 dataSource 保持相同结构的需要更新的数据内容，该指令被执行时，data 内容会被 merge 更新到原 dataSource 中

@param args[0] delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)

@param args[1] dWhen: 指令的条件执行表达式 (optional)





### <a name="Command.GetCommandFromFile">func</a> (\*Command) [GetCommandFromFile](/src/target/command.go?s=2732:2795#L69)
``` go
func (cmd *Command) GetCommandFromFile(path ...string) *Command
```
Load Command from json file.




### <a name="Command.Marshal">func</a> (\*Command) [Marshal](/src/target/command.go?s=2618:2663#L64)
``` go
func (cmd *Command) Marshal() ([]byte, error)
```
Marshal Command to []byte




### <a name="Command.SetOnFinished">func</a> (\*Command) [SetOnFinished](/src/target/command.go?s=3064:3127#L84)
``` go
func (cmd *Command) SetOnFinished(onFinished *Command) *Command
```



### <a name="Command.SetOnInterrupted">func</a> (\*Command) [SetOnInterrupted](/src/target/command.go?s=3179:3248#L89)
``` go
func (cmd *Command) SetOnInterrupted(onInterrupted *Command) *Command
```



### <a name="Command.Unmarshal">func</a> (\*Command) [Unmarshal](/src/target/command.go?s=2496:2544#L59)
``` go
func (cmd *Command) Unmarshal(data []byte) error
```
Unmarshal Command from json []byte.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
