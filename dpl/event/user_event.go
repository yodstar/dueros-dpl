package event

// 设备端通过 SendEvent 指令上报，会对应转换 UserEvent 的事件形式上报给 bot 端，bot 可以通过
// 监听该事件类型，基于获取的组件id、事件类型、事件参数与绑定参数，执行下一步希望进行的交互逻辑。
type UserEvent struct {
	Type      string `json:"type"`      // 固定 UserEvent
	RequestId string `json:"requestId"` // 绑定组件的 token
	Timestamp string `json:"token"`     //
	Token     string `json:"token"`
	Payload   struct {
		ComponentId string `json:"componentId"`
		Source      struct {
			Type          string   `json:"type"`
			Handler       string   `json:"handler"`
			Value         string   `json:"value"`
			SelfArguments []string `json:"selfArguments"`
		} `json:"source"`
	} `json:"payload"`
}

type UserEventRequest struct {
	Request *UserEvent `json:"request"`
}
