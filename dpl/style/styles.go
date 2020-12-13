package style

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"unsafe"

	"github.com/yodstar/dueros-dpl/dpl/util"
)

// 样式
type Styles map[string]interface{}

type xStyles struct {
	Styles
	Error error
}

// Unmarshal Styles from json []byte.
func (s *Styles) Unmarshal(data []byte) error {
	return json.Unmarshal(data, s)
}

// Marshal Styles to []byte
func (s *Styles) Marshal() ([]byte, error) {
	return json.Marshal(s)
}

// Load Styles from json file.
func (s *Styles) GetStylesFromFile(path ...string) *Styles {
	x := (*xStyles)(unsafe.Pointer(s))
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

func (s *Styles) LoadCSSProps(cssProps string) *Styles {
	props := strings.Split(cssProps, ";")
	for i, _ := range props {
		style := strings.SplitN(props[i], ":", 2)
		if len(style) != 2 {
			continue
		}
		prop := strings.TrimSpace(style[0])
		value := strings.TrimSpace(style[1])
		switch value {
		case "true":
			(*s)[prop] = true
		case "false":
			(*s)[prop] = false
		default:
			(*s)[prop] = value
		}
	}
	return s
}

func (s *Styles) GetLastError() error {
	x := (*xStyles)(unsafe.Pointer(s))
	return x.Error
}

func NewStyles(path ...string) *Styles {
	s := (*Styles)(unsafe.Pointer(&xStyles{Styles: make(Styles)}))
	if len(path) > 0 {
		s.GetStylesFromFile(path...)
	}
	return s
}

//------------------------[common]-----------------------------

func (s *Styles) Viewport(width, height string) *Styles {
	(*s)["width"] = width
	(*s)["height"] = height
	return s
}

func (s *Styles) Width(value string) *Styles {
	(*s)["width"] = value
	return s
}

func (s *Styles) MinWidth(value string) *Styles {
	(*s)["min-width"] = value
	return s
}

func (s *Styles) MaxWidth(value string) *Styles {
	(*s)["max-width"] = value
	return s
}

func (s *Styles) Height(value string) *Styles {
	(*s)["height"] = value
	return s
}

func (s *Styles) MinHeight(value string) *Styles {
	(*s)["min-height"] = value
	return s
}

func (s *Styles) MaxHeight(value string) *Styles {
	(*s)["max-height"] = value
	return s
}

// opacity：取值范围为 [0, 1] 区间。默认值是 1，即完全不透明；0 是完全透明；0.5 是 50% 的透明度。
func (s *Styles) Opacity(value int) *Styles {
	(*s)["opacity"] = value
	return s
}

// background-color：设定元素的背景色，可选值为色值，支持RGB（ rgb(255, 0, 0) ）；RGBA（ rgba(255, 0, 0, 0.5) ）；
// 十六进制（ #ff0000 ）；色值关键字（red），默认值以组件的系统默认背景色为准
func (s *Styles) BackgroundColor(value string) *Styles {
	(*s)["background-color"] = value
	return s
}

// background: 设定元素的背景，支持http、https协议的网络图片，支持设定背景色（与background-color属性一致），支持渐变v1.35+。
func (s *Styles) Background(value string) *Styles {
	(*s)["background"] = value
	return s
}

// box-shadow (v1.31.0+): 设定元素的阴影
//
// 格式：inset(可选),offset-x,offset-y, blur-radius,color
// 	示例：
// 	inset 10dp 10dp 20dp red
// 	5dp 10dp 5dp rbga(0,0,0,0,5)
func (s *Styles) BoxShadow(value string) *Styles {
	(*s)["box-shadow"] = value
	return s
}

//------------------------[padding]-----------------------------

// 内边距，内容和边框之间的距离。值类型为 length，默认值 0。支持组合写法，接收1~4个参数。v1.31+支持百分比。
func (s *Styles) Padding(value string) *Styles {
	(*s)["padding"] = value
	return s
}

func (s *Styles) PaddingLeft(value string) *Styles {
	(*s)["padding-left"] = value
	return s
}

func (s *Styles) PaddingRight(value string) *Styles {
	(*s)["padding-right"] = value
	return s
}

func (s *Styles) PaddingTop(value string) *Styles {
	(*s)["padding-top"] = value
	return s
}

func (s *Styles) PaddingBottom(value string) *Styles {
	(*s)["padding-bottom"] = value
	return s
}

func (s *Styles) PaddingVertical(value string) *Styles {
	(*s)["padding-vertical"] = value
	return s
}

func (s *Styles) PaddingHorizontal(value string) *Styles {
	(*s)["padding-horizontal"] = value
	return s
}

//------------------------[margin]-----------------------------

// 外边距，元素和元素之间的空白距离。值类型为 length，默认值 0。支持组合写法，接收1~4个参数。v1.31+支持百分比。
func (s *Styles) Margin(value string) *Styles {
	(*s)["margin"] = value
	return s
}

func (s *Styles) MarginLeft(value string) *Styles {
	(*s)["margin-left"] = value
	return s
}

func (s *Styles) MarginRight(value string) *Styles {
	(*s)["margin-right"] = value
	return s
}

func (s *Styles) MarginTop(value string) *Styles {
	(*s)["margin-top"] = value
	return s
}

func (s *Styles) MarginBottom(value string) *Styles {
	(*s)["margin-bottom"] = value
	return s
}

func (s *Styles) MarginVertical(value string) *Styles {
	(*s)["margin-vertical"] = value
	return s
}

func (s *Styles) MarginHorizontal(value string) *Styles {
	(*s)["margin-horizontal"] = value
	return s
}

//------------------------[border]-----------------------------

// 边框,目前不支持组合写法，通过border-width、border-color、border-radius分别设置。
func (s *Styles) Border(width, color, radius string) *Styles {
	(*s)["border-width"] = width
	(*s)["border-color"] = color
	(*s)["border-radius"] = radius
	return s
}

func (s *Styles) BorderWidth(value string) *Styles {
	(*s)["border-width"] = value
	return s
}

func (s *Styles) BorderLeftWidth(value string) *Styles {
	(*s)["border-left-width"] = value
	return s
}

func (s *Styles) BorderTopWidth(value string) *Styles {
	(*s)["border-top-width"] = value
	return s
}

func (s *Styles) BorderRightWidth(value string) *Styles {
	(*s)["border-right-width"] = value
	return s
}

func (s *Styles) BorderBottomWidth(value string) *Styles {
	(*s)["border-bottom-width"] = value
	return s
}

func (s *Styles) BorderColor(value string) *Styles {
	(*s)["border-color"] = value
	return s
}

func (s *Styles) BorderLeftColor(value string) *Styles {
	(*s)["border-left-color"] = value
	return s
}

func (s *Styles) BorderTopColor(value string) *Styles {
	(*s)["border-top-color"] = value
	return s
}

func (s *Styles) BorderRightColor(value string) *Styles {
	(*s)["border-right-color"] = value
	return s
}

func (s *Styles) BorderBottomColor(value string) *Styles {
	(*s)["border-bottom-color"] = value
	return s
}

func (s *Styles) BorderRadius(value string) *Styles {
	(*s)["border-radius"] = value
	return s
}

func (s *Styles) BorderBottomLeftRadius(value string) *Styles {
	(*s)["border-bottom-left-radius"] = value
	return s
}

func (s *Styles) BorderBottomRightRadius(value string) *Styles {
	(*s)["border-bottom-right-radius"] = value
	return s
}

func (s *Styles) BorderTopLeftRadius(value string) *Styles {
	(*s)["border-top-left-radius"] = value
	return s
}

func (s *Styles) BorderTopRightRadius(value string) *Styles {
	(*s)["border-top-right-radius"] = value
	return s
}

//------------------------[flexbox]-----------------------------

// 设备端样式属性支持布局模型基于 CSS Flexbox，以便所有页面元素的排版能够一致可预测，
// 同时页面布局能适应各种设备或者屏幕尺寸（强烈推荐使用 flex 而非使用绝对定位的方式来进行页面结构的布局）

// flex-direction 定义了 flex 容器中 flex 成员项的排列方向。可选值为 row | column，默认值为 column
// （column：从上到下排列；row：从左到右排列）
func (s *Styles) FlexDirection(value string) *Styles {
	(*s)["flex-direction"] = value
	return s
}

// justify-content 定义了 flex 容器中 flex 成员项在主轴方向上如何排列以处理空白部分。可选值为
// flex-start | flex-end | center | space-between | space-around，默认值为 flex-start
// 	flex-start：是默认值，所有的 flex 成员项都排列在容器的前部；
// 	flex-end：则意味着成员项排列在容器的后部；
// 	center：即中间对齐，成员项排列在容器中间、两边留白；
// 	space-between：表示两端对齐，空白均匀地填充到 flex 成员项之间；
// 	space-around: 表示两端对齐，空白均匀地填充到 flex 成员项之前、之间、之后；
func (s *Styles) JustifyContent(value string) *Styles {
	(*s)["justify-content"] = value
	return s
}

// flex-wrap 指定 flex 元素单行显示还是多行显示。可选值为 nowrap | wrap，默认值为 nowrap
// （nowrap：flex 的元素被摆放到到一行, wrap：flex 元素支持换行，布局到多个行中）
func (s *Styles) FlexWrap(value string) *Styles {
	(*s)["flex-wrap"] = value
	return s
}

// align-items 定义了 flex 容器中 flex 成员项在纵轴方向上如何排列以处理空白部分。可选值为
// stretch | flex-start | center | flex-end，默认值为 stretch
// 	stretch 是默认值，即拉伸高度至 flex 容器的大小；
// 	flex-start 则是上对齐，所有的成员项排列在容器顶部；
// 	flex-end 是下对齐，所有的成员项排列在容器底部；
// 	center 是中间对齐，所有成员项都垂直地居中显示;
func (s *Styles) AlignItems(value string) *Styles {
	(*s)["align-items"] = value
	return s
}

// flex 属性定义了 flex 成员项可以占用容器中剩余空间的大小。如果所有的成员项设置相同的值
// flex: 1，它们将平均分配剩余空间. 如果一个成员项设置的值为 flex: 2，其它的成员项设置的值为 flex: 1，
// 那么这个成员项所占用的剩余空间是其它成员项的2倍。
func (s *Styles) Flex(value int) *Styles {
	(*s)["flex"] = value
	return s
}

//------------------------[position]-----------------------------

// position：设置定位类型。可选值为 relative | absolute ，默认值为 relative
func (s *Styles) Position(value string) *Styles {
	(*s)["position"] = value
	return s
}

func (s *Styles) Top(value string) *Styles {
	(*s)["top"] = value
	return s
}

func (s *Styles) Bottom(value string) *Styles {
	(*s)["bottom"] = value
	return s
}

func (s *Styles) Left(value string) *Styles {
	(*s)["left"] = value
	return s
}

func (s *Styles) Right(value string) *Styles {
	(*s)["right"] = value
	return s
}

//------------------------[text]-----------------------------

// color：文字颜色，可选值为色值，支持 RGB（ rgb(255, 0, 0) ）；RGBA（ rgba(255, 0, 0, 0.5) ）；十六进制（ #ff0000 ）；色值关键字（red）
func (s *Styles) Color(value string) *Styles {
	(*s)["color"] = value
	return s
}

// font-size {number}：文字大小
func (s *Styles) FontSize(value int) *Styles {
	(*s)["font-size"] = value
	return s
}

// font ：设置字体。在安卓平台，可以支持网络、assets、resource、文件、系统自带字体。默认为全局字体（设置过的话）或系统自带字体normal
func (s *Styles) Font(value string) *Styles {
	(*s)["font"] = value
	return s
}

// font-style: 设置文字风格。可选值normal 正常、bold 粗体、italic 斜体
func (s *Styles) FontStyles(value string) *Styles {
	(*s)["font-style"] = value
	return s
}

// text-align: 对齐方式。可选值 left | center | right | justify，默认值为 left
func (s *Styles) TextAlign(value string) *Styles {
	(*s)["text-align"] = value
	return s
}

// text-decoration (v1.35.0+):文字装饰。可选值underline | line-through
func (s *Styles) TextDecoration(value string) *Styles {
	(*s)["text-decoration"] = value
	return s
}

// vertical-align: 垂直方向的对齐方式。可选值 bottom | top | middle，默认值为top
func (s *Styles) VerticalAlign(value string) *Styles {
	(*s)["vertical-align"] = value
	return s
}

// text-overflow：设置内容超长时的省略样式。可选值 clip 截断 | ellipsis 省略，默认值为clip
func (s *Styles) TextOverflow(value string) *Styles {
	(*s)["text-overflow"] = value
	return s
}

// min-lines: 指定文本最少行数。默认值是 0
func (s *Styles) MinLines(value int) *Styles {
	(*s)["min-lines"] = value
	return s
}

// max-lines: 指定文本最大行数
func (s *Styles) MaxLines(value int) *Styles {
	(*s)["max-lines"] = value
	return s
}

// font-padding: 设定文字内边距。在安卓平台上，系统默认会自动为文字计算内边距，以避免文字的部分重叠。默认值为true
func (s *Styles) FontPadding(value bool) *Styles {
	(*s)["font-padding"] = value
	return s
}

// line-spacing: 调整行间距。在安卓平台上，通过line-spacing可以额外增加行间距。单位支持px、dp
func (s *Styles) LineSpacing(value string) *Styles {
	(*s)["line-spacing"] = value
	return s
}

// letter-spacing : 设置字间距。在安卓平台上，通过letter-spacing设置字间距。参数为一个浮点数，以标准字体的倍数作为字间距
func (s *Styles) LetterSpacing(value float64) *Styles {
	(*s)["letter-spacing"] = value
	return s
}

// white-space: 设置是否换行。文本长度超过一行时，继续在一行显示还是换行显示。可选值为normal | nowrap。默认值为normal
func (s *Styles) WhiteSpace(value string) *Styles {
	(*s)["white-space"] = value
	return s
}
