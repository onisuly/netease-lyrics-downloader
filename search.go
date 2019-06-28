package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const pubKey = "010001"
const nonce = "0CoJUm6Qyw8W8jud"
const modulus = "00e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7b725152b3ab17a876aea8a5aa76d2e417629ec" +
	"4ee341f56135fccf695280104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932575cce10b424d813cfe4875d3e82" +
	"047b97ddef52741d546b8e289dc6935b3ece0462db0a22b8e7"
const queryURL = "http://music.163.com/weapi/cloudsearch/get/web?csrf_token="

func SearchSong(text string, transport *http.Transport, timeout time.Duration) (string, error) {
	query := &Query{S: text, Type: 1, Offset: 0, Sub: "false", Limit: 9}
	q, err := json.Marshal(query)
	if err != nil {
		return "", err
	}
	randomString := RandomString(16)
	encText := AESEncrypt(string(q), nonce)
	encText = AESEncrypt(encText, randomString)
	encSecKey := RSAEncrypt(randomString, pubKey, modulus)

	var req http.Request
	err = req.ParseForm()
	if err != nil {
		return "", err
	}
	req.Form.Add("params", encText)
	req.Form.Add("encSecKey", encSecKey)
	body := strings.NewReader(req.Form.Encode())
	newReq, err := http.NewRequest("POST", queryURL, body)
	if err != nil {
		return "", err
	}
	client := &http.Client{Transport: transport, Timeout: timeout}
	newReq.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) "+
		"Chrome/74.0.3729.131 Safari/537.36")
	newReq.Header.Set("Host", "music.163.com")
	newReq.Header.Set("Referer", "http://music.163.com/search/")
	newReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	var resp *http.Response
	for resp, err = client.Do(newReq); err != nil; {
		resp, err = client.Do(newReq)
	}
	defer Close(resp.Body)

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

type Query struct {
	S      string `json:"s"`
	Type   int    `json:"type"`
	Offset int    `json:"offset"`
	Sub    string `json:"sub"`
	Limit  int    `json:"limit"`
}
