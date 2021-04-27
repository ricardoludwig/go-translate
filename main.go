package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ricardoludwig/go-translate/remote/http"
)

func main() {

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
