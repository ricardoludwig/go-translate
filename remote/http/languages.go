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

func Translate(text string, source SupportLanguage, target SupportLanguage) (response string, err error) {

	resp, err := HttpPost("https://libretranslate.com/translate", translateBody(text, source, target))

	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	return resp, nil
}

func translateBody(text string, source SupportLanguage, target SupportLanguage) (body string) {
	//body := "q=hello&source=en&target=pt&api_key=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	return "q=" + text + "&source=" + source.Code + "&target=" + target.Code + "&api_key=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
