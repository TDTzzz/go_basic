package adapter

import "testing"

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	adapter := NewAdapter(adaptee)

	res := adapter.Request()
	if res != "adaptee method" {
		t.Fatalf("%s", res)
	}
}
