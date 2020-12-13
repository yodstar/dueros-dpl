package dpl

import (
	"testing"
)

func TestRenderDocument(t *testing.T) {
	path := []string{
		"./../example/template/dpl-component01.json",
	}

	doc := NewDocument().
		AddParameter("payload").
		AddComponentFromFile(path...)

	if doc.Error != nil {
		t.Log(doc.Error)
		t.Fail()
	}

	if data, err := doc.GetRenderDocument().Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}

	if data, err := NewDocument().Marshal(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(string(data))
	}
}
