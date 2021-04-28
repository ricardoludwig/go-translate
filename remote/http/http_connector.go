package http

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	netHttp "net/http"
	"time"
)

const invalidURLMsg = "Invalid URL"

var netClient = &netHttp.Client{
	Timeout: time.Second * 5,
}

var ErrorInvalidURL = errors.New(invalidURLMsg)

func HttpGet(url string) (response string, err error) {

	if len(url) < 1 {
		log.Println(invalidURLMsg)
		return "", fmt.Errorf("%v %w", invalidURLMsg, ErrorInvalidURL)
	}

	resp, err := netClient.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(body), nil

}
