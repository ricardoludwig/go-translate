package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ricardoludwig/go-translate/remote/http"
)

func main() {

	args := os.Args[1:]
	source := args[0]
	target := args[1]

	fmt.Println("Translate")

	inputReader := bufio.NewReader(os.Stdin)
	text, _ := inputReader.ReadString('\n')

	translated := translate(text, source, target)
	fmt.Println(translated)

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

func translate(text string, source string, target string) (textTranslated string) {

	en := http.SupportLanguage{
		Code:     source,
		Language: "English",
	}

	pt := http.SupportLanguage{
		Code:     target,
		Language: "English",
	}

	resp, err := http.Translate(text, en, pt)

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	return resp

}
