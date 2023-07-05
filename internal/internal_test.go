package internal

import "testing"

func TestInternal_EmptyString(t *testing.T) {
	if EmptyString != "" {
		t.Errorf("empty string not empty")
	}
}
