package dpl

import (
	"encoding/json"
	"io/ioutil"

	"github.com/yodstar/dueros-dpl/dpl/util"
)

// document: 包含渲染环境、指定渲染页面的 document 协议内容数据对象
//
// DPL document 是一个基于 bot 协议的描述、有着完整结构定义的 JSON 对象数据，
// 它定义了协议中对应于在页面中需要渲染展现的动态渲染模板部分
// （包括在页面上渲染展示的所有组件、数据内容和布局结构）以及设定页面模板的默认配置和交互能力
// （无交互超时退出时间、预设事件绑定指令执行机制等）。
//
// 	该字段的解析都必须通过 bot 端下发的方式来执行，目前对于 DPL document 字段属性有两种应用场景：
//
// 	DPL.RenderDocument 指令下发必须包含该字段；
//
// 	DPL.ExecuteCommands 指令中包含 document 更新的指令类型必须包含该字段；
type Document struct {
	Error           error                  `json:"-"`
	Type            string                 `json:"type"`                      // type: 指令类型，即"DPL"
	Version         string                 `json:"version"`                   // version: 定版本信息，目前推荐使用 2.0 版本
	Duration        int                    `json:"duration,omitempty"`        // duration: 用户指定模版无交互超时退出时间，单位毫秒ms, 取值需要大于0, 默认 30 秒
	OnStackDuration int                    `json:"onStackDuration,omitempty"` // onStackDuration: 当模版被压栈时的超时退出时间，单位毫秒ms, 默认 10 分钟
	PageName        string                 `json:"pageName,omitempty"`        // pageName: 全局定义的 pageName 参数，用来给 bot 端定义该页面的名称
	RenderConfig    interface{}            `json:"renderConfig,omitempty"`    // renderConfig: 页面渲染配置信息
	Events          map[string]interface{} `json:"events,omitempty"`          // events: 基于模板的全局 events 描述
	Resources       []interface{}          `json:"resources,omitempty"`       // resources: 用户自定义 style 的属性常量, 可以在 styles 属性中引用属性常量, 引用方式@
	Styles          map[string]interface{} `json:"styles,omitempty"`          // styles: 用户自定义 styles 样式, 可以在 COMPONENT 中引入对应的自定义 styles 名称
	DataSource      map[string]interface{} `json:"dataSource,omitempty"`      // dataSource: 用户自定义的数据部分, 可以在 COMPONENT 中引用 dataSource, 引用方式${}
	Layouts         map[string]Template    `json:"layouts,omitempty"`         // layouts: 用户自定义样式, 可以在 mainTemplate 中的 items 中引用自定义的 layout
	Stylesheet      interface{}            `json:"stylesheet,omitempty"`      // stylesheet: 通过 class 与 mediaQuery 的组合使用，更好的支持基于不同尺寸的设备适配
	Abilities       []interface{}          `json:"abilities,omitempty"`       // abilities: 模板接入的能力，数组: ability[]
	MainTemplate    *Template              `json:"mainTemplate"`              // mainTemplate: 基于嵌套组件组合的模板解析入口
}

// Unmarshal Document from json []byte.
func (doc *Document) Unmarshal(data []byte) error {
	return json.Unmarshal(data, doc)
}

// Marshal Document to []byte
func (doc *Document) Marshal() ([]byte, error) {
	return json.Marshal(doc)
}

// Load Document from json file.
func (doc *Document) GetDocumentFromFile(path ...string) *Document {
	for i, _ := range path {
		data, err := ioutil.ReadFile(path[i])
		if err != nil {
			doc.Error = util.Error(err)
			return doc
		}
		if err = json.Unmarshal(data, doc); err != nil {
			doc.Error = util.Error(err)
			return doc
		}
	}
	return doc
}

// duration: 用户指定模版无交互超时退出时间，单位毫秒ms, 取值需要大于0, 默认 30 秒
func (doc *Document) SetDuration(duration int) *Document {
	doc.Duration = duration
	return doc
}

// onStackDuration: 当模版被压栈时的超时退出时间，单位毫秒ms, 默认 10 分钟
func (doc *Document) SetOnStackDuration(onStackDuration int) *Document {
	doc.OnStackDuration = onStackDuration
	return doc
}

// pageName: 全局定义的 pageName 参数，用来给 bot 端定义该页面的名称
func (doc *Document) SetPageName(pageName string) *Document {
	doc.PageName = pageName
	return doc
}

// renderConfig: 页面渲染配置信息
func (doc *Document) SetRenderConfig(renderConfig interface{}) *Document {
	doc.RenderConfig = renderConfig
	return doc
}

// events: 基于模板的全局 events 描述
func (doc *Document) SetEvents(events map[string]interface{}) *Document {
	doc.Events = events
	return doc
}

