package style

import (
	"encoding/json"
	"io/ioutil"
	"unsafe"

	"github.com/yodstar/dueros-dpl/dpl/util"
)

// 通过class与mediaquery的组合使用，可以更好的支持基于不同尺寸的适配问题
//
// stylesheet字段可以是一个单对象或者一个数组，其中每个对象以 key 表示 className 信息，
// 可以使用,来分割多个 className。基于实际需求的不同，可以是简单的通过stylesheet，设置
// 模板所需要的class类，或者基于不同屏幕尺寸的设备适配需求，配合 mediaquery，来基于设备
// 宽高区分配置不同的样式属性。
type Stylesheet []interface{}

type xStylesheet struct {
	Stylesheet
	Error error
}

// Unmarshal Stylesheet from json []byte.
func (s *Stylesheet) Unmarshal(data []byte) error {
	return json.Unmarshal(data, s)
}

// Marshal Stylesheet to []byte
func (s *Stylesheet) Marshal() ([]byte, error) {
	return json.Marshal(s)
}

// Load Stylesheet from json file.
func (s *Stylesheet) GetStylesheetFromFile(path ...string) *Stylesheet {
	x := (*xStylesheet)(unsafe.Pointer(s))
	for i, _ := range path {
		data, err := ioutil.ReadFile(path[i])
		if err != nil {
			x.Error = util.Error(err)
			return s
		}
		if err = json.Unmarshal(data, s); err != nil {
			x.Error = util.Error(err)
			return s
		}
	}
	return s
}

func (s *Stylesheet) Append(args ...interface{}) *Stylesheet {
	*s = append(*s, args...)
	return s
}

func (s *Stylesheet) GetLastError() error {
	x := (*xStylesheet)(unsafe.Pointer(s))
	return x.Error
}

func NewStylesheet(path ...string) *Stylesheet {
	s := (*Stylesheet)(unsafe.Pointer(&xStylesheet{}))
	if len(path) > 0 {
		s.GetStylesheetFromFile(path...)
	}
	return s
}

// 通过将组件中频繁使用到的一些样式属性聚合起来形成 class 类，存放到 stylesheet 中，可供多个组件
// 去申明使用该class，然后端上渲染对应组件时会将自己引用的 class 在 stylesheet 中检索、提取，
//以达到减少云端下发数据量和便于修改、复用的目的
type ClassList map[string]*Styles

func (c *ClassList) SetClass(name string, class *Styles) *ClassList {
	(*c)[name] = class
	return c
}

// 设置不同类型的设备屏幕分辨率条件，并根据对应的条件，给相应符合条件的设备端调用相对应的样式表，
// 通过宽度和高度的表达式限制，允许内容的呈现针对一个特定范围的输出设备而进行裁剪，而不必改变内容本身。
type ClassMediaQuery struct {
	MediaQuery string     `json:"mediaQuery"`
	ClassList  *ClassList `json:"classList"`
}

//
// 在需要基于不同设备来区分使用不同的样式属性的时候，通过在 stylesheet 中设置 mediaQuery，同时
// 把原本应该的class内容放到classList字段中。
//
// 	1. mediaQuery: "(min-width: 500dp)":
// 	标明该对象下对应 classList 的样式属性可以生效前提是设备的最小宽度必须大于等于 500dp;
//
// 	2. mediaQuery: "(min-width: 500dp) and (max-width: 1000dp)":
// 	标明该对象下对应 classList 的样式属性可以生效得到前提是设备的宽度范围必须在 500dp 到 1000dp之间;
//
// 	3. mediaQuery:  "(max-width: 500dp) , (min-hight: 400dp) ":
// 	标明该对象下对应 classList 的样式属性可以生效得到前提是设备的宽高必须在最大不超过500dp 高度且高度不小于 100dp;
//
// 	4. mediaQuery:  "(500dp <= width <= 1000dp) ":
// 	标明该对象下对应 classList 的样式属性可以生效得到前提是设备的宽度范围必须在 500dp 到 1000dp之间（与2 等价）;
//
// 	5. mediaQuery:  ”x5":
// 	目前专门为5寸设备（2：1宽高比）设定的名称描述方式，等价于 ”(max-height: 480dp)“
func (q *ClassMediaQuery) SetMediaQuery(mediaQuery string) *ClassMediaQuery {
	q.MediaQuery = mediaQuery
	return q
}

func (q *ClassMediaQuery) SetClassList(classList *ClassList) *ClassMediaQuery {
	q.ClassList = classList
	return q
}

func (q *ClassMediaQuery) SetClass(name string, class *Styles) *ClassMediaQuery {
	q.ClassList.SetClass(name, class)
	return q
}

func NewClassMediaQuery(mediaQuery string) *ClassMediaQuery {
	return &ClassMediaQuery{
		MediaQuery: mediaQuery,
		ClassList:  &ClassList{},
	}
}
