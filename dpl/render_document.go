package dpl

import (
	"encoding/json"

	"github.com/dueros/bot-sdk-go/bot/directive"
)

// DPL 渲染指令结构(RenderDocument)
//
// 当一个技能服务收到来自设备端的请求时，它可以以一个 DPL.RenderDocument 页面渲染指令作为响应，
// 在 RenderDocument 指令中包含了一个 document 字段内容作为页面渲染的模板描述，
// 并指定一个唯一的 token 字段作为页面的标识下发该指令到支持 DPL 渲染的设备端上来渲染页面模板。
//
// 用于渲染 DPL 模板指令，通过下发该指令在设备端上渲染一个 DPL 页面
type RenderDocument struct {
	directive.BaseDirective           // type: 目前仅提供 DPL.RenderDocument 作为页面的渲染模板指令
	Token                   string    `json:"token"`    // token: 作为当前页面渲染指令在页面渲染后的唯一标识
	Document                *Document `json:"document"` // document: 包含渲染环境、指定渲染页面的 document 协议内容数据对象
}

// Unmarshal RenderDocument from json []byte.
func (r *RenderDocument) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}

// Marshal RenderDocument to []byte
func (r *RenderDocument) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *RenderDocument) SetDocument(doc *Document) {
	r.Document = doc
}

func NewRenderDocument(doc *Document) *RenderDocument {
	r := &RenderDocument{}
	r.Type = "DPL.RenderDocument"
	r.Token = r.GenToken()
	r.Document = doc
	return r
}
