package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ricardoludwig/go-translate/model"
	"github.com/ricardoludwig/go-translate/remote/http"
)

var languages []model.Language

func main() {

	languages := getLanguages()

	args := os.Args[1:]

	source := args[0]

	if source == "languages" {
		model.Print(languages)
		model.SortByCode(languages)
		model.Print(languages)
		os.Exit(0)
	}

	target := args[1]

	fmt.Println("Translate")

	inputReader := bufio.NewReader(os.Stdin)
	text, _ := inputReader.ReadString('\n')

	translated := translate(text, source, target)
	fmt.Println(translated)

}

func getLanguages() (response []model.Language) {

	resp, err := http.GetLanguages()

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	return resp

}

func translate(text string, source string, target string) (textTranslated string) {

	// sourceName := ""
	// targetName := ""

	// if len(languages) > 0 {

	// }

	langSource := model.Language{
		Code: source,
		Name: "English",
	}

	langTarget := model.Language{
		Code: target,
		Name: "English",
	}

	resp, err := http.Translate(text, langSource, langTarget)

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	return resp

}
