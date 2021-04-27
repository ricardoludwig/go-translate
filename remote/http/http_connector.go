package http

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const invalidURLMsg = "Invalid URL"

var ErrorInvalidURL = errors.New(invalidURLMsg)

func HttpGet(url string) (response string, err error) {

	if len(url) < 1 {
		log.Println(invalidURLMsg)
		return "", fmt.Errorf("%v %w", invalidURLMsg, ErrorInvalidURL)
	}

	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		log.Println(err)
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(body), nil

}
