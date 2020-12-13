

# dpl
`import "github.com/yodstar/dueros-dpl/dpl"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type Document](#Document)
  * [func NewDocument(path ...string) *Document](#NewDocument)
  * [func (doc *Document) AddComponent(components ...interface{}) *Document](#Document.AddComponent)
  * [func (doc *Document) AddComponentFromFile(path ...string) *Document](#Document.AddComponentFromFile)
  * [func (doc *Document) AddParameter(parameters ...string) *Document](#Document.AddParameter)
  * [func (doc *Document) GetDocumentFromFile(path ...string) *Document](#Document.GetDocumentFromFile)
  * [func (doc *Document) GetLastError() error](#Document.GetLastError)
  * [func (doc *Document) GetRenderDocument() *RenderDocument](#Document.GetRenderDocument)
  * [func (doc *Document) Marshal() ([]byte, error)](#Document.Marshal)
  * [func (doc *Document) SetAbilities(abilities []interface{}) *Document](#Document.SetAbilities)
  * [func (doc *Document) SetDataSource(dataSource map[string]interface{}) *Document](#Document.SetDataSource)
  * [func (doc *Document) SetDuration(duration int) *Document](#Document.SetDuration)
  * [func (doc *Document) SetEvents(events map[string]interface{}) *Document](#Document.SetEvents)
  * [func (doc *Document) SetLayouts(layouts map[string]Template) *Document](#Document.SetLayouts)
  * [func (doc *Document) SetMainTemplate(parameters []string, components ...interface{}) *Document](#Document.SetMainTemplate)
  * [func (doc *Document) SetOnStackDuration(onStackDuration int) *Document](#Document.SetOnStackDuration)
  * [func (doc *Document) SetPageName(pageName string) *Document](#Document.SetPageName)
  * [func (doc *Document) SetRenderConfig(renderConfig interface{}) *Document](#Document.SetRenderConfig)
  * [func (doc *Document) SetResources(resources []interface{}) *Document](#Document.SetResources)
  * [func (doc *Document) SetStyles(styles map[string]interface{}) *Document](#Document.SetStyles)
  * [func (doc *Document) SetStylesheet(stylesheet interface{}) *Document](#Document.SetStylesheet)
  * [func (doc *Document) Unmarshal(data []byte) error](#Document.Unmarshal)
* [type RenderConfig](#RenderConfig)
  * [func NewRenderConfig(width, height string, globalVslDisabled bool) *RenderConfig](#NewRenderConfig)
* [type RenderDocument](#RenderDocument)
  * [func NewRenderDocument(doc *Document) *RenderDocument](#NewRenderDocument)
  * [func (r *RenderDocument) Marshal() ([]byte, error)](#RenderDocument.Marshal)
  * [func (r *RenderDocument) SetDocument(doc *Document)](#RenderDocument.SetDocument)
  * [func (r *RenderDocument) Unmarshal(data []byte) error](#RenderDocument.Unmarshal)
* [type Template](#Template)
  * [func NewTemplate(parameters []string, components ...interface{}) *Template](#NewTemplate)


#### <a name="pkg-files">Package files</a>
[bot.go](/src/github.com/yodstar/dueros-dpl/dpl/bot.go) [document.go](/src/github.com/yodstar/dueros-dpl/dpl/document.go) [execute_commands.go](/src/github.com/yodstar/dueros-dpl/dpl/execute_commands.go) [render_config.go](/src/github.com/yodstar/dueros-dpl/dpl/render_config.go) [render_document.go](/src/github.com/yodstar/dueros-dpl/dpl/render_document.go) [resource.go](/src/github.com/yodstar/dueros-dpl/dpl/resource.go) 






## <a name="Document">type</a> [Document](/src/target/document.go?s=925:3198#L22)
``` go
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

