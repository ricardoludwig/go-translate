package main

import (
	"encoding/json"
	"fmt"
	"github.com/ricardoludwig/go-translate/remote/http"
	"os"
)

func main() {

	fmt.Println(translate())

}

func getLanguages() {

	resp, err := http.GetSupportLanguages()

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	response, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	fmt.Println(string(response))

}

func translate() (textTranslated string) {

	en := http.SupportLanguage{
		Code:     "en",
		Language: "English",
	}

	pt := http.SupportLanguage{
		Code:     "pt",
		Language: "English",
	}

	resp, err := http.Translate("hello", en, pt)

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	return resp

}
