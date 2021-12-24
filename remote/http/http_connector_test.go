package http

import (
	"testing"
)

func TestGetWithSucess(t *testing.T) {

	_, err := Get("http://google.com")

	if err != nil {
		t.Error("Fail, no response found")
		return
	}

}

func TestGetInvalidURL(t *testing.T) {

	response, err := Get("")
	if len(response) > 1 {
		t.Error("Fail, no response was expected to invalid request URL")
		return
	}

	if err == nil {
		t.Error("Expected a fail request to invalid URL")
	}

}
