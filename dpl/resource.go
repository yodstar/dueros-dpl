package dpl

type Resource struct {
	Description string            `json:"description"` // description: 对属性常量的描述，解析的时候会被忽略
	Colors      map[string]string `json:"colors"`      // colors: 定义颜色的属性常量，只能是颜色值
	Dimensions  map[string]string `json:"dimensions"`  // dimensions: 定义其他属性值, 如宽高、左右间距
}
