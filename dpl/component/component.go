package component

import (
	"encoding/json"
	"io/ioutil"

	"github.com/yodstar/dueros-dpl/dpl/command"
	"github.com/yodstar/dueros-dpl/dpl/style"
	"github.com/yodstar/dueros-dpl/dpl/util"
)

type Component struct {
	Error       error                         `json:"-"`
	Type        string                        `json:"type"`
	ComponentId string                        `json:"componentId"`
	Class       string                        `json:"class,omitempty"`
	Styles      *style.Styles                 `json:"styles,omitempty"`
	Props       map[string]interface{}        `json:"props,omitempty"`
	Events      map[string][]*command.Command `json:"events,omitempty"`
	Items       []*Component                  `json:"items,omitempty"`
	DShow       string                        `json:"dShow,omitempty"`
	DWhen       string                        `json:"dWhen,omitempty"`
	DFor        string                        `json:"dFor,omitempty"`
}

// Unmarshal Component from json []byte.
func (com *Component) Unmarshal(data []byte) error {
	return json.Unmarshal(data, com)
}

// Marshal Component to []byte
func (com *Component) Marshal() ([]byte, error) {
	return json.Marshal(com)
}

// Load Component from json file.
func (com *Component) GetCommandFromFile(path ...string) *Component {
	for i, _ := range path {
		data, err := ioutil.ReadFile(path[i])
		if err != nil {
			com.Error = util.Error(err)
			return com
		}
		if err = json.Unmarshal(data, com); err != nil {
			com.Error = util.Error(err)
			return com
		}
	}
	return com
}

func (com *Component) SetStyles(styles *style.Styles) *Component {
	com.Styles = styles
	return com
}

func (com *Component) SetStylesJSON(styles string) *Component {
	if com.Styles == nil {
		com.Styles = style.NewStyles()
	}
	if err := json.Unmarshal([]byte(styles), com.Styles); err != nil {
		com.Error = util.Error(err)
	}
	return com
}

func (com *Component) SetStylesCSS(styles string) *Component {
	if com.Styles == nil {
		com.Styles = style.NewStyles()
	}
	com.Styles.LoadCSSProps(styles)
	return com
}

func (com *Component) SetProps(name string, value interface{}) *Component {
	com.Props[name] = value
	return com
}

func (com *Component) SetPropsJSON(props string) *Component {
	if err := json.Unmarshal([]byte(props), &com.Props); err != nil {
		com.Error = util.Error(err)
	}
	return com
}

func (com *Component) SetEvents(name string, events ...*command.Command) *Component {
	com.Events[name] = events
	return com
}

func (com *Component) SetEventsJSON(events string) *Component {
	if err := json.Unmarshal([]byte(events), &com.Events); err != nil {
		com.Error = util.Error(err)
	}
	return com
}

func (com *Component) BindEvents(name string, events ...*command.Command) *Component {
	com.Events[name] = append(com.Events[name], events...)
	return com
}

func (com *Component) AppendComponent(items ...*Component) {
	com.Items = append(com.Items, items...)
}

func (com *Component) GetLastError() error {
	return com.Error
}

/*
@param type: 组件类型

@param args[0] componentId: 组件id

@param args[1:] props: key(string), value(interface{})
*/
func NewComponent(typ string, args ...interface{}) *Component {
	com := &Component{Type: typ}
	com.Props = make(map[string]interface{})
	com.Events = make(map[string][]*command.Command)
	argc := len(args)
	if argc > 0 {
		com.ComponentId = args[0].(string)
		for i := 1; i < argc; i += 2 {
			if i+1 < argc {
				com.Props[args[i].(string)] = args[i+1]
			} else {
				com.Props[args[i].(string)] = nil
			}
		}
	}
	return com
}

