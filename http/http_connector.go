package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

	resp, err := http.Get("https://libretranslate.com/languages")
	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
		return languages, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return languages, err
	}

	json.Unmarshal([]byte(body), &languages)
	if len(languages) < 1 {
		log.Fatal("No response %v %v", string(body), err)
		return languages, err
	}

	return languages, nil

}
