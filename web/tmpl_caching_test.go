package web

import "testing"

func TestTemplateCaching(t *testing.T) {
	if IndexTemplate == nil {
		t.Errorf("die Variable `indexTemplate` ist nil")
	}
}
