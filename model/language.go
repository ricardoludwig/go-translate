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

type Languages []Language

// SortByCode sorted by Language.code
func (ls Languages) SortByCode() {
	sort.Slice(ls, func(i, j int) bool { return ls[i].Code < ls[j].Code })
}

// SortByCodeCopy return a copy of the orign slice sorted by Language.code
func (ls Languages) SortByCodeCopy() (l Languages) {
	l = make(Languages, len(ls), len(ls))
	copy(l, ls)
	sort.Slice(l, func(i, j int) bool { return l[i].Code < l[j].Code })
	return l
}

func (ls Languages) Print() {
	response, err := json.Marshal(ls)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	fmt.Println(string(response))
}

func (ls Languages) Contains(lang Language) bool {
	for _, l := range ls {
		if l.equals(lang) {
			return true
		}
		//	if l == lang {
		//		return true
		//	} else if l.Code == lang.Code {
		//		return true
		//	}
	}
	return false
}

func (l Language) equals(lang Language) bool {
	return l.Code == lang.Code
}
