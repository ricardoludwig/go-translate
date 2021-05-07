package http

import (
	"testing"
)

func TestValidURLExpecteSuccess(t *testing.T) {
	_, err := GetLanguages()
	if err != nil {
		t.Error("Fail, no response found")
		return
	}
}
