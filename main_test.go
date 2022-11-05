package main

import "testing"

func TestHello(t *testing.T) {
	if version != "1.0.0" {
		t.Errorf("Version is %s, expected 1.0.0", version)
	}
}