```
document: 包含渲染环境、指定渲染页面的 document 协议内容数据对象

DPL document 是一个基于 bot 协议的描述、有着完整结构定义的 JSON 对象数据，
它定义了协议中对应于在页面中需要渲染展现的动态渲染模板部分
（包括在页面上渲染展示的所有组件、数据内容和布局结构）以及设定页面模板的默认配置和交互能力
（无交互超时退出时间、预设事件绑定指令执行机制等）。


	该字段的解析都必须通过 bot 端下发的方式来执行，目前对于 DPL document 字段属性有两种应用场景：
	
	DPL.RenderDocument 指令下发必须包含该字段；
	
	DPL.ExecuteCommands 指令中包含 document 更新的指令类型必须包含该字段；







### <a name="NewDocument">func</a> [NewDocument](/src/target/document.go?s=8724:8766#L224)
``` go
func NewDocument(path ...string) *Document
```




### <a name="Document.AddComponent">func</a> (\*Document) [AddComponent](/src/target/document.go?s=7864:7934#L192)
``` go
func (doc *Document) AddComponent(components ...interface{}) *Document
```



### <a name="Document.AddComponentFromFile">func</a> (\*Document) [AddComponentFromFile](/src/target/document.go?s=8063:8130#L198)
``` go
func (doc *Document) AddComponentFromFile(path ...string) *Document
```
Add Component from json file.




### <a name="Document.AddParameter">func</a> (\*Document) [AddParameter](/src/target/document.go?s=7694:7759#L187)
``` go
func (doc *Document) AddParameter(parameters ...string) *Document
```



### <a name="Document.GetDocumentFromFile">func</a> (\*Document) [GetDocumentFromFile](/src/target/document.go?s=3483:3549#L51)
``` go
func (doc *Document) GetDocumentFromFile(path ...string) *Document
```
Load Document from json file.




### <a name="Document.GetLastError">func</a> (\*Document) [GetLastError](/src/target/document.go?s=8655:8696#L220)
``` go
func (doc *Document) GetLastError() error
```



### <a name="Document.GetRenderDocument">func</a> (\*Document) [GetRenderDocument](/src/target/document.go?s=8558:8614#L216)
``` go
func (doc *Document) GetRenderDocument() *RenderDocument
```
DPL 渲染指令结构(RenderDocument)




### <a name="Document.Marshal">func</a> (\*Document) [Marshal](/src/target/document.go?s=3367:3413#L46)
``` go
func (doc *Document) Marshal() ([]byte, error)
```
Marshal Document to []byte




### <a name="Document.SetAbilities">func</a> (\*Document) [SetAbilities](/src/target/document.go?s=7117:7185#L172)
``` go
func (doc *Document) SetAbilities(abilities []interface{}) *Document
```
abilities: 模板接入的能力，数组: ability[]

ability: 能力组件


	Gesture: 支持在端上打开手势识别能力，可以识别 手掌，Ok类型手势，触发对应的事件：
	
		onPalm: 识别手掌手势成功，触发事件
	
		onOk: 识别Ok手势成功，触发事件
	
	Form: 支持表单能力，可以声明多个表单项组件（Input）属于该表单，然后基于调用表单方法与事件统一、部分操控表单项
	
		onReady: 表单中所有表单项满足校验条件，进入可被提交状态
	
		onPrepare: 表单中存在部分表单项未满足校验条件，不可被提交
	
		onSubmit: 表单成功提交
	
		onValidateFail: 表单提交验证失败




### <a name="Document.SetDataSource">func</a> (\*Document) [SetDataSource](/src/target/document.go?s=5427:5506#L112)
``` go
func (doc *Document) SetDataSource(dataSource map[string]interface{}) *Document
```
dataSource: 用户自定义的数据部分, 可以在 COMPONENT 中引用 dataSource, 引用方式${}




### <a name="Document.SetDuration">func</a> (\*Document) [SetDuration](/src/target/document.go?s=3931:3987#L67)
``` go
func (doc *Document) SetDuration(duration int) *Document
```
duration: 用户指定模版无交互超时退出时间，单位毫秒ms, 取值需要大于0, 默认 30 秒




### <a name="Document.SetEvents">func</a> (\*Document) [SetEvents](/src/target/document.go?s=4682:4753#L91)
``` go
func (doc *Document) SetEvents(events map[string]interface{}) *Document
```
events: 基于模板的全局 events 描述




### <a name="Document.SetLayouts">func</a> (\*Document) [SetLayouts](/src/target/document.go?s=5659:5729#L118)
``` go
func (doc *Document) SetLayouts(layouts map[string]Template) *Document
```
layouts: 用户自定义样式, 可以在 mainTemplate 中的 items 中引用自定义的 layout




### <a name="Document.SetMainTemplate">func</a> (\*Document) [SetMainTemplate](/src/target/document.go?s=7518:7612#L182)
``` go
func (doc *Document) SetMainTemplate(parameters []string, components ...interface{}) *Document
```
mainTemplate: 基于嵌套组件组合的模板解析入口


	parameters: 作为指定 dataSource 数据源的引用别名
	
	items: 模版的嵌套子组件数组，基于嵌套的结构和组件的类型提供了整个模板页面的渲染展现和布局、功能




### <a name="Document.SetOnStackDuration">func</a> (\*Document) [SetOnStackDuration](/src/target/document.go?s=4134:4204#L73)
``` go
func (doc *Document) SetOnStackDuration(onStackDuration int) *Document
```
onStackDuration: 当模版被压栈时的超时退出时间，单位毫秒ms, 默认 10 分钟




### <a name="Document.SetPageName">func</a> (\*Document) [SetPageName](/src/target/document.go?s=4356:4415#L79)
``` go
func (doc *Document) SetPageName(pageName string) *Document
```
pageName: 全局定义的 pageName 参数，用来给 bot 端定义该页面的名称




### <a name="Document.SetRenderConfig">func</a> (\*Document) [SetRenderConfig](/src/target/document.go?s=4506:4578#L85)
``` go
func (doc *Document) SetRenderConfig(renderConfig interface{}) *Document
```
renderConfig: 页面渲染配置信息




### <a name="Document.SetResources">func</a> (\*Document) [SetResources](/src/target/document.go?s=4983:5051#L100)
``` go
func (doc *Document) SetResources(resources []interface{}) *Document
```
resources:
用户自定义 style 的属性常量, 可以在 styles 属性中引用属性常量, 引用方式@

- 对应的属性只能是 description、colors、dimensions




### <a name="Document.SetStyles">func</a> (\*Document) [SetStyles](/src/target/document.go?s=5206:5277#L106)
``` go
func (doc *Document) SetStyles(styles map[string]interface{}) *Document
```
styles: 用户自定义 styles 样式, 可以在 COMPONENT 中引入对应的自定义 styles 名称




### <a name="Document.SetStylesheet">func</a> (\*Document) [SetStylesheet](/src/target/document.go?s=6240:6308#L148)
``` go
func (doc *Document) SetStylesheet(stylesheet interface{}) *Document
```
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




### <a name="Document.Unmarshal">func</a> (\*Document) [Unmarshal](/src/target/document.go?s=3243:3292#L41)
``` go
func (doc *Document) Unmarshal(data []byte) error
```
Unmarshal Document from json []byte.




## <a name="RenderConfig">type</a> [RenderConfig](/src/target/render_config.go?s=44:247#L4)
``` go
type RenderConfig struct {
    Viewport          *Viewport `json:"viewport"`          // viewport: 页面可视区域大小
    GlobalVslDisabled bool      `json:"globalVslDisabled"` // globalVslDisabled:
}

