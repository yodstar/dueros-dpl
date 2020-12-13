package dpl

// 页面渲染配置信息
type RenderConfig struct {
	Viewport          *Viewport `json:"viewport"`          // viewport: 页面可视区域大小
	GlobalVslDisabled bool      `json:"globalVslDisabled"` // globalVslDisabled:
}

func NewRenderConfig(width, height string, globalVslDisabled bool) *RenderConfig {
	return &RenderConfig{
		Viewport: &Viewport{
			Width:  width,
			Height: height,
		},
		GlobalVslDisabled: globalVslDisabled,
	}
}

// 目前，dpl 模板中应用的默认自适应方式是以设定 viewport（具体含义可参考下面一小节中的描述）
// 中宽为 960dp的方式，并以在document文档协议描述中，基于使用具体样式宽高属性数值对各个组件
// 使用和布局使用，来实现各个设备上以宽度铺满情况下的默认渲染展现。
type Viewport struct {
	Width  string `json:"width"`
	Height string `json:"height"`
}