// resources:
// 用户自定义 style 的属性常量, 可以在 styles 属性中引用属性常量, 引用方式@
//
//- 对应的属性只能是 description、colors、dimensions
func (doc *Document) SetResources(resources []interface{}) *Document {
	doc.Resources = resources
	return doc
}

// styles: 用户自定义 styles 样式, 可以在 COMPONENT 中引入对应的自定义 styles 名称
func (doc *Document) SetStyles(styles map[string]interface{}) *Document {
	doc.Styles = styles
	return doc
}

// dataSource: 用户自定义的数据部分, 可以在 COMPONENT 中引用 dataSource, 引用方式${}
func (doc *Document) SetDataSource(dataSource map[string]interface{}) *Document {
	doc.DataSource = dataSource
	return doc
}

// layouts: 用户自定义样式, 可以在 mainTemplate 中的 items 中引用自定义的 layout
func (doc *Document) SetLayouts(layouts map[string]Template) *Document {
	doc.Layouts = layouts
	return doc
}

/*
stylesheet: 通过 class 与 mediaQuery 的组合使用，更好的支持基于不同尺寸的设备适配

	[配置形式一]

	'classNameString'
	样式聚合类名

	{'styleNameString': 'styleProp'}
	样式属性对象

	[配置形式二]

	mediaQuery
	设备类型（X5）或设备条件判断属性

	classList
	样式聚合类列表

	'classNameString'
	样式聚合类名

	{'styleNameString': 'styleProp'}
	样式属性对象
*/
func (doc *Document) SetStylesheet(stylesheet interface{}) *Document {
	doc.Stylesheet = stylesheet
	return doc
}

// abilities: 模板接入的能力，数组: ability[]
//
// ability: 能力组件
//
// 	Gesture: 支持在端上打开手势识别能力，可以识别 手掌，Ok类型手势，触发对应的事件：
//
// 		onPalm: 识别手掌手势成功，触发事件
//
// 		onOk: 识别Ok手势成功，触发事件
//
//	Form: 支持表单能力，可以声明多个表单项组件（Input）属于该表单，然后基于调用表单方法与事件统一、部分操控表单项
//
// 		onReady: 表单中所有表单项满足校验条件，进入可被提交状态
//
// 		onPrepare: 表单中存在部分表单项未满足校验条件，不可被提交
//
// 		onSubmit: 表单成功提交
//
//		onValidateFail: 表单提交验证失败
func (doc *Document) SetAbilities(abilities []interface{}) *Document {
	doc.Abilities = abilities
	return doc
}

// mainTemplate: 基于嵌套组件组合的模板解析入口
//
// 	parameters: 作为指定 dataSource 数据源的引用别名
//
// 	items: 模版的嵌套子组件数组，基于嵌套的结构和组件的类型提供了整个模板页面的渲染展现和布局、功能
func (doc *Document) SetMainTemplate(parameters []string, components ...interface{}) *Document {
	doc.MainTemplate = NewTemplate(parameters, components...)
	return doc
}

func (doc *Document) AddParameter(parameters ...string) *Document {
	doc.MainTemplate.Parameters = append(doc.MainTemplate.Parameters, parameters...)
	return doc
}

func (doc *Document) AddComponent(components ...interface{}) *Document {
	doc.MainTemplate.Items = append(doc.MainTemplate.Items, components...)
	return doc
}

// Add Component from json file.
func (doc *Document) AddComponentFromFile(path ...string) *Document {
	var component map[string]interface{}
	for i, _ := range path {
		data, err := ioutil.ReadFile(path[i])
		if err != nil {
			doc.Error = util.Error(err)
			return doc
		}
		if err = json.Unmarshal(data, &component); err != nil {
			doc.Error = util.Error(err)
			return doc
		}
		doc.MainTemplate.Items = append(doc.MainTemplate.Items, component)
	}
	return doc
}

// DPL 渲染指令结构(RenderDocument)
func (doc *Document) GetRenderDocument() *RenderDocument {
	return NewRenderDocument(doc)
}

func (doc *Document) GetLastError() error {
	return doc.Error
}

func NewDocument(path ...string) *Document {
	doc := &Document{
		Type:         "DPL",
		Version:      "2.0",
		MainTemplate: NewTemplate([]string{}),
	}
	if len(path) > 0 {
		doc.GetDocumentFromFile(path...)
	}
	return doc
}

// 基于嵌套组件组合的模板解析入口
//
// parameters
// 作为指定 dataSource 数据源的引用别名
//
// items
// 模版的嵌套子组件数组，基于嵌套的结构和组件的类型提供了整个模板页面的渲染展现和布局、功能
type Template struct {
	Parameters []string      `json:"parameters"` // parameters: 指定 dataSource 的引用别名
	Items      []interface{} `json:"items"`      // items: 模版的子组件数组
}

func NewTemplate(parameters []string, components ...interface{}) *Template {
	return &Template{
		Parameters: parameters,
		Items:      components,
	}
}
