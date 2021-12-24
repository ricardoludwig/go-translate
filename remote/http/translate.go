package http

import (
	"encoding/json"
	"log"

	"github.com/ricardoludwig/translator/model"
)

// GetLanguages - this function make a Http Get
// request to libretranslate.com/languages and obtain an
// array of support languages to translation
func GetLanguages() (response model.Languages, err error) {

	var languages model.Languages

	resp, err := Get("https://libretranslate.com/languages")

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

type TranslatedText struct {
	RespText string `json:"translatedText"`
}

func Translate(text string, source model.Language, target model.Language) (response string, err error) {

	resp, err := Post("https://libretranslate.com/translate", translateBody(text, source, target))

	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	var translatedText TranslatedText
	json.Unmarshal([]byte(resp), &translatedText)

	log.Printf("TranslatedText %v\n", translatedText)

	return translatedText.RespText, nil
}

func translateBody(text string, source model.Language, target model.Language) (body string) {
	return "q=" + text + "&source=" + source.Code + "&target=" + target.Code + "&api_key=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
