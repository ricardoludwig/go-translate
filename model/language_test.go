package model

import (
	"testing"
)

var pt = Language{
	Name: "Português",
	Code: "pt",
}

var en = Language{
	Name: "English",
	Code: "en",
}

var languages = Languages{pt, en}

func TestLanguagesSortByCode(t *testing.T) {

	languages.SortByCode()
	//SortByCode(languages)

	if languages[0].Code != "en" {
		t.Error("Fail, expected english code en but got " + languages[0].Code)
		return
	}

}

func TestContainsLanguageInSlice(t *testing.T) {

	if languages.Contains(en) == false {
		t.Error("Fail, was expect found " + en.Code)
		return
	}
}
