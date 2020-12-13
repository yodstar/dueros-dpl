package command

import (
	"testing"

	"github.com/yodstar/dueros-dpl/dpl/style"
)

func TestComponentSpecified(t *testing.T) {
	if data, err := Scroll("component-specified", "20dp").Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}

	if data, err := ScrollToIndex("component-specified", 0, "center").Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}

	if data, err := ScrollToElement("component-specified", "tid", "center").Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}

	if data, err := AutoPage("component-specified", 500).Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}

	if data, err := SetPage("component-specified", "relative", 0).Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}
}

func TestPageControl(t *testing.T) {
	if data, err := PageDestroy().Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}

	if data, err := ResetNonInteractionExitTime(500).Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}

	if data, err := SetSpeechMonitor("page-control", 1).Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}
}

func TestMedia(t *testing.T) {
	if data, err := ControlMedia("media", "play").Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}
}

func TestDateUpdating(t *testing.T) {
	d := make(map[string]interface{})
	d["data_key"] = "data_value"
	if data, err := UpdateDataSource(d).Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}

	d["styles"] = style.NewStyles().Width("100dp").Padding("5dp")
	if data, err := SetStyles("date-updating", d).Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}

	d["prop_key"] = "prop_value"
	if data, err := SetProps("date-updating", d).Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}

	if data, err := SetState("date-updating", "width", "200dp").Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}

	if data, err := UpdateComponent("date-updating", d).Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}

	if data, err := AppendComponent("date-updating", d).Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}
}

func TestSendEvent(t *testing.T) {
	if data, err := SendEvent("message-interaciton", []interface{}{"arg1"}).Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}
}

func TestMiscellaneous(t *testing.T) {
	if data, err := InvokeMethod("miscellaneous", "method1", []interface{}{"arg1"}).Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}

	if data, err := Animation("miscellaneous", "rotation", "0", "100", "shake(100, 20, 0, 300)", 500).Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}

	commands := make([]*Command, 0)
	commands = append(commands, PageDestroy())
	if data, err := Parallel("miscellaneous", 1, 500, commands).Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}
}
