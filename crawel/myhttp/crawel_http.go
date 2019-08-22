package myhttp

import (
	"crawel/global"
	"io/ioutil"
	"net/http"
)

// 基于爬虫的请求伪装

func SetMyRequestHeader(url string) []byte {
	resp, err := http.Get(url)
	global.CheckErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}
