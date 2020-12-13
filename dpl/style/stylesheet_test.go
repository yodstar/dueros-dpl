package style

import (
	"testing"
)

func TestStylesheet(t *testing.T) {
	st := NewStyles("./styles.json").
		Height("44dp").
		PaddingLeft("24dp").
		PaddingRight("24dp").
		MinWidth("120dp").
		MaxWidth("24dp").
		JustifyContent("center").
		AlignItems("center").
		BackgroundColor("rgba(255, 255, 255, 0.1)").
		BorderRadius("18dp 20dp 20dp 0dp")
	if err := st.GetLastError(); err != nil {
		t.Log(err)
	}

	if data, err := st.Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}

	cl := (&ClassList{}).SetClass("bubble_Hint", st)
	mq := NewClassMediaQuery("x5").SetClass("test-height", NewStyles().Height("220dp"))
	ss := NewStylesheet("./stylesheet.json").Append(cl, mq)
	if err := ss.GetLastError(); err != nil {
		t.Log(err)
	}

	if data, err := ss.Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}
}
