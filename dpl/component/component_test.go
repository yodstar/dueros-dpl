package component

import (
	"testing"

	"github.com/yodstar/dueros-dpl/dpl/command"
)

func TestComponent(t *testing.T) {
	audio := Audio("component-audio", "http://localhost/upload/audio_001.mp3").
		SetPropsJSON(`{"looping":false,"autoplay":true,"type":"music"}`).
		SetStylesCSS("height:220dp;width:300dp").
		BindEvents("onEnd", command.SendEvent("demo_audio_1", []interface{}{"event_audio_end"}))
	if audio.Error != nil {
		t.Log(audio.Error)
	}

	if data, err := audio.Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}

	container := Container("component-container").
		SetPropsJSON(`{"looping":false,"autoplay":true,"type":"music"}`).
		SetStylesCSS("height:220dp;width:300dp").
		BindEvents("onEnd", command.SendEvent("demo_audio_1", []interface{}{"event_audio_end"}))
	if container.Error != nil {
		t.Log(container.Error)
	}

	if data, err := container.Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}
}