/*
音频组件，目前以非可视化的组件提供方式，提供了接入音频的能力（多用于背景音）

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

#props

 	名称       类型        默认值     可选值                    必填      描述
 	src        string      -          -                         是        音频地址
 	looping    boolean     false      -                         否        是否循环, 开启以后音频循环播放
 	mute       boolean     false      -                         否        默认是否静音
 	autoplay   boolean     false      -                         否        是否开启自动播放, 用作背景音时一般为 true
 	type       string      music      music,ring,alarm,system   否        音频播放类型，不同的类型，有不同的优先级

#events

 	名称        参数       参数说明    event说明
 	onPlay       -         -           音频切换到播放状态时触发该事件
 	onPause      -         -           音频切换到暂停状态时触发该事件
 	onPrepared   -         -           音频状态初始化为可播放后可触发该事件
 	onEnd        -         -           音频播放结束后可触发该事件
 	onLoaded     -         -           组件加载完成后可触发该事件（仅在初次渲染完成时触发一次）
*/
func Audio(src string, args ...interface{}) *Component {
	com := NewComponent("Audio", args...)
	com.Props["src"] = src
	return com
}

/*
Container 作为基础的容器组件，通过各类样式属性的应用，可以设置宽高、内外边距以及背景颜色等，
可以嵌套自身以及其他组件，常用来实现页面布局以及嵌套其他组件实现局部样式布局

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

#props

 	名称            类型        默认值     可选值           必填    描述
 	display         string      show       show,hide        否      是否展现该容器
 	clickTimeout    number      1000       -                否      容器可被点击的时间间隔，单位毫秒
 	clickable       string      enable     enable,disable   否      默认容器的可点击态是基于是否有绑定 onClick 事件，增加这个属性，强制禁用可点击（即使有绑定 onClick）
 	enableVoice     boolean     false      -                否      是否支持注册语音话术，具体描述详见
 	voiceAction     string      -          -                否      语音话术类型，具体格式详见
 	voiceConfig     string      -          -                否      语音话术内容，具体格式详见

#events

 	名称           参数     参数说明      event说明
 	onClick        -        -             容器区域被点击后触发该事件
 	onLoaded       -        -             组件加载完成后可触发该事件（仅在初次渲染完成时触发一次）

#items

 	名称            说明
 	firstItem       容器头部插槽，最先展现的子 component
 	lastItem        容器尾部插槽，最后展现的子 component
*/
func Container(args ...interface{}) *Component {
	return NewComponent("Container", args...)
}

/*
通过各类样式属性的应用，提供了多个可组合使用值（day、hour、minute、second）的计时类型组件，可提供灵活组合使用，
作为计时类的展示组件，用于实现计数器展示效果以及相应的计数完成时执行逻辑

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

#组件特殊说明

 	Counter 组件中支持 嵌套任意多的 Container、Image、Text、Frame 组件，当且仅当其内部子组件存在包含的 Text 组件中的
 	对应 props 中 text 属性值为：

 	(1) #Counter.$day" 展示 Counter 组件的日期天数
 	(2) #Counter.$hour" 展示 Counter 组件的日期小时
 	(3) #Counter.$minute" 展示 Counter 组件的日期分钟
 	(4) #Counter.$second" 展示 Counter 组件的日期秒数

#props

 	名称         类型            默认值    可选值                    必填     描述
 	timestamp    number,string  0         -                         否       计时时间参数，单位毫秒
 	tick         number,string  1000      -                         否       计时时间间隔，单位毫秒
 	delay        number         0         -                         否       组件计时的延时执行时间，单位毫秒
 	countMode    string         DECR      DECR,INCR                 否       计时方式，DECR 减量倒计时，INCR 增量顺计时
 	countType    string         SECOND    SECOND,MINUTE,HOUR,DAY    否       计时类型，支持以 SECOND 秒、MINUTE 分钟、HOUR 时、DAY 天 多种方式展现计时
 	placeholder  string         --        -                         否       计时占位符，初始化时展现
 	autoplay     boolean        true      -                         否       是否开始自动执行计时功能

#events

 	名称         参数      参数说明      event说明
 	onEnd        -         -             计时完成后可触发该事件，标识为计时完成状态
 	onLoaded     -         -             组件加载完成后可触发该事件（仅在初次渲染完成时触发一次）

#methods

 	方法名       参数       参数说明     method说明
 	reset        -          -            重置当前计时器，重新开始 delay 后执行计时
 	pause        -          -            暂停当前计时器
 	play         -          -            执行当前计时器计数
*/
func Counter(args ...interface{}) *Component {
	return NewComponent("Counter", args...)
}

