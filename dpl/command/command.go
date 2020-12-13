package command

import (
	"encoding/json"
	"io/ioutil"

	"github.com/yodstar/dueros-dpl/dpl/util"
)

type Command struct {
	Error       error  `json:"-"`
	Type        string `json:"type"`                  // type: 指令类型
	ComponentId string `json:"componentId,omitempty"` // componentId: 组件 id
	Delay       int    `json:"delay,omitempty"`       // delay: 指令在端上的延迟执行时间，单位毫秒，默认为 0 (optional)
	DWhen       string `json:"dWhen,omitempty"`       // dWhen: 指令的条件执行表达式 (optional)

	Align     string        `json:"align,omitempty"`
	Arguments []interface{} `json:"arguments,omitempty"`
	Attribute string        `json:"attribute,omitempty"`

	Command  string     `json:"command,omitempty"`
	Commands []*Command `json:"commands,omitempty"`

	Data                  map[string]interface{} `json:"data,omitempty"`
	DelayInMilliseconds   int                    `json:"delayInMilliseconds,omitempty"`
	DialogType            string                 `json:"dialogType,omitempty"`
	Distance              string                 `json:"distance,omitempty"`
	Document              interface{}            `json:"document,omitempty"`
	Duration              int                    `json:"duration,omitempty"`
	DurationInMillisecond int                    `json:"durationInMillisecond,omitempty"`

	Easing     string `json:"easing,omitempty"`
	From       string `json:"from,omitempty"`
	Index      *int   `json:"index,omitempty"`
	MethodName string `json:"methodName,omitempty"`

	OnFinished    *Command `json:"onFinished,omitempty"`
	OnInterrupted *Command `json:"onInterrupted,omitempty"`

	Params   []interface{}          `json:"params,omitempty"`
	Position string                 `json:"position,omitempty"`
	Props    map[string]interface{} `json:"props,omitempty"`

	RepeatCount interface{} `json:"repeatCount,omitempty"`
	RepeatMode  string      `json:"repeatMode,omitempty"`

	SpeechFinishedPosition int                    `json:"speechFinishedPosition,omitempty"`
	Styles                 map[string]interface{} `json:"styles,omitempty"`
	StateType              string                 `json:"stateType,omitempty"`
	State                  string                 `json:"state,omitempty"`

	TargetComponentId string `json:"targetComponentId,omitempty"`
	To                string `json:"to,omitempty"`

	Value interface{} `json:"value,omitempty"`
}

// Unmarshal Command from json []byte.
func (cmd *Command) Unmarshal(data []byte) error {
	return json.Unmarshal(data, cmd)
}

// Marshal Command to []byte
func (cmd *Command) Marshal() ([]byte, error) {
	return json.Marshal(cmd)
}

// Load Command from json file.
func (cmd *Command) GetCommandFromFile(path ...string) *Command {
	for i, _ := range path {
		data, err := ioutil.ReadFile(path[i])
		if err != nil {
			cmd.Error = util.Error(err)
			return cmd
		}
		if err = json.Unmarshal(data, cmd); err != nil {
			cmd.Error = util.Error(err)
			return cmd
		}
	}
	return cmd
}

func (cmd *Command) SetOnFinished(onFinished *Command) *Command {
	cmd.OnFinished = onFinished
	return cmd
}

func (cmd *Command) SetOnInterrupted(onInterrupted *Command) *Command {
	cmd.OnInterrupted = onInterrupted
	return cmd
}
