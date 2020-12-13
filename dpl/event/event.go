package event

type Event struct {
	Type        string   `json:"type"`
	ComponentId string   `json:"componentId"`
	Arguments   []string `json:"arguments"`
}

func NewEvent(eventType, componentId string, arguments ...string) *Event {
	return &Event{
		Type:        eventType,
		ComponentId: componentId,
		Arguments:   arguments,
	}
}
