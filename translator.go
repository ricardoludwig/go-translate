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

	if len(args) < 1 {
		fmt.Println("\n Error: no args found\n\n")
		help()
		os.Exit(1)
	}

	source := args[0]

	if source == "help" {
		help()
		os.Exit(0)
	}

	if source == "languages" {
		model.Print(languages)
		model.SortByCode(languages)
		model.Print(languages)
		os.Exit(0)
	}

	target := args[1]

	fmt.Println("Translate")
	fmt.Println("Press q plus return to exit")

	for {

		inputReader := bufio.NewReader(os.Stdin)
		text, _ := inputReader.ReadString('\n')

		if text == "q\n" {
			fmt.Println("Good by!")
			os.Exit(0)
		}

		translated := translate(text, source, target)
		fmt.Println(translated)
	}

}

func help() {

	fmt.Println("\nTranslator is a CLI application to translate words or sentences.")
	fmt.Println("I'm using this little programing to put into practice all the new stuff that I'm learning in Go.")
	fmt.Print("Translator uses libreTranslate API, during my tests a got some instability issues,")
	fmt.Println(" so is possible that happens occasionally.")
	fmt.Println("Usage:")
	fmt.Println("\t translator [command]\n")
	fmt.Println("Available commands:")
	fmt.Println("\tlanguages\t get available available languages to translate")
	fmt.Println("\ttranslate\t has three arguments source language, target language, and return (enter) command.")
	fmt.Println("\t\t\t after that, you can type a word or sentence to translate")
	fmt.Println("\thelp     \t show this\n")
	fmt.Println("Take a look at translator command, e.g.")
	fmt.Println("\ttranslator en pt [return]")
	fmt.Println("\tTranslate")
	fmt.Println("\thello")
	fmt.Println("\tOlá.")
	fmt.Println("Press q plus return to exit")

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
