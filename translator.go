package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ricardoludwig/translator/model"
	"github.com/ricardoludwig/translator/remote/http"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("\n Error: Args not found")
		help()
		os.Exit(1)
	}

	source := args[0]
	if source == "help" {
		help()
		os.Exit(0)
	}

	languages := getLanguages()
	if source == "languages" {
		languages.SortByCodeCopy().Print()
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

		langSource := model.Language{
			Code: source,
		}

		langTarget := model.Language{
			Code: target,
		}

		translated := translate("hello", langSource, langTarget, languages)
		fmt.Println(translated)
	}

}

func getLanguages() (response model.Languages) {

	resp, err := http.GetLanguages()

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	return resp

}

func translate(text string, source model.Language, target model.Language, languages model.Languages) (textTranslated string) {

	if languages.Contains(source) == false {
		fmt.Printf("Invalid language %v\n", source)
		os.Exit(1)
	}

	if languages.Contains(target) == false {
		fmt.Printf("Invalid language %v", source)
		os.Exit(1)
	}

	resp, err := http.Translate(text, source, target)

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	fmt.Printf("resposta %v\n", resp)
	return resp

}

func help() {

	fmt.Println("\nTranslator is a CLI application to translate words or sentences.")
	fmt.Println("I'm using this little programing to put into practice all the new stuff that I'm learning in Go.")
	fmt.Print("Translator uses libreTranslate API, during my tests a got some instability issues,")
	fmt.Println(" so is possible that happens occasionally.")
	fmt.Println("Usage:")
	fmt.Println("\t translator [command]")
	fmt.Println("Available commands:")
	fmt.Println("\tlanguages\t get available available languages to translate")
	fmt.Println("\ttranslate\t has three arguments source language, target language, and return (enter) command.")
	fmt.Println("\t\t\t after that, you can type a word or sentence to translate")
	fmt.Println("\thelp     \t show this")
	fmt.Println("Take a look at translator command, e.g.")
	fmt.Println("\ttranslator en pt [return]")
	fmt.Println("\tTranslate")
	fmt.Println("\thello")
	fmt.Println("\tOlá.")
	fmt.Println("Press q plus return to exit")

}
