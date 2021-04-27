package http

import (
	"testing"
)

func giveValidURLExpecteSuccess(t *testing.T) {
	_, err := GetSupportLanguages()
	if err != nil {
		t.Error("Fail, no response found")
		return
	}
}