/*
通过各类样式属性的应用，使用多张图片实现逐帧动画效果, 用户可以自定义帧动画的执行间隔、执行次数以及设置结束的触发事件

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

*/
func FrameAnimation(args ...interface{}) *Component {
	return NewComponent("FrameAnimation", args...)
}

/*
图片组件，用来在设备端上通过各类样式属性的应用呈现图片, 支持使用 png、jpg、gif、svg、apng（设备37版本及以上）、
webp（设备37版本及以上）格式

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

*/
func Image(args ...interface{}) *Component {
	return NewComponent("Image", args...)
}

/*
用于提供在设备端上输入能力的组件

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

*/
func Input(args ...interface{}) *Component {
	return NewComponent("Input", args...)
}

/*
用于定义一个列表, 可以设置列表的宽高位置, 指定列表项的展现方向, 通过绑定的列表数据渲染列表项。
还可以通过firstItem和lastItem来指定列表的头部和尾部样式

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

*/
func List(args ...interface{}) *Component {
	return NewComponent("List", args...)
}

/*
用定义滑动页面组件, 实现一个可以左右或者上下滑动的滑动页面组件， 形式上类似于手机 app 中的引导页,
可以设置宽高。 区别 List 组件的页面项和绑定数据项结构都是一样的, Pager 定义多少个页面项就会展示
多少页面, 各个页面项可以不一样, 绑定的数据源结构也可以不一样。 通过 firstItem 和 lastItem，Pager
也可以自定义首页和尾页样式

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

*/
func Pager(args ...interface{}) *Component {
	return NewComponent("Pager", args...)
}

/*
用于定义一个带滚动条滑动窗口, ScrollView 是一个宽高可伸缩的特殊的 Container。 当包含的内容宽高超出
设置的指定值, 就会产生滚动条, 基于指定的滑动方向滑动可以显示隐藏内容。ScrollView 多用于篇幅较长的
内容展示

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

*/
func ScrollView(args ...interface{}) *Component {
	return NewComponent("ScrollView", args...)
}

/*
展示基于部分样式属性的单行或多行文本内容

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

*/
func Text(args ...interface{}) *Component {
	return NewComponent("Text", args...)
}

/*
定时器组件，用于实现定时执行逻辑, 不展示UI样式

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

*/
func Timer(args ...interface{}) *Component {
	return NewComponent("Timer", args...)
}

/*
用于在设备端上展现浏览器网页内容的组件

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

*/
func WebView(args ...interface{}) *Component {
	return NewComponent("WebView", args...)
}

/*
用于定义视频组件, 可以指定播放界面的宽高, 通过相关属性设置播放组件的显示样式, 如上一个、下一个、
暂停、弹幕、进度条等, 还可以通过绑定onPlay、 onPause等事件实现数据状态上报以及和云端的交互，
如通过onPause上报当前播放状态，或者通过onNext上报云端，云端再返回下一个播放的url

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

*/
func Video(args ...interface{}) *Component {
	return NewComponent("Video", args...)
}

/*
可伴随页面下发时的 outspeech 中的内容屏幕同步滚动播报多行文本

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

*/
func Section(args ...interface{}) *Component {
	return NewComponent("Section", args...)
}

/*
通常用于定义头部 title 栏展示和增加通用的返回/页面点击关闭能力， 如果需要在页面中显示头部，
推荐将Header放在最外层容器的顶部，并根据合理的布局让它在合适的位置展示

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

*/
func Header(args ...interface{}) *Component {
	return NewComponent("Header", args...)
}

/*
通常用于定义底部 hint 栏展示和增加通用的点击跳转能力, 如果需要在页面中显示底部Hint提示，
建议将 Footer 放在最外层容器的底部，并根据合理的布局让它在合适的位置展示

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

*/
func Footer(args ...interface{}) *Component {
	return NewComponent("Footer", args...)
}

/*
针对拉取逻辑封装的弹幕组件

 	@param args[0] componentId: 组件id
 	@param args[1:] props: key(string), value(interface{})

*/
func Barrage(src string, args ...interface{}) *Component {
	com := NewComponent("Barrage", args...)
	com.Props["src"] = src
	return com
}
