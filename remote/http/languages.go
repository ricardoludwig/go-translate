package http

import (
	"encoding/json"
	"log"
)

//SupportLanguage -  representation type of suport languages
type SupportLanguage struct {
	Code     string `json:"code"`
	Language string `json:"name"`
}

// GetSupportLanguages - this function make a Http Get
// request to libretranslate.com/languages and obtain an
// array of support languages to translation
func GetSupportLanguages() (response []SupportLanguage, err error) {

	var languages []SupportLanguage

	resp, err := HttpGet("https://libretranslate.com/languages")

	if err != nil {
		log.Fatalln(err)
		return languages, err
	}

	json.Unmarshal([]byte(resp), &languages)
	if len(languages) < 1 {
		log.Fatal("No response")
		return languages, err
	}

	return languages, nil

}
