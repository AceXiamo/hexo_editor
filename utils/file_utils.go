package utils

import (
	"io/ioutil"
	"net/http"
)

func UrlToByte(url string) []byte {
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}