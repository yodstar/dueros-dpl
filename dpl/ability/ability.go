package ability

import (
	"github.com/yodstar/dueros-dpl/dpl/event"
)

// Ability: 能力组件，目前支持了手势（Gesture），表单（Form）能力
type Ability struct {
	// type: {{ENUM}}
	Type string `json:"type"`
	// componentId:
	ComponentId string `json:"componentId"`
	// events:
	Events map[string][]*event.Event `json:"events"`
}

func (a *Ability) SetEvents(name string, events []*event.Event) {
	a.Events[name] = events
}

func NewAbility(abilityType, componentId string) *Ability {
	return &Ability{
		Type:        abilityType,
		ComponentId: componentId,
	}
}
