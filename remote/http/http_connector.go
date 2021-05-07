package http

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	netHttp "net/http"
	"strings"
	"time"
)

const invalidURLMsg = "Invalid URL"

var ErrorInvalidURL = errors.New(invalidURLMsg)

var getClient = &netHttp.Client{
	Timeout: time.Second * 35,
}

func HttpGet(url string) (response string, err error) {

	if validURL(url) != nil {
		return "", fmt.Errorf("%v %w", invalidURLMsg, ErrorInvalidURL)
	}

	resp, err := getClient.Get(url)
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

var postClient = &netHttp.Client{
	Timeout: time.Second * 35,
}

func HttpPost(url string, body string) (response string, err error) {

	if validURL(url) != nil {
		return "", fmt.Errorf("%v %w", invalidURLMsg, ErrorInvalidURL)
	}

	resp, err := postClient.Post(url, "application/x-www-form-urlencoded", strings.NewReader(body))

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return "", err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(respBody), nil

}

func validURL(url string) (err error) {
	if len(url) < 1 {
		log.Println(invalidURLMsg)
		return ErrorInvalidURL
	}
	return nil
}