```
页面渲染配置信息







### <a name="NewRenderConfig">func</a> [NewRenderConfig](/src/target/render_config.go?s=251:331#L9)
``` go
func NewRenderConfig(width, height string, globalVslDisabled bool) *RenderConfig
```




## <a name="RenderDocument">type</a> [RenderDocument](/src/target/render_document.go?s=620:1037#L16)
``` go
type RenderDocument struct {
    directive.BaseDirective           // type: 目前仅提供 DPL.RenderDocument 作为页面的渲染模板指令
    Token                   string    `json:"token"`    // token: 作为当前页面渲染指令在页面渲染后的唯一标识
    Document                *Document `json:"document"` // document: 包含渲染环境、指定渲染页面的 document 协议内容数据对象
}

```
DPL 渲染指令结构(RenderDocument)

当一个技能服务收到来自设备端的请求时，它可以以一个 DPL.RenderDocument 页面渲染指令作为响应，
在 RenderDocument 指令中包含了一个 document 字段内容作为页面渲染的模板描述，
并指定一个唯一的 token 字段作为页面的标识下发该指令到支持 DPL 渲染的设备端上来渲染页面模板。

用于渲染 DPL 模板指令，通过下发该指令在设备端上渲染一个 DPL 页面







### <a name="NewRenderDocument">func</a> [NewRenderDocument](/src/target/render_document.go?s=1383:1436#L36)
``` go
func NewRenderDocument(doc *Document) *RenderDocument
```




### <a name="RenderDocument.Marshal">func</a> (\*RenderDocument) [Marshal](/src/target/render_document.go?s=1220:1270#L28)
``` go
func (r *RenderDocument) Marshal() ([]byte, error)
```
Marshal RenderDocument to []byte




### <a name="RenderDocument.SetDocument">func</a> (\*RenderDocument) [SetDocument](/src/target/render_document.go?s=1304:1355#L32)
``` go
func (r *RenderDocument) SetDocument(doc *Document)
```



### <a name="RenderDocument.Unmarshal">func</a> (\*RenderDocument) [Unmarshal](/src/target/render_document.go?s=1088:1141#L23)
``` go
func (r *RenderDocument) Unmarshal(data []byte) error
```
Unmarshal RenderDocument from json []byte.




## <a name="Template">type</a> [Template](/src/target/document.go?s=9236:9439#L243)
``` go
type Template struct {
    Parameters []string      `json:"parameters"` // parameters: 指定 dataSource 的引用别名
    Items      []interface{} `json:"items"`      // items: 模版的子组件数组
}

```
基于嵌套组件组合的模板解析入口

parameters
作为指定 dataSource 数据源的引用别名

items
模版的嵌套子组件数组，基于嵌套的结构和组件的类型提供了整个模板页面的渲染展现和布局、功能







### <a name="NewTemplate">func</a> [NewTemplate](/src/target/document.go?s=9443:9517#L248)
``` go
func NewTemplate(parameters []string, components ...interface{}) *Template
```








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
