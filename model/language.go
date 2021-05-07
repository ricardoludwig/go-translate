package model

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

//Language -  representation type of suport languages
type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func SortByCode(l []Language) {
	sort.Slice(l, func(i, j int) bool { return l[i].Code < l[j].Code })
}

func Print(resp []Language) {
	response, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	fmt.Println(string(response))
}
